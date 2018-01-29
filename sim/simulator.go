/*
Package sim provides a simple simulator for interacting with the Motorola
68000 emulator package.

The simulator provides standard I/O using a similar protocol to the Teesside
simulator.

See: http://www.easy68k.com/QuickStart/TrapTasks.htm
*/
package sim

import (
	"fmt"
	"io"
	"os"

	"github.com/cavaliercoder/go-m68k"
)

type Simulator struct {
	p *m68k.Processor
	w io.Writer
	r io.Reader

	b [255]byte
}

func NewClassic() (s *Simulator, err error) {
	p := &m68k.Processor{}
	p.Reset()
	s = &Simulator{
		p: p,
		w: os.Stdout,
		r: os.Stdin,
	}
	err = p.RegisterTrapHandler(47, s)
	return
}

// New returns a new Simulator and registers each of its trap handlers with
// the given processor.
func New(p *m68k.Processor, w io.Writer, r io.Reader) (s *Simulator, err error) {
	s = &Simulator{p: p, w: w, r: r}
	err = p.RegisterTrapHandler(47, s)
	return
}

func (c *Simulator) Run() error {
	return c.p.Run()
}

func (c *Simulator) Processor() *m68k.Processor {
	return c.p
}

// Exception handler TRAP 15 instructions from the underlying processor.
func (c *Simulator) Exception(p *m68k.Processor, v int) (err error) {
	switch p.D[0] {
	default:
		return fmt.Errorf("unrecognized trap task identifier: %d", p.D[0])

	case 0:
		// Display string at (A1), D1.W bytes long (max 255) with carriage return
		// and line feed (CR, LF). (see task 13)
		n := uint16(c.p.D[1])
		_, err = c.p.M.Read(int(c.p.A[1]), c.b[:n])
		if err != nil {
			return
		}
		c.b[n], c.b[n+1] = '\r', '\n'
		_, err = c.w.Write(c.b[:n+2])

	case 1:
		// Display string at (A1), D1.W bytes long (max 255) without CR, LF. (see
		// task 14)
		n := uint16(c.p.D[1])
		_, err = c.p.M.Read(int(c.p.A[1]), c.b[:n])
		if err != nil {
			return
		}
		_, err = c.w.Write(c.b[:n])

	case 6:
		// Display single character in D1.B.
		_, err = c.w.Write([]byte{byte(c.p.D[1])})

	case 9: // halt simulator
		err = io.EOF
	}

	return
}
