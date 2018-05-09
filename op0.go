package m68k

import "fmt"

// opMove implements MOVE, MOVEA and MOVEP
func opMove(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "move",
		n:    1,

		// decode unusual size encoding for move instructions
		sz: []uint16{0xFFFF, SizeByte, SizeLong, SizeWord}[(c.op&0x3000)>>12],
	}
	c.PC += 2

	src, dst := decodeEA(c.op), decodeExEA(c.op)
	if src&0x38 == 0x08 && dst&0x40 == 0x40 {
		t.op = "movep"
	}
	if dst&0x38 == 0x08 {
		t.op = "movea"
		// TODO: sign extension and drop support for bytes
	}

	c.SR &= 0xFFFFFFF0

	switch t.sz {
	case SizeByte:
		var b byte
		b, t.src, t.err = c.readByte(src)
		if t.err != nil {
			break
		}
		t.dst, t.err = c.writeByte(dst, b)
		if b == 0 {
			c.SR |= StatusZero
		}

	case SizeWord:
		var v uint16
		v, t.src, t.err = c.readWord(src)
		if t.err != nil {
			break
		}
		t.dst, t.err = c.writeWord(dst, v)
		if v == 0 {
			c.SR |= StatusZero
		}

	case SizeLong:
		var v uint32
		v, t.src, t.err = c.readLong(src)
		if t.err != nil {
			break
		}
		t.dst, t.err = c.writeLong(dst, v)
		if v == 0 {
			c.SR |= StatusZero
		}
	}

	return
}

