package m68k

import (
	"errors"
	"fmt"
)

var (
	errNoProgram      = errors.New("no program loaded or memory device attached")
	errBadAddress     = errors.New("illegal effective address")
	errBadOpcode      = errors.New("illegal operation code")
	errBadOpSize      = errors.New("illegal operand size")
	errNotImplemented = errors.New("operation code not implemented")
)

type traceableError struct {
	Addr   uint32
	Opcode uint16
	Err    error
}

func newTraceableError(addr uint32, op uint16, err error) error {
	return &traceableError{addr, op, err}
}

func (e *traceableError) Error() string {
	return fmt.Sprintf("%s (Op: 0x%04X, PC: 0x%X)", e.Err.Error(), e.Opcode, e.Addr)
}
