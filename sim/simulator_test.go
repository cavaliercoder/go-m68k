package sim

import (
	"testing"

	"github.com/cavaliercoder/go-m68k/m68ktest"
)

func TestSim(t *testing.T) {
	p := m68ktest.LoadBytes(t, []byte{0x4E, 0x4F})
	_, err := New(p)
	if err != nil {
		t.Fatal(err)
	}
	m68ktest.AssertRun(t, p)
}
