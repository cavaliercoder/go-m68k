package m68k_test

import (
	"testing"

	. "github.com/cavaliercoder/go-m68k/m68ktest"
)

func TestBra(t *testing.T) {
	p := LoadBytes(t, []byte{
		0x66, 0x06, //
	})
	AssertRun(t, p)
}
