package m68k

import (
	"errors"
	"fmt"
)

var (
	ErrNoProgram          = errors.New("no program loaded or memory device attached")
	ErrAddressOutOfBounds = errors.New("memory address out of bounds")
)

type opcodeError uint16

func (e opcodeError) Error() string {
	return fmt.Sprintf("unrecognized opcode: %04X", uint16(e))
}

func newOpcodeError(op uint16) error {
	return opcodeError(op)
}
