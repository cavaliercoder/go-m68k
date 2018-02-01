package m68k

//go:generate ./cmd/m68kgen/m68kgen -out opcodes.go

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"unsafe"
)

const (
	SizeByte = iota
	SizeWord
	SizeLong
)

// Status Register bits.
// See section 1.1.4
const (
	StatusCarry    = 1 << iota // C
	StatusOverflow             // V
	StatusZero                 // Z
	StatusNegative             // N
	StatusExtend               // X

	StatusMask = 0xA71F // Section 1.3.2
)

var (
	ErrNoProgram = errors.New("no program loaded or memory device attached")
)

// A Processor emulates the Motorola 68000 microprocessor.
type Processor struct {
	D  [8]uint32 // Data registers
	A  [8]uint32 // Address registers
	SR uint32    // Status Register
	PC uint32    // Program Counter
	M  Memory    // System memory controller

	// TraceWriter specifies where trace log output should be written to. If
	// TraceWriter is nil, no logging is performed.
	TraceWriter io.Writer

	buf      [64]byte // general purpose buffer
	op       uint16
	err      error
	handlers [256]TrapHandler
}

type TrapHandler interface {
	Exception(*Processor, int) error
}
type TrapHandlerFunc func(*Processor, int) error

func (f TrapHandlerFunc) Exception(p *Processor, v int) error {
	return f(p, v)
}

func (c *Processor) RegisterTrapHandler(vector int, f TrapHandler) error {
	if vector < 0 || vector >= len(c.handlers) {
		return errors.New("vector out of range")
	}
	c.handlers[vector] = f
	return nil
}

// Reset resets all processor and memory state.
func (c *Processor) Reset() {
	// TODO: this seems important: https://stackoverflow.com/a/38244832/5809680
	*c = Processor{
		TraceWriter: c.TraceWriter,
	}
	c.Jump(0x1000)
	if c.M == nil {
		c.M = NewRAM(0x40000) // 256KB
	}
	clearMemory(c.M)
}

func clearMemory(m Memory) {
	if m == nil {
		return
	}
	addr := 0
	zero := [32]byte{}
	for {
		n, err := m.Write(addr, zero[:])
		if err == io.ErrShortWrite || err == ErrAddressOutOfBounds {
			break
		}
		if err != nil {
			panic(err)
		}
		addr += n
	}
}

// Jump sets the value of the program counter to the given memory address.
func (c *Processor) Jump(addr uint32) error {
	c.PC = addr
	return nil
}

// Run executes any program loaded into memory, starting from the program
// counter value, running until completion.
func (c *Processor) Run() error {
	if c.M == nil {
		return ErrNoProgram
	}
	for c.err == nil {
		c.Step()
	}
	return c.err
}

// Step executes the single instruction located at the address specified by the
// program counter register.
func (c *Processor) Step() error {
	// read from program pointer into op
	if _, c.err = c.M.Read(int(c.PC), c.buf[0:2]); c.err != nil {
		return c.err
	}
	c.op = uint16(c.buf[0])<<8 + uint16(c.buf[1])
	if c.op == 0 {
		// TODO: this is a valid instruction. Find a way to terminate programs
		// without buffer overrun.
		c.err = io.EOF
		return c.err
	}

	// try new fnmap
	done := false
	fn := defaultFuncMap.Resolve(c.op)
	if fn != nil {
		t := fn(c)
		if c.err == nil {
			c.trace(t)
			done = true
		} else if _, ok := c.err.(opcodeError); ok {
			c.err = nil
		} else {
			done = true
		}
	}

	// use generated opcodes
	if !done {
		f := c.mapFn(c.op)
		if f == nil {
			c.err = newOpcodeError(c.op)
		} else {
			f()
		}
	}
	return c.err
}

func (c *Processor) tracef(format string, a ...interface{}) {
	if c.TraceWriter == nil {
		return
	}
	fmt.Fprintf(c.TraceWriter, format, a...)
}

func (c *Processor) trace(v ...interface{}) {
	if c.TraceWriter == nil {
		return
	}
	fmt.Fprintln(c.TraceWriter, v...)
}

// DumpState prints the state of of a 68000 processor to the given writer in
// hexidecimal format.
func DumpState(w io.Writer, p *Processor) {
	for i := 0; i < len(p.D); i++ {
		fmt.Fprintf(w, "D%d: %08X A%d: %08X\n", i, p.D[i], i, p.A[i])
	}
}

