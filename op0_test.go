package m68k_test

import (
	"testing"

	"github.com/cavaliercoder/go-m68k"

	. "github.com/cavaliercoder/go-m68k/m68ktest"
)

func TestAddi(t *testing.T) {
	p := LoadFile(t, "testdata/test-op-addi.h68")
	AssertRun(t, p)
	AssertDataRegister(t, p, 0, 0xFE)
	AssertDataRegister(t, p, 1, 0xFFFE)
	AssertDataRegister(t, p, 2, 0xFFFFFFFE)
	AssertDataRegister(t, p, 3, 0xFFFFFFFF)
}

func TestMoveL(t *testing.T) {
	p := LoadFile(t, "testdata/test-op-move-l.h68")
	AssertRun(t, p)

	// source: data register
	AssertDataRegister(t, p, 1, 0x44304430)
	AssertAddressRegister(t, p, 1, 0x44304430)
	AssertLong(t, p, 0x2000, 0x44304430)
	AssertLong(t, p, 0x2004, 0x44314431)
	AssertLong(t, p, 0x2008, 0x44324432)
	// AssertLong(t, p, 0x200C, 0x44334433)
	AssertLong(t, p, 0x2010, 0x44344434)
	AssertLong(t, p, 0x12000, 0x44344434)

	// source: address register
	AssertDataRegister(t, p, 2, 0x41304130)
	AssertAddressRegister(t, p, 2, 0x41304130)
	AssertLong(t, p, 0x3000, 0x41304130)
	AssertLong(t, p, 0x3004, 0x41314131)
	AssertLong(t, p, 0x3008, 0x41324132)
	AssertLong(t, p, 0x300C, 0x41334133)
	AssertLong(t, p, 0x13000, 0x41334133)

	// source: immediate
	AssertDataRegister(t, p, 3, 0x49304930)
	AssertAddressRegister(t, p, 3, 0x49304930)
	AssertLong(t, p, 0xF000, 0x49304930)
	AssertLong(t, p, 0xF004, 0x49314931)
	AssertLong(t, p, 0xF008, 0x49324932)
	AssertLong(t, p, 0xF00C, 0x49334933)
	AssertLong(t, p, 0x1F000, 0x49334933)
}

func TestMove(t *testing.T) {
	p := LoadBytes(t, []byte{
		0x20, 0x3C, 0xFF, 0xFF, 0xFF, 0xFF, // move.l #$FFFFFFFF,D0
		0x22, 0x00, // move.l D0,D1
		0x22, 0x40, // movea.l D0,A1
		0x20, 0x7C, 0x00, 0x00, 0x20, 0x00, // movea.l #$2000,A0
		0x20, 0x80, // move.l D0,(A0)
	})
	AssertRun(t, p)
	AssertDataRegister(t, p, 0, 0xFFFFFFFF)
	AssertDataRegister(t, p, 1, 0xFFFFFFFF)
	AssertAddressRegister(t, p, 1, 0xFFFFFFFF)
	AssertAddressRegister(t, p, 0, 0x2000)
	AssertLong(t, p, 0x2000, 0xFFFFFFFF)
}

func TestOri(t *testing.T) {
	p := LoadFile(t, "testdata/test-op-ori.h68")
	p.D[2] = 0x22
	p.D[3] = 0x2222
	p.D[4] = 0x22222222
	AssertRun(t, p)
	AssertStatusRegister(t, p, 0x1F)
	AssertDataRegister(t, p, 1, 0xFFFFFFFF)
	AssertDataRegister(t, p, 2, 0x33)
	AssertDataRegister(t, p, 3, 0x3333)
	AssertDataRegister(t, p, 4, 0x33333333)
}

