package m68k

import (
	"fmt"
	"io"
)

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
		c.err = errBadAddress
		return

	case 0x02: // memory address
		c.A[an] = c.A[reg]
		t.src = fmt.Sprintf("(A%d)", reg)

	case 0x07: // other
		switch reg {
		default:
			c.err = errBadAddress
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
			t.src = fmt.Sprintf("($%X,PC)", disp)
		}
	}

	return
}

// opMovem implementes MOVEM (pg. 4-128)
func opMovem(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		n:    1,
		op:   "movem",
		sz:   []uint16{SizeWord, SizeLong}[c.op&0x0040>>6],
	}
	c.PC += 2
	regl, _, _ := c.readImmWord() // register list
	dir := c.op & 0x0400 >> 10    // direction

	if dir == 0 { // register to memory
		// TODO: implement movem register to memory
		c.err = errNotImplemented
		return
	} else { // memory to register
		t.dst = fmt.Sprintf("$%X", regl)
		for i := uint16(0); i < 16; i++ {
			ok := regl & (1 << i)
			if ok == 0 {
				continue // skip unset registers
			}
			if t.sz == SizeWord { // movem.w <ea>,<list>
				var n uint16
				n, t.src, c.err = c.readWord(c.op)
				if i < 8 {
					c.D[i] = uint32(wordToInt32(n))
				} else {
					c.A[i-8] = uint32(wordToInt32(n))
				}
			} else { // movem.l <ea>,<list>
				var n uint32
				n, t.src, c.err = c.readLong(c.op)
				if i < 8 {
					c.D[i] = n
				} else {
					c.A[i-8] = n
				}
			}
		}
	}

	return
}

// opTrap implements TRAP (4-188)
func opTrap(c *Processor) (t *stepTrace) {
	v := c.op & 0x000F
	t = &stepTrace{
		addr: c.PC,
		n:    1,
		op:   "trap",
		sz:   noSize,
		src:  fmt.Sprintf("#%d", v),
	}
	c.PC += 2
	v += 32
	if c.handlers[v] != nil {
		// TODO: should trap handlers be executed inline?
		c.err = c.handlers[v].Exception(c, int(v))
	}
	return
}

// opMoveUsp implements MOVE USP (pg. 6-21)
func opMoveUsp(c *Processor) (t *stepTrace) {
	reg := c.op & 0x07
	t = &stepTrace{
		addr: c.PC,
		op:   "move",
		n:    1,
		sz:   noSize,
	}
	c.PC += 2
	dir := c.op & 0x08 >> 3
	if dir == 0 { // An to USP
		t.src, t.dst = fmt.Sprintf("A%d", reg), "USP"
		c.A[7] = c.A[reg]
	} else { // USP to An
		t.src, t.dst = "USP", fmt.Sprintf("A%d", reg)
		c.A[reg] = c.A[7]
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
	n, c.err = c.M.Word(int(c.PC))
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

// opJsr implements JSR (4-109)
func opJsr(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "jsr",
		n:    1,
		sz:   noSize,
	}
	c.PC += 2

	mod := c.op & 0x38 >> 3
	reg := c.op & 0x07
	switch mod {
	// TODO: implement remaining address modes for JSR
	default:
		c.err = errBadAddress
		return

	case 0x02: // address register
		c.PC = c.A[reg]
		t.src = fmt.Sprintf("(A%d)", reg)

	case 0x07: // other
		switch reg {
		default:
			c.err = errBadAddress
			return

		case 0x00: // absolute word
			var n uint16
			n, _, c.err = c.readImmWord()
			c.PC = uint32(wordToInt32(n))
			t.src = fmt.Sprintf("$%X", int32(n))

		case 0x01: // absolute long
			var n uint32
			n, _, c.err = c.readImmLong()
			c.PC = n
			t.src = fmt.Sprintf("$%X", int32(n))

		case 0x02: // program counter with displacement
			var n uint16
			n, _, c.err = c.readImmWord()
			disp := wordToInt32(n)
			c.PC = uint32(int32(c.PC) + disp - 2)
			t.src = fmt.Sprintf("($%X,PC)", disp)
		}
	}
	return
}

// opTst implements TST (4-192)
func opTst(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "tst",
		sz:   c.op & 0xC0 >> 6,
		n:    1,
	}
	c.PC += 2
	c.SR &= 0xFFF0
	switch t.sz {
	case SizeByte:
		var b byte
		b, t.src, c.err = c.readByte(c.op)
		if b == 0 {
			c.SR |= StatusZero
		}
		if b&0x80 != 0 {
			c.SR |= StatusNegative
		}

	case SizeWord:
		var n uint16
		n, t.src, c.err = c.readWord(c.op)
		if n == 0 {
			c.SR |= StatusZero
		}
		if n&0x8000 != 0 {
			c.SR |= StatusNegative
		}

	case SizeLong:
		var n uint32
		n, t.src, c.err = c.readLong(c.op)
		if n == 0 {
			c.SR |= StatusZero
		}
		if n&0x8000000 != 0 {
			c.SR |= StatusNegative
		}
	}
	return
}

// opRts implements RTS (Return from Subroutine) (4-169)
func opRts(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "rts",
		sz:   noSize,
		n:    1,
	}
	c.PC = c.A[7]
	c.A[7] += 4
	return
}