// byteToInt32 sign extends the given byte to an Int32.
func byteToInt32(b byte) int32 {
	if b&0x80 == 0 {
		return int32(b)
	}
	v := 0xFFFFFF00 | uint32(b)
	return *(*int32)(unsafe.Pointer(&v))
}

// wordToInt32 sign extends the given word to an Int32.
func wordToInt32(n uint16) int32 {
	if n&0x8000 == 0 {
		return int32(n)
	}
	v := 0xFFFF0000 | uint32(n)
	return *(*int32)(unsafe.Pointer(&v))
}

// decodeEA returns the effective address segment (bits 0 - 5) of an opcode.
func decodeEA(op uint16) uint16 {
	return op & 0x003F
}

// decodeExEA returns the effective destination address segment of an opcode,
// for instructions that have an effective address in bit 6 - 11. The result is
// flipped so that the mode and register portions of the ea match the encoding
// used by typical ea fields.
func decodeExEA(op uint16) uint16 {
	return ((op & 0x01C0) >> 3) | ((op & 0x0E00) >> 9)
}

func (c *Processor) Byte(addr int) (b byte, err error) {
	if _, err = c.M.Read(addr, c.buf[:1]); err != nil {
		return
	}
	b = c.buf[0]
	return
}

func (c *Processor) Word(addr int) (n uint16, err error) {
	if _, err = c.M.Read(addr, c.buf[:2]); err != nil {
		return
	}
	n = uint16(c.buf[0])<<8 | uint16(c.buf[1])
	return
}

func (c *Processor) SWord(addr int) (n int16, err error) {
	var u uint16
	u, err = c.Word(addr)
	if err != nil {
		return
	}
	n = *(*int16)(unsafe.Pointer(&u))
	return
}

func (c *Processor) Long(addr int) (n uint32, err error) {
	if _, err = c.M.Read(addr, c.buf[:4]); err != nil {
		return
	}
	n = uint32(c.buf[0])<<24 | uint32(c.buf[1])<<16 | uint32(c.buf[2])<<8 | uint32(c.buf[3])
	return
}

func (c *Processor) SLong(addr int) (n int32, err error) {
	var u uint32
	if u, err = c.Long(addr); err != nil {
		return
	}
	n = *(*int32)(unsafe.Pointer(&u))
	return
}

