package main

import "io"

type IndentWriter struct {
	w      io.Writer
	n      int
	prefix []byte
}

func NewIndentWriter(w io.Writer, prefix string, n int) *IndentWriter {
	return &IndentWriter{
		w:      w,
		n:      n,
		prefix: []byte(prefix),
	}
}

func (w *IndentWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}
	var nn int
	indent := true
	for i := 0; i < len(p); i++ {
		if p[i] == '\n' {
			indent = true
		}
		if indent && p[i] != '\n' {
			indent = false
			nn, err = w.indent()
			n += nn
			if err != nil {
				break
			}
		}
		nn, err = w.w.Write(p[i : i+1])
		n += nn
		if err != nil {
			break
		}
	}
	return
}

func (w *IndentWriter) indent() (n int, err error) {
	var nn int
	for i := 0; i < w.n; i++ {
		nn, err = w.w.Write(w.prefix)
		n += nn
		if err != nil {
			return
		}
	}
	return
}

func (w *IndentWriter) Increment(n int) int {
	w.n += n
	return w.n
}

func (w *IndentWriter) Decrement(n int) int {
	w.n -= n
	return w.n
}

func (w *IndentWriter) Reset() int {
	return w.Set(0)
}

func (w *IndentWriter) Set(n int) int {
	w.n = n
	return w.n
}
