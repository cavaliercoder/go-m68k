package m68k

import (
	"fmt"
)

type stepTrace struct {
	addr         uint32
	op, src, dst string
	sz           uint16
	// err          error
	n int // cycle count
}

func (c *stepTrace) String() string {
	operands := c.dst
	if c.dst == "" {
		operands = c.src
	} else if c.src != "" {
		operands = fmt.Sprintf("%s,%s", c.src, c.dst)
	}
	if c.sz < 0 {
		return fmt.Sprintf("%04X %s %s", c.addr, c.op, operands)
	}
	sz := []string{"b", "w", "l"}[c.sz]
	return fmt.Sprintf("%04X %s.%s %s", c.addr, c.op, sz, operands)
}
