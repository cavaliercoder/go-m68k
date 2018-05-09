package m68kmem

const (
	MinAddress uint32 = 0
	MaxAddress uint32 = 0xFFFFFF
)

// Memory is an interface for any IO device that is addressable via memory
// mapping. This interface will be satisfied by random access memory, a virtual
// memory manager, peripheral devices, etc.
type Memory interface {
	Read(addr int, p []byte) (n int, err error)
	Write(addr int, p []byte) (n int, err error)
	Reset() (err error)
}
