# romanumeral
Roman numerals in Go


[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/romanumeral)](https://goreportcard.com/report/github.com/wzshiming/romanumeral)
[![GoDoc](https://godoc.org/github.com/wzshiming/romanumeral?status.svg)](https://godoc.org/github.com/wzshiming/romanumeral)
[![GitHub license](https://img.shields.io/github/license/wzshiming/romanumeral.svg)](https://github.com/wzshiming/romanumeral/blob/master/LICENSE)

## Install

``` shell
go get -u -v github.com/wzshiming/romanumeral
```

## Example

``` golang

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

```

## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/romanumeral/blob/master/LICENSE) for the full license text.
