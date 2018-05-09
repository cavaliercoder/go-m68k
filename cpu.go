package m68k

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/cavaliercoder/go-m68k/m68kmem"
)

const (
	SizeByte = iota
	SizeWord
	SizeLong
)

// Status Register bits.
// See section 1.1.4.
const (
	StatusCarry    = uint32(1) << iota // C
	StatusOverflow                     // V
	StatusZero                         // Z
	StatusNegative                     // N
	StatusExtend                       // X

	StatusMask = 0xA71F // Section 1.3.2
)

// Comparison conditions.
// See table 3-19.
const (
	CondTrue = iota
	CondFalse
	CondHigh
	CondLowOrSame
	CondCarryClear
	CondCarrySet
	CondNotEqual
	CondEqual
	CondOverflowClean
	CondOverflowSet
	CondPlus
	CondMinus
	CondGreaterOrEqual
	CondLessThan
	CondGreaterThan
	CondLessOrEqual
)

// string suffix for conditional opcodes
// see: table 3-19
var conditionStrings = []string{
	"t",  // true
	"f",  // false
	"hi", // high
	"ls", // low or same
	"cc", // (hi) carry clear
	"cs", // (lo) carry set
	"ne", // not equal
	"eq", // equal
	"vc", // overflow clear
	"vs", // overflow set
	"pl", // plus
	"mi", // minus
	"ge", // greater or equal
	"lt", // less than
	"gt", // greater than
	"le", // less or equal
}

// RunState specifies the current running state of the Processor.
type RunState uint32

const (
	// RunStateStopped indicates that the Processor is no longer incrementing the
	// program counter and executing instructions.
	RunStateStopped RunState = iota

	// RunStateRunning indicates that the Processor is currently executing
	// instructions until interrupted.
	RunStateRunning
)

// A Processor emulates the Motorola 68000 microprocessor.
type Processor struct {
	D  [8]uint32        // Data registers
	A  [8]uint32        // Address registers
	SR uint32           // Status Register
	PC uint32           // Program Counter
	M  *m68kmem.Decoder // System memory controller

	// TraceWriter specifies where trace log output should be written to. If
	// TraceWriter is nil, no logging is performed.
	TraceWriter io.Writer

	buf      [64]byte         // general purpose buffer
	op       uint16           // current opcode
	runState RunState         // current state
	err      error            // most recent error
	handlers [256]TrapHandler // registered trap handlers
}

type TrapHandler interface {
	Trap(c *Processor, vector int) error
}
type TrapHandlerFunc func(c *Processor, vector int) error

func (f TrapHandlerFunc) HandleTrap(p *Processor, v int) error {
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
		M:           c.M,
		TraceWriter: c.TraceWriter,
		handlers:    c.handlers,
	}
	c.M.Reset()
}

// Run executes any program loaded into memory, starting from the program
// counter value, running until completion.
func (c *Processor) Run() error {
	if c.M == nil {
		return errNoProgram
	}
	c.err = nil
	c.runState = RunStateRunning
	for c.err == nil && c.runState == RunStateRunning {
		c.Step()
	}
	c.runState = RunStateStopped
	return c.err
}

// Step executes the single instruction located at the address specified by the
// program counter register.
func (c *Processor) Step() error {
	// TODO: the returned error must be a traceableError.

	// TODO: consider prefetching a whole page of instructions and measure

	// read from program pointer into op
	if _, c.err = c.M.Read(int(c.PC), c.buf[0:2]); c.err != nil {
		return c.err
	}
	c.op = uint16(c.buf[0])<<8 + uint16(c.buf[1])

	if c.op == 0 {
		// TODO: 0 is a valid instruction. Find a way to terminate programs
		// without buffer overrun.
		c.runState = RunStateStopped
		return c.err
	}

	// dispatch opcode to function
	fn := dispatch(c.op)
	if fn == nil {
		c.err = newTraceableError(c.PC, c.op, errBadOpcode)
		return c.err
	}
	t := fn(c)

	// BUG: trace data might be corrupted by an execution error.
	c.trace(t)
	if c.err != nil {
		c.err = newTraceableError(t.addr, c.op, c.err)
	}
	return c.err
}

// Stop instructs the processor to stop processing instructions.
func (c *Processor) Stop() (err error) {
	c.runState = RunStateStopped
	return
}

// print a formatted message to the configured TraceWriter.
func (c *Processor) tracef(format string, a ...interface{}) {
	if c.TraceWriter == nil {
		return
	}
	fmt.Fprintf(c.TraceWriter, format, a...)
}

