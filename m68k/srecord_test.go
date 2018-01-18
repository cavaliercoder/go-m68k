package m68k

import (
	"strings"
	"testing"
)

func TestSRecord(t *testing.T) {
	tests := []struct {
		Input   string
		Expect  string
		Address int
	}{
		{
			Input: `S004000020DB
S125100041FA001212186708103C00064E4F60F44E7227005468697320697320746865206D65D3
S10B102273736167650D0A0098
S9030000FC
`,
			Expect:  "This is the message",
			Address: 0,
		},
	}

	for _, test := range tests {
		srecords, err := ReadSRecords(strings.NewReader(test.Input))
		if err != nil {
			t.Fatal(err)
		}
		m := NewMemory(0x2000)
		for _, s := range srecords {
			if s.IsData() {
				if err := s.Load(m); err != nil {
					t.Fatal(err)
				}
			}
		}
		b := make([]byte, 19)
		m.Read(0x1014, b)
		actual := string(b)
		if test.Expect != actual {
			t.Fatalf("expected '%s' @ 0x1014, got '%s'", test.Expect, actual)
		}
	}
}
