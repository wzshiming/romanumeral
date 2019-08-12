package romanumeral

import (
	"testing"
)

func TestRoman(t *testing.T) {
	for i := Roman(1); i != 4000; i++ {
		tmp, err := i.EncodeToString()
		if err != nil {
			t.Error(err)
		}
		var d Roman
		off, err := d.DecodeString(tmp)
		if err != nil {
			t.Error(err)
		}
		if d != i {
			t.Fatal(d, i)
		}
		if len(tmp) != off {
			t.Fatal(len(tmp), off)
		}
	}
}

func TestRomanIsValid(t *testing.T) {
	var r Roman
	if r.IsValid() {
		t.Fail()
	}
}
