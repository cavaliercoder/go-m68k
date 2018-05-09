package m68kmem

import (
	"encoding/hex"
	"fmt"
	"io"
)

type tracer struct {
	m Memory
	w io.Writer
	s string
}

// NewTracer wraps a memory device to log all reads and write to the given
// io.Writer.
func NewTracer(name string, w io.Writer, m Memory) Memory {
	return &tracer{m: m, w: w, s: name}
}

func (c *tracer) Read(addr int, p []byte) (n int, err error) {
	n, err = c.m.Read(addr, p)
	if err != nil {
		fmt.Fprintf(c.w, "%08X [%s] error reading %d bytes: %v\n", addr, c.s, len(p), err)
		return
	}
	fmt.Fprintf(c.w, "%08X [%s] read %d bytes: %s\n", addr, c.s, n, hex.EncodeToString(p))
	return
}

func (c *tracer) Write(addr int, p []byte) (n int, err error) {
	n, err = c.m.Write(addr, p)
	if err != nil {
		fmt.Fprintf(c.w, "%08X [%s] error writing %d bytes: %v\n", addr, c.s, len(p), err)
		return
	}
	fmt.Fprintf(c.w, "%08X [%s] wrote %d bytes: %s\n", addr, c.s, n, hex.EncodeToString(p))
	return
}

func (c *tracer) Reset() (err error) {
	return c.m.Reset()
}