func (c *Processor) readByte(ea uint16) (b byte, opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = newOpcodeError(c.op)
		return

	case 0x00: // data register
		b = byte(c.D[reg])
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		b = byte(c.A[reg])
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		b, err = c.Byte(int(c.A[reg]))
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		b, err = c.Byte(int(c.A[reg]))
		c.A[reg]++
		if reg == 7 {
			c.A[reg]++ // align stack pointer to 16-bit
		}
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg]--
		if reg == 7 {
			c.A[reg]-- // align stack pointer to 16-bit
		}
		b, err = c.Byte(int(c.A[reg]))
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.SWord(int(c.PC))
		if err != nil {
			return
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		b, err = c.Byte(addr)
		opr = fmt.Sprintf("(%d,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = newOpcodeError(c.op)
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			b, err = c.Byte(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			b, err = c.Byte(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x04: // immediate byte
			var n uint16
			n, err = c.Word(int(c.PC))
			b = byte(n)
			c.PC += 2
			opr = fmt.Sprintf("#$%X", b)
		}
	}
	return
}

func (c *Processor) readWord(ea uint16) (n uint16, opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = newOpcodeError(c.op)
		return

	case 0x00: // data register
		n = uint16(c.D[reg])
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		n = uint16(c.A[reg])
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		n, err = c.Word(int(c.A[reg]))
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		n, err = c.Word(int(c.A[reg]))
		c.A[reg] += 2
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 2
		n, err = c.Word(int(c.A[reg]))
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.SWord(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		n, err = c.Word(addr)
		opr = fmt.Sprintf("(%d,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = newOpcodeError(c.op)
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			n, err = c.Word(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			n, err = c.Word(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x04: // immediate word
			n, c.err = c.Word(int(c.PC))
			c.PC += 2
			opr = fmt.Sprintf("#$%X", n)
		}
	}
	return
}

func (c *Processor) readLong(ea uint16) (n uint32, opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = newOpcodeError(c.op)
		return

	case 0x00: // data register
		n = c.D[reg]
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		n = c.A[reg]
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		n, err = c.Long(int(c.A[reg]))
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		n, err = c.Long(int(c.A[reg]))
		c.A[reg] += 4
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 4
		n, err = c.Long(int(c.A[reg]))
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.SWord(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		n, err = c.Long(addr)
		opr = fmt.Sprintf("(%d,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = newOpcodeError(c.op)
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			n, err = c.Long(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			n, err = c.Long(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x04: // immediate long
			n, c.err = c.Long(int(c.PC))
			c.PC += 4
			opr = fmt.Sprintf("#$%X", n)
		}
	}
	return
}

func (c *Processor) writeByte(ea uint16, b byte) (opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = newOpcodeError(c.op)
		return

	case 0x00: // data register
		c.D[reg] = c.D[reg]&0xFFFFFF00 | uint32(b)
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		c.A[reg] = c.A[reg]&0xFFFFFF00 | uint32(b)
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		_, err = c.M.Write(int(c.A[reg]), []byte{b})
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		_, err = c.M.Write(int(c.A[reg]), []byte{b})
		c.A[reg]++
		if reg == 7 {
			c.A[reg]++ // align stack pointer to 16-bit
		}
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg]--
		if reg == 7 {
			c.A[reg]-- // align stack pointer to 16-bit
		}
		_, err = c.M.Write(int(c.A[reg]), []byte{b})
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.SWord(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		_, err = c.M.Write(int(addr), []byte{b})
		opr = fmt.Sprintf("(%d,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = newOpcodeError(c.op)
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			_, err = c.M.Write(int(addr), []byte{b})
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			_, err = c.M.Write(int(addr), []byte{b})
			opr = fmt.Sprintf("$%X", addr)
		}
	}
	return
}

func (c *Processor) writeWord(ea uint16, v uint16) (opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = newOpcodeError(c.op)
		return

	case 0x00: // data register
		c.D[reg] = c.D[reg]&0xFFFF0000 | uint32(v)
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		c.A[reg] = c.A[reg]&0xFFFF0000 | uint32(v)
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		binary.BigEndian.PutUint16(c.buf[:2], v)
		_, err = c.M.Write(int(c.A[reg]), c.buf[:2])
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		binary.BigEndian.PutUint16(c.buf[:2], v)
		_, err = c.M.Write(int(c.A[reg]), c.buf[:2])
		c.A[reg] += 2
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 2
		binary.BigEndian.PutUint16(c.buf[:2], v)
		_, err = c.M.Write(int(c.A[reg]), c.buf[:2])
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.SWord(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		binary.BigEndian.PutUint16(c.buf[:2], v)
		_, err = c.M.Write(addr, c.buf[:2])
		opr = fmt.Sprintf("(%d,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = newOpcodeError(c.op)
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			binary.BigEndian.PutUint16(c.buf[:2], v)
			_, err = c.M.Write(int(addr), c.buf[:2])
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			binary.BigEndian.PutUint16(c.buf[:2], v)
			_, err = c.M.Write(int(addr), c.buf[:2])
			opr = fmt.Sprintf("$%X", addr)
		}
	}
	return
}

func (c *Processor) writeLong(ea uint16, v uint32) (opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = newOpcodeError(c.op)
		return

	case 0x00: // data register
		c.D[reg] = v
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		c.A[reg] = v
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		binary.BigEndian.PutUint32(c.buf[:4], v)
		_, err = c.M.Write(int(c.A[reg]), c.buf[:4])
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		binary.BigEndian.PutUint32(c.buf[:4], v)
		_, err = c.M.Write(int(c.A[reg]), c.buf[:4])
		c.A[reg] += 4
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 4
		binary.BigEndian.PutUint32(c.buf[:4], v)
		_, err = c.M.Write(int(c.A[reg]), c.buf[:4])
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.SWord(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		binary.BigEndian.PutUint32(c.buf[:4], v)
		_, err = c.M.Write(addr, c.buf[:4])
		opr = fmt.Sprintf("(%d,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = newOpcodeError(c.op)
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			binary.BigEndian.PutUint32(c.buf[:4], v)
			_, err = c.M.Write(int(addr), c.buf[:4])
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			binary.BigEndian.PutUint32(c.buf[:4], v)
			_, err = c.M.Write(int(addr), c.buf[:4])
			opr = fmt.Sprintf("$%X", addr)
		}
	}
	return
}

type stepFunc func(*Processor) *stepTrace
