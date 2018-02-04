package m68kmem

type nop struct{}

// NewNop returns a No-Operation memory device that neither stores nor
// retrieves data. It serves only for debug or placeholder purposes.
func NewNop() Memory {
	return &nop{}
}

func (c *nop) Read(addr int, p []byte) (n int, err error) {
	n = len(p)
	return
}

func (c *nop) Write(addr int, p []byte) (n int, err error) {
	n = len(p)
	return
}
