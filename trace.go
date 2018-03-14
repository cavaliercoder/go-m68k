package m68k

import (
	"fmt"
)

const noSize = uint16(0xFFFF)

type stepTrace struct {
	addr         uint32
	op, src, dst string
	sz           uint16
	n            int // cycle count
}

func (c *stepTrace) String() string {
	operands := c.dst
	if c.dst == "" {
		operands = c.src
	} else if c.src != "" {
		operands = fmt.Sprintf("%s,%s", c.src, c.dst)
	}
	if c.sz == noSize {
		return fmt.Sprintf("%08X %s %s", c.addr, c.op, operands)
	}
	sz := []string{"b", "w", "l"}[c.sz]
	return fmt.Sprintf("%08X %s.%s %s", c.addr, c.op, sz, operands)
}
