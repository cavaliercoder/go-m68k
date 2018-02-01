package main

import (
	"fmt"
	"os"

	"github.com/cavaliercoder/go-m68k"
	"github.com/cavaliercoder/go-m68k/sim"
	"github.com/cavaliercoder/go-m68k/srec"
)

func main() {
	if len(os.Args) < 2 {
		usage(1)
	}

	// read program file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	// create simulator
	s, err := sim.NewClassic()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	s.Processor().TraceWriter = os.Stderr

	// load program into memory
	records, err := srec.Read(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := srec.Load(s.Processor().M, records); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// run program
	if err := s.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// dump state
	m68k.DumpState(os.Stderr, s.Processor())
	m68k.Dump(os.Stderr, s.Processor().M)
}

func usage(n int) {
	w := os.Stdout
	if n > 0 {
		w = os.Stderr
	}
	fmt.Fprintf(w, "usage: %s [program]\n", os.Args[0])
	os.Exit(n)
}
