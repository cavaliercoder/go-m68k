package m68k

import (
	"fmt"
	"io"
)

// op41 routes for:
// - LEA (pg. 4-110)
// - CHK (pg. 4-69)
func op41(c *Processor) (t *stepTrace) {
	// ensure bits 7 - 8 are set
	if c.op&0x0180 != 0x0180 {
		c.err = newOpcodeError(c.op)
		return
	}
	if c.op&0x40 == 0 {
		// TODO: CHK
		c.err = newOpcodeError(c.op)
		return
	}

	// LEA
	return opLea(c)
}

// op4E routes for:
// - STOP (pg.6-85)
func op4E(c *Processor) (t *stepTrace) {
	if c.op == 0x4E72 {
		return opStop(c)
	}
	c.err = newOpcodeError(c.op)
	return
}

// opLea implements LEA (pg. 4-110)
func opLea(c *Processor) (t *stepTrace) {
	an := (c.op & 0x0E00) >> 9
	t = &stepTrace{
		addr: c.PC,
		op:   "lea",
		sz:   noSize,
		dst:  fmt.Sprintf("A%d", an),
	}
	c.PC += 2

	ea := decodeEA(c.op)
	mod := ea & 0x38 >> 3
	reg := ea & 0x07
	switch mod {
	default:
		c.err = newOpcodeError(c.op)
		return

	case 0x02: // memory address
		c.A[an] = c.A[reg]
		t.src = fmt.Sprintf("(A%d)", reg)

	case 0x07: // other
		switch reg {
		default:
			c.err = newOpcodeError(c.op)
			return

		case 0x00: // absolute word
			var n uint16
			n, _, c.err = c.readImmWord()
			c.A[an] = uint32(wordToInt32(n))
			t.src = fmt.Sprintf("$%X", int32(n))

		case 0x01: // absolute long
			var n uint32
			n, _, c.err = c.readImmLong()
			c.A[an] = n
			t.src = fmt.Sprintf("$%X", int32(n))

		case 0x02: // program counter with displacement
			var n uint16
			n, _, c.err = c.readImmWord()
			disp := wordToInt32(n)
			c.A[an] = uint32(int32(c.PC) + disp - 2)
			t.src = fmt.Sprintf("$%X(PC)", disp)
		}
	}

	return
}

// opStop implements STOP (6-85)
func opStop(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "stop",
		sz:   noSize,
	}
	c.PC += 2

	var n uint16
	n, c.err = c.Word(int(c.PC))
	if c.err != nil {
		return
	}
	c.PC += 2
	t.dst = fmt.Sprintf("#$%X", n)
	c.SR = uint32(n) & StatusMask
	if c.SR&0x0700 == 0x0700 {
		c.err = io.EOF // end program if interrupt mask is maximum
	}
	return
}
