package m68k

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

// opClr implements CLR (4-73).
// Clear operand.
func opClr(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		n:    1,
		op:   "clr",
		sz:   c.op >> 6 & 0x3,
	}
	c.PC += 2

	switch t.sz {
	default:
		c.err = errBadOpSize

	case SizeByte:
		t.dst, c.err = c.writeByte(c.op, 0)

	case SizeWord:
		t.dst, c.err = c.writeWord(c.op, 0)

	case SizeLong:
		t.dst, c.err = c.writeLong(c.op, 0)
	}
	if c.err != nil {
		return
	}
	c.SR &= 0xFFFFFFF0
	c.SR |= StatusZero
	return
}

// opMoveToSr implements MOVE to SR (6-19).
func opMoveToSr(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		n:    1,
		op:   "move",
		sz:   noSize,
		dst:  "SR",
	}
	c.PC += 2

	var n uint16
	n, t.src, c.err = c.readWord(c.op)
	if c.err != nil {
		return
	}

	c.SR = uint32(n & StatusMask)
	return
}

// opMovem implements MOVEM (pg. 4-128)
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

	if dir == 0 {
		// register to memory
		t.src = fmt.Sprintf("$%X", regl)
		b := &bytes.Buffer{}
		for i := uint16(0); i < 16; i++ {
			ok := regl & (1 << i)
			if ok == 0 {
				continue // skip unset registers
			}

			// read register value
			var v uint32
			if i < 8 {
				v = c.A[7-i]
			} else {
				v = c.D[7-(i-8)]
			}

			// write value to buffer
			binary.BigEndian.PutUint32(c.buf[:4], v)
			if t.sz == SizeWord { // movem.w <list>,<ea>
				b.Write(c.buf[2:4])
			} else { // movem.l <list>,<ea>
				b.Write(c.buf[:4])
			}
		}

		// flush buffer to memory
		t.dst, c.err = c.writeBytes(c.op, b.Bytes())
		return
	}

	// memory to register
	t.dst = fmt.Sprintf("$%X", regl)

	// count registers
	regc := 0
	for i := uint16(0); i < 16; i++ {
		if regl&(1<<i) != 0 {
			regc++
		}
	}

	// read values from memory
	buflen := regc * 2
	if t.sz == SizeLong {
		buflen = regc * 4
	}
	t.src, c.err = c.readBytes(c.op, c.buf[:buflen])
	if c.err != nil {
		return
	}

	// set register values
	x := 0
	for i := uint16(0); i < 16; i++ {
		ok := regl & (1 << i)
		if ok == 0 {
			continue // skip unset registers
		}
		var v uint32
		if t.sz == SizeWord { // movem.w <ea>,<list>
			v = uint32(wordToInt32(binary.BigEndian.Uint16(c.buf[x : x+2])))
			x += 2
		} else { // movem.l <ea>,<list>
			v = binary.BigEndian.Uint32(c.buf[x : x+4])
			x += 4
		}
		if i < 8 {
			c.D[i] = v
		} else {
			c.A[i-8] = v
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
	// TODO: handle case where interrupt mask is maximum (0x0700)
	c.err = c.Stop()
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
	c.PC, c.err = c.M.Long(int(c.A[7]))
	if c.err != nil {
		return
	}
	c.A[7] += 4
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

	// compute jump address and mutate PC to return address
	var jmp uint32
	mod := c.op & 0x38 >> 3
	reg := c.op & 0x07
	switch mod {
	// TODO: implement remaining address modes for JSR
	default:
		c.err = errBadAddress
		return

	case 0x02: // address register indirect
		jmp = c.A[reg]
		t.src = fmt.Sprintf("(A%d)", reg)

	case 0x07: // other
		switch reg {
		default:
			c.err = errBadAddress
			return

		case 0x00: // absolute word
			var n uint16
			n, _, c.err = c.readImmWord()
			jmp = uint32(wordToInt32(n))
			t.src = fmt.Sprintf("$%X", int32(n))

		case 0x01: // absolute long
			var n uint32
			n, _, c.err = c.readImmLong()
			jmp = n
			t.src = fmt.Sprintf("$%X", int32(n))

		case 0x02: // program counter with displacement
			var n uint16
			n, _, c.err = c.readImmWord()
			disp := wordToInt32(n)
			jmp = uint32(int32(c.PC) + disp - 2)
			t.src = fmt.Sprintf("($%X,PC)", disp)
		}
	}

	// push return address (current PC) to stack
	c.A[7] -= 4
	binary.BigEndian.PutUint32(c.buf[:4], c.PC)
	_, c.err = c.M.Write(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}

	// jump
	c.PC = jmp
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
	c.SR &= 0xFFFFFFF0 // reset last 4 bits only
	switch t.sz {
	case SizeByte:
		var b byte
		b, t.src, c.err = c.readByte(c.op)
		c.setStatusForByte(b)

	case SizeWord:
		var n uint16
		n, t.src, c.err = c.readWord(c.op)
		c.setStatusForWord(n)

	case SizeLong:
		var n uint32
		n, t.src, c.err = c.readLong(c.op)
		c.setStatusForLong(n)
	}
	return
}
