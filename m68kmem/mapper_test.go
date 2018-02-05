package m68kmem

import (
	"fmt"
)

func ExampleMapper() {
	// create a memory mapper
	m := NewMapper()

	// map a ROM device to 0x00
	dev1 := NewROM([]byte("This is device one.............\n"))
	m.Map(dev1, 0, 0x1F)

	// map a ROM device to 0x20
	dev2 := NewROM([]byte("This is device two.............\n"))
	m.Map(dev2, 0x20, 0x3F)

	// map a ROM device to 0x40
	dev3 := NewROM([]byte("This is device three...........\n"))
	m.Map(dev3, 0x40, 0x5F)

	// read 96 bytes from the memory mapper
	b := make([]byte, 0x60)
	m.Read(0, b)
	fmt.Print(string(b))

	// Output:
	// This is device one.............
	// This is device two.............
	// This is device three...........
}
