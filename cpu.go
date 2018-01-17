/*
Manual: http://www.easy68k.com/paulrsm/doc/68kprm.pdf

Glossary:

EA: Effective Address - Addressing mode and location
Dn: Data Register Direct Mode
An: Address Register Direct Mode
*/
package m68k

import "fmt"

type Register uint32

// See 1.1.4
type ConditionCodeRegister Register

type RegisterBank struct {
	D0, D1, D2, D3, D4, D5, D6, D7 Register // Data registers
	A0, A1, A2, A3, A4, A5, A6     Register // Address registers
	A7                             Register // Program Counter
	CCR                            ConditionCodeRegister
}

type Microprocessor struct {
	RegisterBank
}

func (c *Microprocessor) Run(b []byte) {
	for i := 0; i < len(b); i++ {
		if b[i]&0xD0 == 0xD0 {
			fmt.Printf("%04d ADD %08b\n", i, b[i])
		} else {
			// fmt.Printf("%04d Unknown: %08b\n", i, b[i])
		}
	}
}
