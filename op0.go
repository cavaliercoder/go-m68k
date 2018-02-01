package m68k

// This file contains opcodes starting in 0x0000

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
	if src&0x38 == 0x08 && dst&0x20 == 0x20 {
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
		b, t.src, c.err = c.readByte(src)
		if c.err != nil {
			break
		}
		t.dst, c.err = c.writeByte(dst, b)
		if b == 0 {
			c.SR |= StatusZero
		}

	case SizeWord:
		var v uint16
		v, t.src, c.err = c.readWord(src)
		if c.err != nil {
			break
		}
		t.dst, c.err = c.writeWord(dst, v)
		if v == 0 {
			c.SR |= StatusZero
		}

	case SizeLong:
		var v uint32
		v, t.src, c.err = c.readLong(src)
		if c.err != nil {
			break
		}
		t.dst, c.err = c.writeLong(dst, v)
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
		c.err = newOpcodeError(c.op)
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, c.err = c.readImmByte()
		if c.err != nil {
			break
		}
		if ea == 0x3C { // ori.b to CCR
			c.SR |= uint32(src) & StatusMask
			t.sz = noSize
			t.dst = "CCR"
			break
		}

		// ori.b
		dst, _, c.err = c.readByte(ea)
		if c.err != nil {
			break
		}
		b := dst | src
		t.dst, c.err = c.writeByte(ea, b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, c.err = c.readImmWord()
		if c.err != nil {
			break
		}
		if ea == 0x3C { // ori.w to SR
			c.SR |= uint32(src) & StatusMask
			t.sz = noSize
			t.dst = "SR"
			break
		}

		// ori.w
		dst, _, c.err = c.readWord(ea)
		if c.err != nil {
			break
		}
		v := dst | src
		t.dst, c.err = c.writeWord(ea, v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, c.err = c.readImmLong()
		if c.err != nil {
			break
		}
		dst, _, c.err = c.readLong(ea)
		if c.err != nil {
			break
		}
		v := dst | src
		t.dst, c.err = c.writeLong(ea, v)
	}
	return
}

// opAndi implements ANDI, ANDI to CCR and ANDI to SR
func opAndi(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "andi",
		n:    1,
		sz:   (c.op & 0xC0) >> 6,
	}
	c.PC += 2
	ea := decodeEA(c.op)

	switch t.sz {
	default:
		c.err = newOpcodeError(c.op)
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, c.err = c.readImmByte()

		if ea == 0x3C { // and.b to CCR
			c.SR |= uint32(src) & StatusMask
			t.sz = noSize
			t.dst = "CCR"
			break
		}

		// and.b
		dst, _, c.err = c.readByte(ea)
		if c.err != nil {
			break
		}
		b := dst & src
		t.dst, c.err = c.writeByte(ea, b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, c.err = c.readImmWord()

		if ea == 0x3C { // andi.w to SR
			c.SR |= uint32(src) & StatusMask
			t.sz = noSize
			t.dst = "SR"
			break
		}

		// ori.w
		dst, _, c.err = c.readWord(ea)
		if c.err != nil {
			break
		}
		v := dst & src
		t.dst, c.err = c.writeWord(ea, v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, c.err = c.readImmLong()

		// andi.l
		dst, _, c.err = c.readLong(ea)
		if c.err != nil {
			break
		}
		v := dst & src
		t.dst, c.err = c.writeLong(ea, v)
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
		c.err = newOpcodeError(c.op)
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, c.err = c.readImmByte()
		if c.err != nil {
			break
		}
		dst, _, c.err = c.readByte(ea)
		if c.err != nil {
			break
		}
		dst -= src
		t.dst, c.err = c.writeByte(ea, dst)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, c.err = c.readImmWord()
		if c.err != nil {
			break
		}
		dst, _, c.err = c.readWord(ea)
		if c.err != nil {
			break
		}
		dst -= src
		t.dst, c.err = c.writeWord(ea, dst)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, c.err = c.readImmLong()
		if c.err != nil {
			break
		}
		dst, _, c.err = c.readLong(ea)
		if c.err != nil {
			break
		}
		dst -= src
		t.dst, c.err = c.writeLong(ea, dst)
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
		c.err = newOpcodeError(c.op)
		return
	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, c.err = c.readImmByte()
		if c.err != nil {
			break
		}
		dst, _, c.err = c.readByte(ea)
		if c.err != nil {
			break
		}
		b := dst + src
		t.dst, c.err = c.writeByte(ea, b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, c.err = c.readImmWord()
		if c.err != nil {
			break
		}
		dst, _, c.err = c.readWord(ea)
		if c.err != nil {
			break
		}
		v := dst + src
		t.dst, c.err = c.writeWord(ea, v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, c.err = c.readImmLong()
		if c.err != nil {
			break
		}
		dst, _, c.err = c.readLong(ea)
		if c.err != nil {
			break
		}
		v := dst + src
		t.dst, c.err = c.writeLong(ea, v)
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
		c.err = newOpcodeError(c.op)
		return

	case SizeByte:
		// read immediate byte
		var src, dst byte
		src, t.src, c.err = c.readImmByte()
		if c.err != nil {
			break
		}

		if ea == 0x3C { // eori.b to CCR
			c.SR = (c.SR ^ uint32(src)) & StatusMask
			t.sz = noSize
			t.dst = "CCR"
			break
		}

		// eori.b
		dst, _, c.err = c.readByte(ea)
		if c.err != nil {
			break
		}
		b := dst ^ src
		t.dst, c.err = c.writeByte(ea, b)

	case SizeWord:
		// read immediate word
		var src, dst uint16
		src, t.src, c.err = c.readImmWord()
		if c.err != nil {
			break
		}

		if ea == 0x3C { // eori.w to SR
			c.SR = (c.SR ^ uint32(src)) & StatusMask
			t.sz = noSize
			t.dst = "SR"
			break
		}

		// eori.w
		dst, _, c.err = c.readWord(ea)
		if c.err != nil {
			break
		}
		v := dst ^ src
		t.dst, c.err = c.writeWord(ea, v)

	case SizeLong:
		// read immediate long
		var src, dst uint32
		src, t.src, c.err = c.readImmLong()
		if c.err != nil {
			break
		}

		// eori.l
		dst, _, c.err = c.readLong(ea)
		if c.err != nil {
			break
		}
		v := dst ^ src
		t.dst, c.err = c.writeLong(ea, v)
	}
	return
}
