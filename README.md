# go-m68k

[![GoDoc](https://godoc.org/github.com/cavaliercoder/go-m68k?status.svg)](https://godoc.org/github.com/cavaliercoder/go-m68k) [![Build Status](https://travis-ci.org/cavaliercoder/go-m68k.svg)](https://travis-ci.org/cavaliercoder/go-m68k) [![Go Report Card](https://goreportcard.com/badge/github.com/cavaliercoder/go-m68k)](https://goreportcard.com/report/github.com/cavaliercoder/go-m68k)

This project contains a number of tools and libraries for emulating the Motorola
68000 chipset; used in devices like the Apple Macintosh, Sega Megadrive and
Commodore Amiga.

__WARNING__ This project is an incomplete work in progress.

The goal of this project is less about performance or competing with existing
emulators. It's more about enabling a better development and learning experience
for novices like me.

## Installation

```bash
$ go get github.com/cavaliercoder/go-m68k/...
```


## Example

Compile the following Motorola 68000 assembly using a supported compiler:

```plain
* Test 68000 simulator program

WRCHAR  EQU     6                       Trap # for write D1.B char to screen

        ORG $1000

        LEA      MESSAGE(PC),A0         Point A0 to start of message
NEXT    MOVE.B  (A0)+,D1                Get character, increment pointer
        BEQ.S   FINISH                  Exit if end
        MOVE.B  #WRCHAR,D0              Set up trap to write to screen
        TRAP    #15                     Write char. to screen
        BRA.S   NEXT                    ..and loop back
FINISH  STOP    #$2700                  Halt
MESSAGE DC.B    'Hello world!',$0D,$0A,0

        END     $1000
```

Run the program:

```plain
$ m68k hello-world.h68
Hello world!
```
