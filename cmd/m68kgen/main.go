package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cavaliercoder/go-m68k/generator"
)

var (
	outFile = flag.String("out", "", "output file path")
)

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func main() {
	flag.Parse()
	w := os.Stdout
	if outFile != nil && *outFile != "" {
		f, err := os.OpenFile(*outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			die("error opening output file: %v", err)
		}
		defer f.Close()
		w = f
	}
	if err := generator.Generate(w); err != nil {
		die("error generating code: %v\n", err)
	}
}
