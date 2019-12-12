# romanumeral
Roman numerals in Go


[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/romanumeral)](https://goreportcard.com/report/github.com/wzshiming/romanumeral)
[![GoDoc](https://godoc.org/github.com/wzshiming/romanumeral?status.svg)](https://godoc.org/github.com/wzshiming/romanumeral)
[![GitHub license](https://img.shields.io/github/license/wzshiming/romanumeral.svg)](https://github.com/wzshiming/romanumeral/blob/master/LICENSE)
[![gocover.io](https://gocover.io/_badge/github.com/wzshiming/romanumeral)](https://gocover.io/github.com/wzshiming/romanumeral)

## Install

``` shell
go get -u -v github.com/wzshiming/romanumeral
```

## Example

``` golang
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
			t.Fatal(int(i), d, i, off, tmp)
		}
	}
}
```

## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/romanumeral/blob/master/LICENSE) for the full license text.
