package m68k

import (
	"errors"
	"fmt"
)

var (
	ErrNoProgram      = errors.New("no program loaded or memory device attached")
	ErrBadAddress     = errors.New("illegal effective address")
	ErrBadOpSize      = errors.New("illegal operand size")
	ErrNotImplemented = errors.New("opcode not implemented")
)

type opcodeError struct {
	op uint16
}

func (e *opcodeError) Error() string {
	return fmt.Sprintf("unrecognized opcode: %04X", e.op)
}

func newOpcodeError(op uint16) error {
	return &opcodeError{op}
}
