package m68kmem

import (
	"io"
)

// Memory is an interface for any IO device that is addressable via memory
// mapping. This interface will be satisfied by random access memory, a virtual
// memory manager, peripheral devices, etc.
type Memory interface {
	Read(addr int, p []byte) (n int, err error)
	Write(addr int, p []byte) (n int, err error)
}

// Clear wipes the given Memory device, resetting all reachable addresses to
// zero.
func Clear(m Memory) {
	if m == nil {
		return
	}
	addr := 0
	zero := [32]byte{}
	for {
		n, err := m.Write(addr, zero[:])
		// TODO: clear breaks on access violation
		if err == io.ErrShortWrite || err == io.EOF || isAccessViolationError(err) {
			break
		}
		if err != nil {
			panic(err)
		}
		addr += n
	}
}
