package romanumeral

import (
	"testing"
)

func TestRoman(t *testing.T) {
	for i := Roman(1); i != 4000; i++ {
		tmp, err := i.Encode()
		if err != nil {
			t.Error(err)
		}
		var d Roman
		_, err = d.Decode(tmp)
		if err != nil {
			t.Error(err)
		}
		if d != i {
			t.Fail()
		}
	}
}

func TestRomanIsValid(t *testing.T) {
	var r Roman
	if r.IsValid() {
		t.Fail()
	}
}
