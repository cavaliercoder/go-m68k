package m68k_test

import (
	"testing"

	. "github.com/cavaliercoder/go-m68k/m68ktest"
)

func TestLea(t *testing.T) {
	t.Run("WithAbsoluteShort", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0x41, 0xF8, 0x7F, 0xFF, // lea $7FFF.w,A0
		})
		AssertRun(t, p)
		AssertAddressRegister(t, p, 0, 0x7FFF)
	})

	t.Run("WithAbsoluteLong", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0x41, 0xF9, 0xFF, 0xFF, 0xFF, 0xFF, // lea $FFFFFFFF.l,A0
		})
		AssertRun(t, p)
		AssertAddressRegister(t, p, 0, 0xFFFFFFFF)
	})
}