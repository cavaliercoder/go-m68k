package m68k

import "fmt"

// opAbcd implements ABCD (pg. 4-2)
func opAbcd(c *Processor) (t *stepTrace) {
	t = &stepTrace{
		addr: c.PC,
		op:   "abcd",
		sz:   noSize,
		n:    1,
	}
	c.PC += 2

	// read src and dst values
	var vx, vy byte
	rx, ry := c.op&0x0E>>9, c.op&0x07
	x := byte(c.SR & StatusExtend >> 4)
	mod := c.op & 0x08 >> 3
	if mod == 0 {
		// data register to data register
		t.dst = fmt.Sprintf("D%d", rx)
		t.src = fmt.Sprintf("D%d", ry)
		vx, vy = byte(c.D[rx]), byte(c.D[ry])
	} else {
		// memory to memory with pre-decrement
		t.dst = fmt.Sprintf("-(A%d)", rx)
		t.src = fmt.Sprintf("-(A%d)", ry)
		c.A[rx]--
		c.A[ry]--
		vx, c.err = c.M.Byte(int(c.A[rx]))
		if c.err != nil {
			return
		}
		vy, c.err = c.M.Byte(int(c.A[ry]))
		if c.err != nil {
			return
		}
	}

	// reset satus
	c.SR &= ^(StatusCarry | StatusExtend)

	// packed bcd add
	v := (uint16(vx) & 0x0F) + (uint16(vy) & 0x0F) + uint16(x)
	if v > 9 {
		v += 6
	}
	v += (uint16(vx) & 0xF0) + (uint16(vy) & 0xF0)
	if v > 0x99 {
		c.SR |= StatusCarry | StatusExtend
		v -= 0xA0
	}
	if v != 0 {
		// clear Z if result is non-zero
		c.SR &= ^StatusZero
	}

	// write result
	if mod == 0 {
		c.D[rx] = (c.D[rx] & 0xFFFFFF00) | uint32(v&0xFF)
	} else {
		_, c.err = c.M.Write(int(c.A[rx]), []byte{byte(v)})
	}

	return
}
