package dump_test

import (
	"os"

	"github.com/cavaliercoder/go-m68k/dump"
	. "github.com/cavaliercoder/go-m68k/m68ktest"
)

func ExampleProcessor() {
	p := LoadFile(nil, "../testdata/test-sim.h68")
	p.Run()

	dump.Processor(os.Stdout, p)

	// output:
	// PC: 00001014 SR: 00002700
	// D0: 00000006 A0: 00001023
	// D1: 00000000 A1: 00000000
	// D2: 00000000 A2: 00000000
	// D3: 00000000 A3: 00000000
	// D4: 00000000 A4: 00000000
	// D5: 00000000 A5: 00000000
	// D6: 00000000 A6: 00000000
	// D7: 00000000 A7: 00000000
}

func ExampleMemory() {
	p := LoadFile(nil, "../testdata/test-sim.h68")
	p.Run()

	dump.Memory(os.Stdout, p.M)

	// output:
	// *
	// 00001000  41 FA 00 12 12 18 67 08  10 3C 00 06 4E 4F 60 F4  |A.....g..<..NO`.|
	// 00001010  4E 72 27 00 48 65 6C 6C  6F 20 77 6F 72 6C 64 21  |Nr'.Hello world!|
	// 00001020  0D 0A 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
	// *
}
