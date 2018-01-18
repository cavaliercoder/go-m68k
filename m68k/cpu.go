package m68k

/*
Manual: http://www.easy68k.com/paulrsm/doc/68kprm.pdf

Glossary:

EA: Effective Address - Addressing mode and location
Dn: Data Register Direct Mode
An: Address Register Direct Mode
*/

type Processor struct {
	D      [8]uint32 // Data registers
	A      [8]uint32 // Address registers
	CCR    uint32    // Condition Code Register
	Memory Memory
}

func (c *Processor) Reset() {
	for i := 0; i < len(c.D); i++ {
		c.D[i] = 0
		c.A[i] = 0
	}
	c.CCR = 0
	c.Jump(0x1000)
	if c.Memory == nil {
		c.Memory = NewMemory(0x2000)
	} else {
		// c.Memory.Reset()
	}
}

func (c *Processor) Jump(addr int) error {
	c.A[7] = uint32(addr)
	return nil
}

func (c *Processor) Run() error {
	ins := make([]byte, 16)
	for {
		if _, err := c.Memory.Read(int(c.A[7]), ins[0:2]); err != nil {
			return err
		}
		c.A[7] += 2
	}
}
