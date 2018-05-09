package m68k

// opCmp implements CMP (pg. 4-75)
func opCmp(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "cmp",
		n:    1,
		sz:   c.op & 0x1c >> 6,
	}
	c.PC += 2

	ea := decodeEA(c.op)
	reg := c.op & 0x0E00 >> 9

	var a, b Int
	a = NewInt(c.D[reg], int(t.sz))
	b, t.src, t.err = c.readInt(ea, t.sz)

	a = a.Sub(b)
	c.SR &= 0xFFFFFFF0
	c.SR |= a.Status()
	return
}