// RunState returns the current run state of the Processor.
func (c *Processor) RunState() RunState {
	return c.runState
}

// print a message to the configured TraceWriter.
func (c *Processor) trace(v ...interface{}) {
	if c.TraceWriter == nil {
		return
	}
	fmt.Fprintln(c.TraceWriter, v...)
}

// testCode returns true if the given condition code is true for the current
// processor state.
func testCond(c *Processor, cc uint16) bool {
	switch cc {
	case CondTrue:
		return true

	case CondFalse:
		return false

	case CondNotEqual:
		return c.SR&StatusZero == 0

	case CondEqual:
		return c.SR&StatusZero != 0

		// TODO: implement all evaluation of all condition codes
	}
	return false
}

// readByte reads an 8-bit byte of data from the given effective address.
func (c *Processor) readByte(ea uint16) (b byte, opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = errBadAddress
		return

	case 0x00: // data register
		b = byte(c.D[reg])
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		b = byte(c.A[reg])
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		b, err = c.M.Byte(int(c.A[reg]))
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		b, err = c.M.Byte(int(c.A[reg]))
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
		b, err = c.M.Byte(int(c.A[reg]))
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.M.Sword(int(c.PC))
		if err != nil {
			return
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		b, err = c.M.Byte(addr)
		opr = fmt.Sprintf("($%X,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			b, err = c.M.Byte(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			b, err = c.M.Byte(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x04: // immediate byte
			var n uint16
			n, err = c.M.Word(int(c.PC))
			b = byte(n)
			c.PC += 2
			opr = fmt.Sprintf("#$%X", byteToInt32(b))
		}
	}
	return
}

// readWord reads a 16-bit word of data from the given effective address.
func (c *Processor) readWord(ea uint16) (n uint16, opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = errBadAddress
		return

	case 0x00: // data register
		n = uint16(c.D[reg])
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		n = uint16(c.A[reg])
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		n, err = c.M.Word(int(c.A[reg]))
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		n, err = c.M.Word(int(c.A[reg]))
		c.A[reg] += 2
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 2
		n, err = c.M.Word(int(c.A[reg]))
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.M.Sword(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		n, err = c.M.Word(addr)
		opr = fmt.Sprintf("($%X,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			n, err = c.M.Word(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			n, err = c.M.Word(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x04: // immediate word
			n, c.err = c.M.Word(int(c.PC))
			c.PC += 2
			opr = fmt.Sprintf("#$%X", wordToInt32(n))
		}
	}
	return
}

// readLong reads a 32-bit long-word of data from the given effective address.
func (c *Processor) readLong(ea uint16) (n uint32, opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = errBadAddress
		return

	case 0x00: // data register
		n = c.D[reg]
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		n = c.A[reg]
		opr = fmt.Sprintf("A%d", reg)

	case 0x02: // memory address
		n, err = c.M.Long(int(c.A[reg]))
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		n, err = c.M.Long(int(c.A[reg]))
		c.A[reg] += 4
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 4
		n, err = c.M.Long(int(c.A[reg]))
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.M.Sword(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		n, err = c.M.Long(addr)
		opr = fmt.Sprintf("($%X,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			n, err = c.M.Long(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			n, err = c.M.Long(int(addr))
			opr = fmt.Sprintf("$%X", addr)

		case 0x04: // immediate long
			n, c.err = c.M.Long(int(c.PC))
			c.PC += 4
			opr = fmt.Sprintf("#$%X", int32(n))
		}
	}
	return
}

// readBytes reads a slice of byte data from the given effective address.
func (c *Processor) readBytes(ea uint16, b []byte) (opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = errBadAddress
		return

	case 0x02: // memory address
		_, c.err = c.M.Read(int(c.A[reg]), b)
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		_, c.err = c.M.Read(int(c.A[reg]), b)
		c.A[reg] += 4
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 4
		_, c.err = c.M.Read(int(c.A[reg]), b)
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.M.Sword(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		_, c.err = c.M.Read(addr, b)
		opr = fmt.Sprintf("($%X,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			_, c.err = c.M.Read(int(addr), b)
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			_, c.err = c.M.Read(int(addr), b)
			opr = fmt.Sprintf("$%X", addr)
		}
	}
	return
}

