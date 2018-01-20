package m68k

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrNoProgram = errors.New("No program loaded or memory device attached")
)

// A Processor emulate the Motorola 68000 microprocessor.
type Processor struct {
	D   [8]uint32 // Data registers
	A   [8]uint32 // Address registers
	CCR uint32    // Condition Code Register
	M   Memory

	TraceWriter io.Writer

	op  [16]byte
	err error
}

func (c *Processor) Reset() {
	for i := 0; i < len(c.D); i++ {
		c.D[i] = 0
		c.A[i] = 0
	}
	c.CCR = 0
	c.Jump(0x1000)
	if c.M == nil {
		c.M = NewRAM(0x4000) // 16KB
	}
	// TODO: c.Memory.Reset()
}

// Jump sets the value of the program counter to the given memory address.
func (c *Processor) Jump(addr uint32) error {
	c.A[7] = uint32(addr)
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
		if _, c.err = c.M.Read(int(c.A[7]), c.op[0:2]); c.err != nil {
			return c.err
		}
		if c.op[0] == 0 && c.op[1] == 0 {
			// this might not always be true...
			c.err = io.EOF
			return c.err
		}

		switch c.op[0] & 0xF0 {
		case 0x00:
			c.err = c.opcode00()
		case 0xD0:
			c.err = c.opcodeD0()
		default:
			c.trace("%04X Opcode: %04X (%0X)\n", c.A[7], c.op[:2], c.op[0]&0xF0)
			c.A[7] += 2
		}
	}
	return c.err
}

func (c *Processor) trace(format string, a ...interface{}) {
	if c.TraceWriter == nil {
		return
	}
	fmt.Fprintf(c.TraceWriter, format, a...)
}

func (c *Processor) opcode00() error {
	addr := c.A[7]
	c.A[7] += 2
	switch c.op[0] & 0x0F {
	case 0x06: // addi
		// compute immediate value size
		szid := (c.op[1] & 0xC0) >> 6
		sz := []uint32{2, 2, 4}[szid]
		mask := []uint32{0xFF, 0xFFFF, 0xFFFFFFFF}[szid]

		// read immediate value
		b := make([]byte, sz)
		if _, err := c.M.Read(int(c.A[7]), b); err != nil {
			return err
		}

		// decode value
		v := uint32(b[0]) // immediate value
		for i := uint32(1); i < sz; i++ {
			v <<= 8
			v += uint32(b[i])
		}

		// compute destination register
		d := c.op[1] & 0x07 // destination data register

		// assign value
		c.D[d] = (c.D[d] + v) & mask // add

		szs := []byte{'b', 'w', 'l'}
		c.trace("%04X addi.%c #$%04X,D%d\n", addr, szs[szid], v, d)
		c.A[7] += sz

	default:
		c.trace("%04X Unknown opcode: %04X (%0X)\n", addr, c.op[:2], c.op[0]&0xF0)
	}
	return nil
}

func (c *Processor) opcodeD0() error {
	addr := c.A[7]
	c.A[7] += 2

	mod := c.op[1] >> 4
	reg := c.op[1] & 0x07
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
		sz := []uint32{1, 2, 4}[(c.op[1]&0xC0)>>6]
		mask := []uint32{0xFF, 0xFFFF, 0xFFFFFFFF}[(c.op[1]&0xC0)>>6]
		b := make([]byte, sz)
		if _, err := c.M.Read(int(c.A[7]), b); err != nil {
			return err
		}
		c.A[7] += sz
		v := uint32(b[0])
		for i := uint32(1); i < sz; i++ {
			v <<= 8
			v += uint32(b[i])
		}
		d := c.op[0] & 0x0E
		c.D[d] += v
		c.D[d] &= mask

		c.trace("%04X add.w #$%04X,D%d\n", addr, v, d)
	}

	// c.trace("ADD? %s - %d\n", addrMode(c.op[1]), mod)
	return nil
}
