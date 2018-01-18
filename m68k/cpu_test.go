package m68k

import (
	"strings"
	"testing"
)

func TestCPU(t *testing.T) {
	prg := `S004000020DB
S11B100030382000D0782002D0782004044001F031C020064E72270043
S9030000FC
`
	cpu := &Processor{}
	cpu.Reset()
	srecords, err := ReadSRecords(strings.NewReader(prg))
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range srecords {
		if s.IsData() {
			if err := s.Load(cpu.Memory); err != nil {
				t.Fatal(err)
			}
		}
	}
	cpu.Run()
}
