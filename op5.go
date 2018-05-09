package m68k

import "fmt"

// opDBcc implements DBcc (pg. 4-90)
// Test, decrement and branch.
func opDBcc(c *Processor) (t *stepTrace) {
	reg := c.op & 0x07
	cc := c.op & 0x0F00 >> 8
	t = &stepTrace{
		addr: c.PC,
		n:    1,
		op:   fmt.Sprintf("db%s", conditionStrings[cc]),
		src:  fmt.Sprintf("D%d", reg),
		sz:   noSize,
	}
	c.PC += 2

	// compute displacement
	var dw uint16
	dw, _, t.err = c.readImmWord()
	if t.err != nil {
		return
	}
	disp := wordToInt32(dw)
	t.dst = fmt.Sprintf("($%X,PC)", disp)

	// test condition and return if true
	if testCond(c, cc) {
		return
	}

	// decrement register
	var n uint16
	n = uint16(c.D[reg]) - 1
	c.D[reg] = c.D[reg]&0xFFFF0000 | uint32(n)

	// branch if n != -1
	if n != 0xFFFF {
		c.PC = uint32(int32(c.PC) + disp - 2)
	}
	return
}
