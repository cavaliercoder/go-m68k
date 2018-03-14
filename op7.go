package m68k

import "fmt"

// opMoveq implements MOVEQ (pg. 4-134)
func opMoveq(c *Processor) (t *stepTrace) {
	b := byte(c.op)
	reg := c.op & 0x0E00 >> 9
	t = &stepTrace{
		op:   "moveq",
		addr: c.PC,
		n:    1,
		src:  fmt.Sprintf("#$%X", b),
		dst:  fmt.Sprintf("D%d", reg),
		sz:   noSize,
	}
	c.PC += 2
	c.SR &= 0xFFFFFFF0
	c.D[reg] = uint32(byteToInt32(b))
	if b == 0 {
		c.SR |= StatusZero
	}
	if b&0x80 != 0 {
		c.SR |= StatusNegative
	}
	return
}
