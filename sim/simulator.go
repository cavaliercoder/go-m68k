/*
Package sim provides a simple simulator for interacting with the Motorola
68000 emulator package.

The simulator provides standard I/O using a similar protocol to the Teesside
simulator.

See: http://www.easy68k.com/QuickStart/TrapTasks.htm
*/
package sim

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cavaliercoder/go-m68k"
	"github.com/cavaliercoder/go-m68k/m68kmem"
)

var (
	ErrNoProcessor = errors.New("no processor configured")
	ErrNoWriter    = errors.New("write operation requested but no writer is attached")
	ErrNoReader    = errors.New("read operation requested but no reader is attached")
)

// A Simulator provides standard I/O for a Motorola 68000 emulator using a
// similar protocol to the Teesside simulator.
type Simulator struct {
	Processor *m68k.Processor
	Writer    io.Writer
	Reader    io.Reader

	b [255]byte
}

// New returns a new simulator with default configuration.
func New() (s *Simulator, err error) {
	p := &m68k.Processor{
		PC: 0x1000,
		M:  &m68kmem.Decoder{M: m68kmem.NewRAM(0x10000)}, // 64KB
	}
	s = &Simulator{
		Processor: p,
		Writer:    os.Stdout,
		Reader:    os.Stdin,
	}
	err = s.Register()
	return
}

// Register configures the given Processor for use with the Simulator.
func (c *Simulator) Register() (err error) {
	if c.Processor == nil {
		return ErrNoProcessor
	}
	err = c.Processor.RegisterTrapHandler(47, c)
	return
}

// Trace configures the simulator to write trace messages to the given Writer.
func (c *Simulator) Trace(w io.Writer) {
	c.Processor.TraceWriter = w
}

// Run configures the underlying Processor and executes any program loaded
// into memory, starting from the location of the Program Counter.
func (c *Simulator) Run() (err error) {
	if err = c.Register(); err != nil {
		return
	}
	return c.Processor.Run()
}

func (c *Simulator) write(p []byte) (n int, err error) {
	if c.Writer == nil {
		return 0, ErrNoWriter
	}
	return c.Writer.Write(p)
}

func (c *Simulator) read(p []byte) (n int, err error) {
	if c.Reader == nil {
		return 0, ErrNoReader
	}
	return c.Reader.Read(p)
}

// Trap handles TRAP 15 instructions from the underlying Processor.
func (c *Simulator) Trap(p *m68k.Processor, v int) (err error) {
	switch p.D[0] {
	default:
		return fmt.Errorf("unrecognized trap task identifier: %d", p.D[0])

	case 0:
		// Display string at (A1), D1.W bytes long (max 255) with carriage return
		// and line feed (CR, LF). (see task 13)
		n := uint16(c.Processor.D[1])
		_, err = c.Processor.M.Read(int(c.Processor.A[1]), c.b[:n])
		if err != nil {
			return
		}
		c.b[n], c.b[n+1] = '\r', '\n'
		_, err = c.write(c.b[:n+2])

	case 1:
		// Display string at (A1), D1.W bytes long (max 255) without CR, LF. (see
		// task 14)
		n := uint16(c.Processor.D[1])
		_, err = c.Processor.M.Read(int(c.Processor.A[1]), c.b[:n])
		if err != nil {
			return
		}
		_, err = c.write(c.b[:n])

	case 6:
		// Display single character in D1.B.
		_, err = c.write([]byte{byte(c.Processor.D[1])})

	case 9:
		// halt simulator
		c.Processor.Stop()

	case 13:
		// Display the NULL terminated string at (A1) with CR, LF.
		addr := int(p.A[1])
		b := make([]byte, 1)
		for {
			_, err = p.M.Read(addr, b)
			if err != nil {
				return
			}
			if b[0] == 0 {
				break
			}
			c.write(b)
			addr++
		}
		c.write([]byte{'\r', '\n'})

	case 14:
		// Display the NULL terminated string at (A1) without CR, LF.
		addr := int(p.A[1])
		b := make([]byte, 1)
		for {
			_, err = p.M.Read(addr, b)
			if err != nil {
				return
			}
			if b[0] == 0 {
				break
			}
			c.write(b)
			addr++
		}
	}
	return
}
