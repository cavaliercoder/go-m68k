/*
Package sim provides a simple simulator for interacting with the Motorola
68000 emulator package.

The simulator provides standard I/O using a similar protocol to the Easy68K
simulator.

See: http://www.easy68k.com/QuickStart/TrapTasks.htm
*/
package sim

import (
	"fmt"

	"github.com/cavaliercoder/go-m68k"
)

type Simulator struct {
	p *m68k.Processor
}

// New returns a new Simulator and registers each of its trap handlers with
// the given processor.
func New(p *m68k.Processor) (s *Simulator, err error) {
	s = &Simulator{p: p}
	err = p.RegisterTrapHandler(47, s)
	return
}

func (c *Simulator) Exception(p *m68k.Processor, v int) error {
	fmt.Printf("===> %d <===\n", v)
	return nil
}
