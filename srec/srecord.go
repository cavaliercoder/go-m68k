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

	"github.com/cavaliercoder/go-m68k/m68kmem"
)

var (
	ErrInvalidChecksum = errors.New("checksum validation failed")
)

// A Record is a single record from a Motorola 68000 S-Record encoded file. This
// format is used to share 68000 programs in plain text and is detailed in
// Appendix C of the M68000 Family Programmer's Reference Manual.
type Record struct {
	b []byte
}

// Read reads and parses S-Records from the given Reader. Records are delimited
// by line breaks.
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

// Parse decodes the given byte array into a Record struct. The input is
// expected to be a heximdecimal encoded record with the 'S' prefix.
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

// Load copies all data from the given slice of records into a 68000 memory
// bank.
func Load(m m68kmem.Memory, records []*Record) error {
	for _, s := range records {
		if s.IsData() {
			_, err := m.Write(s.Address(), s.Data())
			if err != nil {
				return err
			}
		}
	}
	return nil
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

// Type returns the type of the Record. E.g. "S1".
func (s *Record) Type() string {
	return string(s.b[:2])
}

// IsData returns true if the Record contains code or data.
func (s *Record) IsData() bool {
	return s.b[1] > 0 && s.b[1] < 4
}

// Address specifies the destination memory address where data will be loaded
// for a data record.
func (s *Record) Address() int {
	addr := 0
	w := s.addressWidth()
	for i := 0; i < w; i++ {
		addr |= int(s.b[3+i]) << (8 * uint(w-1-i))
	}
	return addr
}

// Data return the raw data section of the Record.
func (s *Record) Data() []byte {
	return s.b[3+s.addressWidth() : len(s.b)-1]
}

// Checksum returns the checksum byte for the Record.
func (s *Record) Checksum() byte {
	return s.b[len(s.b)-1]
}

// String returns a string represenation of the Record.
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
