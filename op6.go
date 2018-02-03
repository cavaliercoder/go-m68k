package m68k

import "fmt"

// opBcc implements:
// - BRA
// - BSR
// - Bcc (4-25)
func opBcc(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		sz:   noSize,
		n:    1,
	}
	c.PC += 2

	cc := c.op & 0x0F00 >> 8
	t.op = []string{
		"bra", // bra
		"bsr", // bsr (this should never happen)
		"bhi", // high
		"bls", // low or same
		"bcc", // (hi) carry clear
		"bcs", // (lo) carry set
		"bne", // not equal
		"beq", // equal
		"bvc", // overflow clear
		"bvs", // overflow set
		"bpl", // plus
		"bmi", // minus
		"bge", // greater or equal
		"blt", // less than
		"bgt", // greater than
		"ble", // less or equal
	}[cc]

	// compute displacement
	var disp int32
	switch c.op & 0xFF {
	default: // 8-bit displacement
		disp = byteToInt32(byte(c.op))

	case 0: // 16-bit displacement
		var n uint16
		n, _, c.err = c.readImmWord()
		disp = wordToInt32(n)

	case 0xFF: // 32-bit displacement
		if cc < 2 {
			// 32-bit displacement is not implemented on the 68000 for bra and bsr
			c.err = newOpcodeError(c.op)
			return
		}
		var n uint32
		n, _, c.err = c.readImmLong()
		disp = longToInt32(n)
	}
	if c.err != nil {
		return
	}
	t.src = fmt.Sprintf("$%X(PC)", disp)

	// decide if to branch
	switch cc {
	default:
		c.err = newOpcodeError(c.op)
		return

	case 0: // bra
		// always branch

	case CondNotEqual: // bne
		if c.SR&StatusZero != 0 {
			return
		}
	}

	// branch
	c.PC = uint32((int32(t.addr+2) + disp))
	return
}
