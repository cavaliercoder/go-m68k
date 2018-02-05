package m68kmem

import "fmt"

type accessViolationError uint32

func (e accessViolationError) Error() string {
	return fmt.Sprintf("access violation: 0x%X", uint32(e))
}

// AccessViolationError returns an error indicating a memory address access
// violation for the given address.
func AccessViolationError(addr uint32) error {
	return accessViolationError(addr)
}

func isAccessViolationError(err error) bool {
	_, ok := err.(accessViolationError)
	return ok
}
