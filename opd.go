package m68k

import "fmt"

// opAdd implements ADD (pg. 4-4)
func opAdd(c *Processor) (t *stepTrace) {
	sz := c.op & 0xC0 >> 6
	if sz == 3 {
		return opAdda(c)
	}
	t = &stepTrace{
		addr: c.PC,
		n:    1,
		op:   "add",
		sz:   sz,
	}
	c.PC += 2
	reg := c.op & 0x0E00 >> 9
	dir := c.op & 0x0100 >> 8

	if dir == 0 { // register to ea
		t.src = fmt.Sprintf("D%d", reg)
		switch t.sz {
		case SizeByte:
			var src, dst, v byte
			src = byte(c.D[reg])
			dst, t.dst, c.err = c.readByte(c.op)
			if c.err != nil {
				return
			}
			v = dst + src
			_, c.err = c.writeByte(c.op, v)

		case SizeWord:
			var src, dst, v uint16
			src = uint16(c.D[reg])
			dst, t.dst, c.err = c.readWord(c.op)
			if c.err != nil {
				return
			}
			v = dst + src
			_, c.err = c.writeWord(c.op, v)

		case SizeLong:
			var src, dst, v uint32
			src = c.D[reg]
			dst, t.dst, c.err = c.readLong(c.op)
			if c.err != nil {
				return
			}
			v = dst + src
			_, c.err = c.writeLong(c.op, v)
		}
	} else { // ea to register
		t.dst = fmt.Sprintf("D%d", reg)
		switch t.sz {
		case SizeByte:
			var src, dst, v byte
			dst = byte(c.D[reg])
			src, t.src, c.err = c.readByte(c.op)
			if c.err != nil {
				return
			}
			v = dst + src
			c.D[reg] = c.D[reg]&0xFFFFFF00 | uint32(v)

		case SizeWord:
			var src, dst, v uint16
			dst = uint16(c.D[reg])
			src, t.src, c.err = c.readWord(c.op)
			if c.err != nil {
				return
			}
			v = dst + src
			c.D[reg] = c.D[reg]&0xFFFF0000 | uint32(v)

		case SizeLong:
			var src uint32
			src, t.src, c.err = c.readLong(c.op)
			if c.err != nil {
				return
			}
			c.D[reg] += src
		}
	}

	return
}

// opAdda implements ADDA (pg. 4-7)
func opAdda(c *Processor) (t *stepTrace) {
	reg := c.op & 0x0E00 >> 9
	t = &stepTrace{
		addr: c.PC,
		n:    1,
		op:   "adda",
		sz:   []uint16{SizeWord, SizeLong}[c.op&0x01>>9],
		dst:  fmt.Sprintf("A%d", reg),
	}
	c.PC += 2

	switch t.sz {
	case SizeWord:
		var src, dst uint16
		dst = uint16(c.A[reg])
		src, t.src, c.err = c.readWord(c.op)
		if c.err != nil {
			return
		}
		c.A[reg] = uint32(wordToInt32(dst + src))

	case SizeLong:
		var src uint32
		src, t.src, c.err = c.readLong(c.op)
		if c.err != nil {
			return
		}
		c.A[reg] += src
	}
	return
}
