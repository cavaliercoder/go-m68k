package m68k_test

import (
	"testing"

	"github.com/cavaliercoder/go-m68k"

	. "github.com/cavaliercoder/go-m68k/m68ktest"
)

func TestAbcd(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0xC1, 0x01, // abcd D1,D0
		})
		p.D[0], p.D[1] = 0x09, 0x09
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0x18)
		AssertStatusRegister(t, p, 0)
	})

	t.Run("WithExtend", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0xC1, 0x01, // abcd D1,D0
		})
		p.SR |= m68k.StatusExtend
		p.D[0], p.D[1] = 0x09, 0x09
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0x19)
		AssertStatusRegister(t, p, 0)
	})

	t.Run("WithInvalidInput", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0xC1, 0x01, // abcd D1,D0
		})
		p.D[0], p.D[1] = 0xFF, 0xFF
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0x64)
		AssertStatusRegister(t, p, 0x11)
	})

	t.Run("WithInvalidInputAndExtend", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0xC1, 0x01, // abcd D1,D0
		})
		p.SR |= m68k.StatusExtend
		p.D[0], p.D[1] = 0xFF, 0xFF
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0x65)
		AssertStatusRegister(t, p, 0x11)
	})
}
