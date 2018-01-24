package m68k

//go:generate ./cmd/m68kgen/m68kgen -out opcodes.go

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrNoProgram = errors.New("No program loaded or memory device attached")
)

// A Processor emulates the Motorola 68000 microprocessor.
type Processor struct {
	D   [8]uint32 // Data registers
	A   [8]uint32 // Address registers
	CCR uint32    // Condition Code Register
	PC  uint32    // Program Counter
	M   Memory    // System memory controller

	// TraceWriter specifies where trace log output should be written to. If
	// TraceWriter is nil, no logging is performed.
	TraceWriter io.Writer

	buf [64]byte // general purpose buffer
	op  uint16
	err error
}

// Reset resets all processor and memory state.
func (c *Processor) Reset() {
	*c = Processor{
		TraceWriter: c.TraceWriter,
	}
	c.Jump(0x1000)
	if c.M == nil {
		c.M = NewRAM(0x40000) // 256KB
	}
	clearMemory(c.M)
}

func clearMemory(m Memory) {
	if m == nil {
		return
	}
	addr := 0
	zero := [32]byte{}
	for {
		n, err := m.Write(addr, zero[:])
		if err == io.ErrShortWrite {
			break
		}
		if err != nil {
			panic(err)
		}
		addr += n
	}
}

// Jump sets the value of the program counter to the given memory address.
func (c *Processor) Jump(addr uint32) error {
	c.PC = addr
	return nil
}

// Run executes any program loaded into memory, starting from the program
// counter value, running until completion.
func (c *Processor) Run() error {
	if c.M == nil {
		return ErrNoProgram
	}
	for c.err == nil {
		c.Step()
	}
	return c.err
}

// Step executes the single instruction located at the address specified by the
// program counter register.
func (c *Processor) Step() error {
	// read from program pointer into op
	if _, c.err = c.M.Read(int(c.PC), c.buf[0:2]); c.err != nil {
		return c.err
	}
	c.op = uint16(c.buf[0])<<8 + uint16(c.buf[1])
	if c.op == 0 {
		// TODO: this is a valid instruction. Find a way to terminate programs
		// without buffer overrun.
		c.err = io.EOF
		return c.err
	}
	f := c.mapFn(c.op)
	if f == nil {
		c.err = fmt.Errorf("unrecognised opcode: %04X @ %04X", c.op, c.PC)
	} else {
		f()
	}
	return c.err
}

func (c *Processor) tracef(format string, a ...interface{}) {
	if c.TraceWriter == nil {
		return
	}
	fmt.Fprintf(c.TraceWriter, format, a...)
}

func (c *Processor) read(addr uint32, b []byte) {
	_, c.err = c.M.Read(int(addr), b)
	// TODO: create processor exception for IO errors
}

func (c *Processor) readLong(addr uint32) uint32 {
	_, c.err = c.M.Read(int(addr), c.buf[:4])
	return uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
}

func (c *Processor) writeLong(addr uint32, n uint32) {
	c.buf[0] = byte(n >> 24)
	c.buf[1] = byte(n >> 16)
	c.buf[2] = byte(n >> 8)
	c.buf[3] = byte(n)
	_, c.err = c.M.Write(int(addr), c.buf[:4])
}

// readImm reads an immediate value of size n where n is one of:
//     0 - byte
//     1 - word
//     2 - long
func (c *Processor) readImm(n uint16) uint32 {
	sz := []uint32{1, 2, 4}[n]
	if _, c.err = c.M.Read(int(c.PC), c.buf[:sz]); c.err != nil {
		return 0 // TODO: handle error
	}
	c.PC += sz
	v := uint32(c.buf[0])
	for i := uint32(1); i < sz; i++ {
		v <<= 8
		v |= uint32(c.buf[i])
	}
	return v
}

func (c *Processor) readImmByte() byte {
	if _, c.err = c.M.Read(int(c.PC), c.buf[:1]); c.err != nil {
		return 0 // TODO: handle error
	}
	c.PC++
	return c.buf[0]
}

func (c *Processor) readImmWord() uint16 {
	if _, c.err = c.M.Read(int(c.PC), c.buf[:2]); c.err != nil {
		return 0 // TODO: handle error
	}
	c.PC += 2
	return uint16(c.buf[0])<<8 | uint16(c.buf[1])
}

func (c *Processor) readImmLong() uint32 {
	if _, c.err = c.M.Read(int(c.PC), c.buf[:4]); c.err != nil {
		return 0 // TODO: handle error
	}
	c.PC += 4
	return uint32(c.buf[0])<<24 | uint32(c.buf[1])<<16 | uint32(c.buf[2])<<8 | uint32(c.buf[3])
}

// DumpState prints the state of of a 68000 processor to the given writer in
// hexidecimal format.
func DumpState(w io.Writer, p *Processor) {
	for i := 0; i < len(p.D); i++ {
		fmt.Fprintf(w, "D%d: %08X A%d: %08X\n", i, p.D[i], i, p.A[i])
	}
}
