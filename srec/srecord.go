/*
Package srec provides decoding of the Motorola S-Record file format.
*/
package srec

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/cavaliercoder/go-m68k/m68k"
)

var (
	ErrInvalidChecksum = errors.New("checksum validation failed")
)

type Record struct {
	b []byte
}

func Read(r io.Reader) ([]*Record, error) {
	v := make([]*Record, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		sr, err := Parse(s.Bytes())
		if err != nil {
			return nil, err
		}
		v = append(v, sr)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return v, nil
}

func Parse(b []byte) (*Record, error) {
	s := &Record{}
	s.b = make([]byte, len(b)/2+1)
	s.b[0] = 'S'
	s.b[1] = b[1] - '0'
	if _, err := hex.Decode(s.b[2:], b[2:]); err != nil {
		return nil, err
	}

	// validate checksum
	var sum uint32
	for i := 2; i < len(s.b)-1; i++ {
		sum += uint32(s.b[i])
	}
	sum = ^(sum | 0xFFFFFF00)
	if byte(sum) != s.Checksum() {
		return nil, ErrInvalidChecksum
	}
	return s, nil
}

// Load copies the data from a data record into the given Memory.
func (s *Record) Load(m m68k.Memory) (err error) {
	if !s.IsData() {
		return errors.New("s-record is not a data record")
	}
	_, err = m.Write(s.Address(), s.Data())
	return
}

func (s *Record) addressWidth() int {
	if s.b[1] == 4 || s.b[1] > 9 {
		panic("unrecognized s-record type")
	}
	return []int{
		2,  // S0
		2,  // S1
		3,  // S2
		4,  // S3
		-1, // S4 - Reserved
		2,  // S5
		3,  // S6
		4,  // S7
		3,  // S8
		2,  // S9
	}[s.b[1]]
}

func (s *Record) Type() string {
	return string(s.b[:2])
}

func (s *Record) IsData() bool {
	return s.b[1] > 0 && s.b[1] < 4
}

func (s *Record) Address() int {
	addr := 0
	w := s.addressWidth()
	for i := 0; i < w; i++ {
		addr |= int(s.b[3+i]) << (8 * uint(w-1-i))
	}
	return addr
}

func (s *Record) Data() []byte {
	return s.b[3+s.addressWidth() : len(s.b)-1]
}

func (s *Record) Checksum() byte {
	return s.b[len(s.b)-1]
}

func (s *Record) String() string {
	if s.b[1] > 9 {
		return "unknown"
	}
	return []func() string{
		func() string { return "header" },                                           // S0
		func() string { return fmt.Sprintf("data @ 0x%04X", s.Address()) },          // S1
		func() string { return fmt.Sprintf("data @ 0x%06X", s.Address()) },          // S2
		func() string { return fmt.Sprintf("data @ 0x%08X", s.Address()) },          // S3
		func() string { return "reserved" },                                         // S4
		func() string { return fmt.Sprintf("record count (%d)", s.Address()) },      // S5
		func() string { return fmt.Sprintf("record count (%d)", s.Address()) },      // S6
		func() string { return fmt.Sprintf("start address @ 0x%08X", s.Address()) }, // S7
		func() string { return fmt.Sprintf("start address @ 0x%06X", s.Address()) }, // S8
		func() string { return fmt.Sprintf("start address @ 0x%04X", s.Address()) }, // S9
	}[s.b[1]]()
}