// readImmByte reads an 8-bit byte of immediate data from the current program
// counter address and increments the counter by 2 (16-bits) for word alignment.
func (c *Processor) readImmByte() (b byte, opr string, err error) {
	var n uint16
	n, err = c.M.Word(int(c.PC))
	if err != nil {
		return
	}
	b = byte(n) // are bytes signed???
	c.PC += 2
	opr = fmt.Sprintf("#$%X", byteToInt32(b))
	return
}

// readImmWord reads a 16-bit word of immediate data from the current program
// counter address and increments the counter by 2 (16-bits).
func (c *Processor) readImmWord() (n uint16, opr string, err error) {
	n, err = c.M.Word(int(c.PC))
	if err != nil {
		return
	}
	c.PC += 2
	opr = fmt.Sprintf("#$%X", wordToInt32(n))
	return
}

// readImmLong reads a 32-bit long-word of immediate data from the current
// program counter address and increments the counter by 4 (32-bits).
func (c *Processor) readImmLong() (n uint32, opr string, err error) {
	n, err = c.M.Long(int(c.PC))
	if err != nil {
		return
	}
	c.PC += 4
	opr = fmt.Sprintf("#$%X", int32(n))
	return
}

func (c *Processor) writeByte(ea uint16, b byte) (opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = errBadAddress
		return

	case 0x00: // data register
		c.D[reg] = c.D[reg]&0xFFFFFF00 | uint32(b)
		opr = fmt.Sprintf("D%d", reg)

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
		d, err = c.M.Sword(int(c.PC))
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
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			_, err = c.M.Write(int(addr), []byte{b})
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
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
		err = errBadAddress
		return

	case 0x00: // data register
		c.D[reg] = c.D[reg]&0xFFFF0000 | uint32(v)
		opr = fmt.Sprintf("D%d", reg)

	case 0x01: // address register
		c.A[reg] = uint32(wordToInt32(v))
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
		d, err = c.M.Sword(int(c.PC))
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
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			binary.BigEndian.PutUint16(c.buf[:2], v)
			_, err = c.M.Write(int(addr), c.buf[:2])
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
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
		err = errBadAddress
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
		d, err = c.M.Sword(int(c.PC))
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
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			binary.BigEndian.PutUint32(c.buf[:4], v)
			_, err = c.M.Write(int(addr), c.buf[:4])
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
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

func (c *Processor) writeBytes(ea uint16, b []byte) (opr string, err error) {
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		err = errBadAddress
		return

	case 0x02: // memory address
		_, err = c.M.Write(int(c.A[reg]), b)
		opr = fmt.Sprintf("(A%d)", reg)

	case 0x03: // memory address with post-increment
		_, err = c.M.Write(int(c.A[reg]), b)
		c.A[reg] += 4
		opr = fmt.Sprintf("(A%d)+", reg)

	case 0x04: // memory address with pre-decrement
		c.A[reg] -= 4
		_, err = c.M.Write(int(c.A[reg]), b)
		opr = fmt.Sprintf("-(A%d)", reg)

	case 0x05: // memory address with displacement
		var d int16
		d, err = c.M.Sword(int(c.PC))
		if err != nil {
			break
		}
		c.PC += 2
		addr := int(c.A[reg]) + int(d)
		_, err = c.M.Write(addr, b)
		opr = fmt.Sprintf("(%d,A%d)", d, reg)

	case 0x07: // other
		switch reg {
		default:
			err = errBadAddress
			return

		case 0x00: // absolute word
			var addr uint16
			addr, err = c.M.Word(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 2
			_, err = c.M.Write(int(addr), b)
			opr = fmt.Sprintf("$%X", addr)

		case 0x01: // absolute long
			var addr uint32
			addr, err = c.M.Long(int(c.PC))
			if err != nil {
				return
			}
			c.PC += 4
			_, err = c.M.Write(int(addr), b)
			opr = fmt.Sprintf("$%X", addr)
		}
	}
	return
}

func (c *Processor) setStatusForByte(b byte) {
	if b == 0 {
		c.SR |= StatusZero
	}
	if b&0x80 != 0 {
		c.SR |= StatusNegative
	}
}

func (c *Processor) setStatusForWord(v uint16) {
	if v == 0 {
		c.SR |= StatusZero
	}
	if v&0x8000 != 0 {
		c.SR |= StatusNegative
	}
}

func (c *Processor) setStatusForLong(v uint32) {
	if v == 0 {
		c.SR |= StatusZero
	}
	if v&0x80000000 != 0 {
		c.SR |= StatusNegative
	}
}
