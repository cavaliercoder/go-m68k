package m68k

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrNoProgram = errors.New("No program loaded or memory device attached")
)

// A Processor emulates the Motorola 68000 microprocessor.
type Processor struct {
	D   [8]uint32 // Data registers
	A   [8]uint32 // Address registers
	CCR uint32    // Condition Code Register
	PC  uint32    // Program Counter
	M   Memory    // System memory controller

	// TraceWriter specifies where trace log output should be written to. If
	// TraceWriter is nil, no logging is performed.
	TraceWriter io.Writer

	buf [64]byte // general purpose buffer
	op  uint16
	err error
}

// Reset resets all processor and memory state.
func (c *Processor) Reset() {
	*c = Processor{
		TraceWriter: c.TraceWriter,
	}
	c.Jump(0x1000)
	if c.M == nil {
		c.M = NewRAM(0x4000) // 16KB
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
		if err == io.ErrShortWrite {
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
// counter value in A7.
func (c *Processor) Run() error {
	if c.M == nil {
		return ErrNoProgram
	}
	for c.err == nil {
		// read from program pointer into op
		if _, c.err = c.M.Read(int(c.PC), c.buf[0:2]); c.err != nil {
			return c.err
		}
		c.op = uint16(c.buf[0])<<8 + uint16(c.buf[1])
		if c.op == 0 {
			// this might not always be true...
			c.err = io.EOF
			return c.err
		}

		switch c.op & 0xF000 {
		case 0x0000:
			c.opcode00()
		case 0x2000:
			c.opcode20()
		case 0xD000:
			c.opcodeD0()
		default:
			c.err = fmt.Errorf("Unregistered opcode @0x%04X: 0x%04X", c.PC, c.op&0xF000)
			c.PC += 2
		}
	}
	return c.err
}

func (c *Processor) read(addr uint32, b []byte) {
	_, c.err = c.M.Read(int(addr), b)
	// TODO: create processor exception for IO errors
}

func (c *Processor) trace(format string, a ...interface{}) {
	if c.TraceWriter == nil {
		return
	}
	fmt.Fprintf(c.TraceWriter, format, a...)
}

func (c *Processor) opcode00() {
	addr := c.PC
	c.PC += 2
	switch c.op & 0x0F00 {
	case 0x0600: // addi
		// compute immediate value size
		szid := (c.op & 0x00C0) >> 6
		sz := []uint32{2, 2, 4}[szid]
		mask := []uint32{0xFF, 0xFFFF, 0xFFFFFFFF}[szid]

		// read immediate value
		b := make([]byte, sz)
		c.read(c.PC, b)

		// decode value
		v := uint32(b[0]) // immediate value
		for i := uint32(1); i < sz; i++ {
			v <<= 8
			v += uint32(b[i])
		}

		// compute destination
		d := c.op & 0x0007
		switch c.op & 0x0038 {
		case 0x00: // data register direct
			c.D[d] = (c.D[d] + v) & mask // add
			c.trace("%04X addi.%c #$%04X,D%d\n", addr, []byte{'b', 'w', 'l'}[szid], v, d)

		case 0x02: // address register indirect
			// ...
		}
		c.PC += sz

	default:
		c.trace("%04X Unknown opcode: %04X (%0X)\n", addr, c.op, c.op&0xF000)
	}
}

// opcode 0x20 is always MOVE.L. See section 4-116.
func (c *Processor) opcode20() {
	addr := c.PC
	c.PC += 2

	dm := (c.op & 0x01C0) >> 6 // dest mode
	dr := (c.op & 0x0E00) >> 9 // dest register
	sm := (c.op & 0x0038) >> 3 // source mode
	sr := c.op & 0x0007        // source register
	var v uint32
	switch sm {
	case 0x00: // data register
		v = c.D[sr]

	case 0x07:
		switch sr {
		case 0x04: // immediate
			v = c.imm(0x02)
		}

	default:
		panic("TODO")
	}

	switch dm {
	case 0x00: // data register
		c.D[dr] = v

	case 0x01: // address register
		c.A[dr] = v

	default:
		panic(dm)
	}

	c.trace("%04X move.l TODO\n", addr)
}

func (c *Processor) opcodeD0() {
	addr := c.PC
	c.PC += 2

	mod := (c.op & 0xFF) >> 4
	reg := c.op & 0x0007
	if mod == 0x07 {
		// now mode is:
		// 0x07 absolute word
		// 0x08 absolute long
		// 0x0B immediate
		mod += reg
	}

	// opmode := c.op[1] & 0xE0

	switch mod {
	case 0x00: // Direct data register
		c.trace("D%d", reg)
	case 0x01: // Direct address register
		c.trace("A%d", reg)
	case 0x02: // Indirect register address
		c.trace("(A%d)", reg)
	case 0x03: // Address with postincrement
		c.trace("(A%d)+", reg)
	case 0x04: // Address with predecrement
		c.trace("-(A%d)", reg)
	case 0x05: // Address with displacement
		c.trace("(d₁₆,A%d)", reg)
	case 0x06: // Address Register Indirect with Index
		c.trace("(d₈,A%d,Xn)", reg) // TODO: 8-bit vs. Base displacement
	case 0x07: // absolute word
		c.trace("(xxx).w")
	case 0x08: // absolute long
		c.trace("(xxx).l")
	case 0x0B: // immediate
		v := c.imm((c.op & 0x00C0) >> 6)
		m := []uint32{0xFF, 0xFFFF, 0xFFFFFFFF}[(c.op&0x00C0)>>6]
		d := c.op & 0x0E00
		c.D[d] += v
		c.D[d] &= m

		c.trace("%04X add.w #$%04X,D%d\n", addr, v, d)
	}

	// c.trace("ADD? %s - %d\n", addrMode(c.op[1]), mod)
}

// return an immediate value of size n where n is one of:
//     0 - byte
//     1 - word
//     2 - long
func (c *Processor) imm(n uint16) uint32 {
	sz := []uint32{1, 2, 4}[n]
	if _, c.err = c.M.Read(int(c.PC), c.buf[:sz]); c.err != nil {
		return 0 // TODO: handle error
	}
	c.PC += sz
	v := uint32(c.buf[0])
	for i := uint32(1); i < sz; i++ {
		v <<= 8
		v += uint32(c.buf[i])
	}
	return v
}