// opOri implements ORI, ORI to CCR and ORI to SR.
func opOri(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "ori",
		n:    1,
		sz:   (c.op & 0xC0) >> 6,
	}
	c.PC += 2
	ea := decodeEA(c.op)

	switch t.sz {
	default:
		t.err = errBadOpSize
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, t.err = c.readImmByte()
		if t.err != nil {
			break
		}
		if ea == 0x3C { // ori.b to CCR
			c.SR |= uint32(src) & StatusMask
			t.sz = noSize
			t.dst = "CCR"
			break
		}

		// ori.b
		dst, _, t.err = c.readByte(ea)
		if t.err != nil {
			break
		}
		b := dst | src
		t.dst, t.err = c.writeByte(ea, b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, t.err = c.readImmWord()
		if t.err != nil {
			break
		}
		if ea == 0x3C { // ori.w to SR
			c.SR |= uint32(src) & StatusMask
			t.sz = noSize
			t.dst = "SR"
			break
		}

		// ori.w
		dst, _, t.err = c.readWord(ea)
		if t.err != nil {
			break
		}
		v := dst | src
		t.dst, t.err = c.writeWord(ea, v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, t.err = c.readImmLong()
		if t.err != nil {
			break
		}
		dst, _, t.err = c.readLong(ea)
		if t.err != nil {
			break
		}
		v := dst | src
		t.dst, t.err = c.writeLong(ea, v)
	}
	return
}

// opAndi implements ANDI (4-18), ANDI to CCR and ANDI to SR
func opAndi(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "andi",
		n:    1,
		sz:   (c.op & 0xC0) >> 6,
	}
	c.PC += 2
	ea := decodeEA(c.op)
	sr := c.SR         // cache value in case destination is status register
	c.SR &= 0xFFFFFFF0 // reset last 4 bits only
	switch t.sz {
	default:
		t.err = errBadOpSize
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, t.err = c.readImmByte()

		if ea == 0x3C { // and.b to CCR
			c.SR = sr | uint32(src)&StatusMask
			t.sz = noSize
			t.dst = "CCR"
			break
		}

		// and.b
		dst, _, t.err = c.readByte(ea)
		if t.err != nil {
			break
		}
		b := dst & src
		t.dst, t.err = c.writeByte(ea, b)
		c.setStatusForByte(b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, t.err = c.readImmWord()

		if ea == 0x3C { // andi.w to SR
			c.SR = sr | uint32(src)&StatusMask
			t.sz = noSize
			t.dst = "SR"
			break
		}

		// ori.w
		dst, _, t.err = c.readWord(ea)
		if t.err != nil {
			break
		}
		v := dst & src
		t.dst, t.err = c.writeWord(ea, v)
		c.setStatusForWord(v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, t.err = c.readImmLong()

		// andi.l
		dst, _, t.err = c.readLong(ea)
		if t.err != nil {
			break
		}
		v := dst & src
		t.dst, t.err = c.writeLong(ea, v)
		c.setStatusForLong(v)
	}

	return
}

// opSubi implements SUBI
func opSubi(c *Processor) (t *stepTrace) {
	// TODO: eliminate illegal address schemes
	t = &stepTrace{
		addr: c.PC,
		op:   "subi",
		n:    1,
		sz:   (c.op & 0xC0) >> 6,
	}
	c.PC += 2
	ea := decodeEA(c.op)

	switch t.sz {
	default:
		t.err = errBadOpSize
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, t.err = c.readImmByte()
		if t.err != nil {
			break
		}
		dst, _, t.err = c.readByte(ea)
		if t.err != nil {
			break
		}
		dst -= src
		t.dst, t.err = c.writeByte(ea, dst)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, t.err = c.readImmWord()
		if t.err != nil {
			break
		}
		dst, _, t.err = c.readWord(ea)
		if t.err != nil {
			break
		}
		dst -= src
		t.dst, t.err = c.writeWord(ea, dst)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, t.err = c.readImmLong()
		if t.err != nil {
			break
		}
		dst, _, t.err = c.readLong(ea)
		if t.err != nil {
			break
		}
		dst -= src
		t.dst, t.err = c.writeLong(ea, dst)
	}
	return
}

// opAddi implements ADDI
func opAddi(c *Processor) (t *stepTrace) {
	// TODO: eliminate illegal address schemes
	t = &stepTrace{
		addr: c.PC,
		op:   "addi",
		n:    1,
		sz:   (c.op & 0xC0) >> 6,
	}
	c.PC += 2
	ea := decodeEA(c.op)

	switch t.sz {
	default:
		t.err = errBadOpSize
		return
	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, t.err = c.readImmByte()
		if t.err != nil {
			break
		}
		dst, _, t.err = c.readByte(ea)
		if t.err != nil {
			break
		}
		b := dst + src
		t.dst, t.err = c.writeByte(ea, b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, t.err = c.readImmWord()
		if t.err != nil {
			break
		}
		dst, _, t.err = c.readWord(ea)
		if t.err != nil {
			break
		}
		v := dst + src
		t.dst, t.err = c.writeWord(ea, v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, t.err = c.readImmLong()
		if t.err != nil {
			break
		}
		dst, _, t.err = c.readLong(ea)
		if t.err != nil {
			break
		}
		v := dst + src
		t.dst, t.err = c.writeLong(ea, v)
	}
	return
}

// opEori implements EORI, EORI to CCR and EORI to SR.
func opEori(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "eori",
		n:    1,
		sz:   (c.op & 0xC0) >> 6,
	}
	c.PC += 2
	ea := decodeEA(c.op)
	switch t.sz {
	default:
		t.err = errBadOpSize
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, t.err = c.readImmByte()
		if t.err != nil {
			break
		}

		if ea == 0x3C { // eori.b to CCR
			c.SR = (c.SR ^ uint32(src)) & StatusMask
			t.sz = noSize
			t.dst = "CCR"
			break
		}

		// eori.b
		dst, _, t.err = c.readByte(ea)
		if t.err != nil {
			break
		}
		b := dst ^ src
		t.dst, t.err = c.writeByte(ea, b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, t.err = c.readImmWord()
		if t.err != nil {
			break
		}

		if ea == 0x3C { // eori.w to SR
			c.SR = (c.SR ^ uint32(src)) & StatusMask
			t.sz = noSize
			t.dst = "SR"
			break
		}

		// eori.w
		dst, _, t.err = c.readWord(ea)
		if t.err != nil {
			break
		}
		v := dst ^ src
		t.dst, t.err = c.writeWord(ea, v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, t.err = c.readImmLong()
		if t.err != nil {
			break
		}

		// eori.l
		dst, _, t.err = c.readLong(ea)
		if t.err != nil {
			break
		}
		v := dst ^ src
		t.dst, t.err = c.writeLong(ea, v)
	}
	return
}

// opBtst implements BTST (4-61)
func opBtst(c *Processor) (t *stepTrace) {
	// TODO: implement bit number as immediate value (4-63)
	t = &stepTrace{
		addr: c.PC,
		op:   "btst",
		sz:   noSize,
		n:    1,
	}
	c.PC += 2

	// read bit number from data register
	src := (c.op >> 9) & 0x7
	t.src = fmt.Sprintf("D%d", src)
	bit := c.D[src]

	// read destination value
	mod := c.op & 0x38 >> 3
	var v, width uint32
	if mod == 0 { // data register
		v, t.dst, t.err = c.readLong(c.op)
		if t.err != nil {
			return
		}
		width = 32
	} else {
		var b byte
		b, t.dst, t.err = c.readByte(c.op)
		v = uint32(b)
		width = 8
	}

	// set status bit
	bit %= width
	if v&bit == 0 {
		c.SR |= StatusZero
	} else {
		c.SR &= ^StatusZero
	}
	return
}
