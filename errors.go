package m68k

import "fmt"

type opcodeError uint16

func (e opcodeError) Error() string {
	return fmt.Sprintf("unrecognized opcode: %04X", uint16(e))
}

func newOpcodeError(op uint16) error {
	return opcodeError(op)
}
