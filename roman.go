package romanumeral

import (
	"bytes"
	"errors"
	"fmt"
)

var romanLetterValue = map[rune]Roman{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}
var romanValues = []Roman{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanLetters = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

type Roman uint

func (r Roman) IsValid() bool {
	return r > 0 && r < 4000
}

func (r *Roman) Decode(s string) (n int, err error) {
	if r == nil {
		return 0, errors.New("numerals: Roman decode on nil pointer")
	}
	var pre Roman
	var curr Roman
	var ret Roman
	for i, ch := range s {
		curr = romanLetterValue[ch]
		if curr == 0 {
			*r = ret
			return i, nil
		}
		if pre >= curr {
			ret += pre
		} else {
			ret -= pre
		}
		pre = curr
	}
	*r = ret + curr
	return len(s), nil
}

func (r Roman) Encode() (string, error) {
	if !r.IsValid() {
		return "", errors.New("numerals: Roman access beyond the numeral")
	}
	buf := bytes.NewBuffer(nil)
	for i := 0; i != len(romanValues); {
		v := romanValues[i]
		if r >= v {
			r -= v
			buf.WriteString(romanLetters[i])
		} else {
			i++
		}
	}
	return buf.String(), nil
}

func (r Roman) String() string {
	ret, err := r.Encode()
	if err != nil {
		return fmt.Sprintf("Roman(%d): %s", uint(r), err.Error())
	}
	return ret
}