func TestAndi(t *testing.T) {
	b := []byte{
		0x02, 0x01, 0x00, 0x0F, // andi.b #$0F,D1
		0x02, 0x41, 0x0F, 0xFF, // andi.w #$0FFF,D1
		0x02, 0x81, 0x0F, 0x0F, 0xFF, 0xFF, // andi.l #$0F0FFFFF,D1
	}
	// p := LoadFile(t, "testdata/test-op-andi.h68")
	p := LoadBytes(t, b)
	p.D[1] = 0x55555555
	AssertRun(t, p)
	// AssertStatusRegister(t, p, 0x1F)
	AssertDataRegister(t, p, 1, 0x05050505)
}

func TestSubi(t *testing.T) {
	t.Run("FromPositive", func(t *testing.T) {
		// pos - pos = pos
		p := LoadBytes(t, []byte{
			0x04, 0x80, 0x00, 0x00, 0x00, 0x01, // subi.l #$1,D0
			0x04, 0x41, 0x00, 0x01, // subi.w #$1,D1
			0x04, 0x02, 0x00, 0x01, // subi.b #$1,D2
		})
		p.D[0], p.D[1], p.D[2] = 0x01234567, 0x01234567, 0x01234567
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0x01234566)
		AssertDataRegister(t, p, 1, 0x01234566)
		AssertDataRegister(t, p, 2, 0x01234566)
		// TODO: assert status register
	})

	t.Run("FromNegative", func(t *testing.T) {
		// neg - pos = neg
		p := LoadBytes(t, []byte{
			0x04, 0x80, 0x00, 0x00, 0x00, 0x01, // subi.l #$1,D0
			0x04, 0x41, 0x00, 0x01, // subi.w #$1,D1
			0x04, 0x02, 0x00, 0x01, // subi.b #$1,D2
		})
		p.D[0], p.D[1], p.D[2] = 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF // -1
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0xFFFFFFFE)
		AssertDataRegister(t, p, 1, 0xFFFFFFFE)
		AssertDataRegister(t, p, 2, 0xFFFFFFFE)
	})

	t.Run("SetSignBit", func(t *testing.T) {
		// pos - pos = neg
		p := LoadBytes(t, []byte{
			0x04, 0x80, 0x02, 0x02, 0x02, 0x02, // subi.l #$2020202,D0
			0x04, 0x41, 0x02, 0x02, // subi.w #$202,D1
			0x04, 0x02, 0x00, 0x02, // subi.b #$2,D2
		})
		p.D[0], p.D[1], p.D[2] = 0x01010101, 0x01010101, 0x01010101
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0xFEFEFEFF)
		AssertDataRegister(t, p, 1, 0x0101FEFF)
		AssertDataRegister(t, p, 2, 0x010101FF)
	})

	t.Run("ClearSignBit", func(t *testing.T) {
		// neg - neg = pos
		p := LoadBytes(t, []byte{
			0x04, 0x80, 0xFF, 0xFF, 0xFF, 0xFE, // subi.l #$-2,D0
			0x04, 0x41, 0xFF, 0xFE, // subi.w #$-2,D1
			0x04, 0x02, 0x00, 0xFE, // subi.b #$-2,D2
		})
		p.D[0], p.D[1], p.D[2] = 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF
		AssertRun(t, p)
		AssertDataRegister(t, p, 0, 0x00000001)
		AssertDataRegister(t, p, 1, 0xFFFF0001)
		AssertDataRegister(t, p, 2, 0xFFFFFF01)
	})
}

func TestEori(t *testing.T) {
	t.Run("ToCCR", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0x0A, 0x7C, 0xFF, 0xFF, // eori #$FFFF,SR
		})
		p.SR = m68k.StatusMask
		AssertRun(t, p)
		AssertStatusRegister(t, p, 0)
	})

	t.Run("ToSR", func(t *testing.T) {
		p := LoadBytes(t, []byte{
			0x0A, 0x3C, 0xFF, 0xFF, // eori #$FFFF,SR
		})
		p.SR = m68k.StatusMask
		AssertRun(t, p)
		AssertStatusRegister(t, p, 0xA700)
	})
}
