# bytesort

[![GoDoc](https://godoc.org/github.com/nochso/bytesort?status.svg)](https://godoc.org/github.com/nochso/bytesort)
[![Build Status](https://travis-ci.org/nochso/bytesort.svg?branch=master)](https://travis-ci.org/nochso/bytesort)
[![Coverage Status](https://coveralls.io/repos/github/nochso/bytesort/badge.svg?branch=master)](https://coveralls.io/github/nochso/bytesort?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nochso/bytesort)](https://goreportcard.com/report/github.com/nochso/bytesort)

Package bytesort encodes common Go types as binary/byte slices that are bytewise
sortable.

The output is intended for binary/bytewise comparison and sorting.
More specifically for creating the keys used in indexes of key value stores.

- [Install](#install)
- [Usage](#usage)
	- [Output example](#output-example)
- [Change log and versioning](#change-log-and-versioning)
- [License](#license)

## Install

```sh
go get github.com/nochso/bytesort
```

## Usage

[Full documentation is available at godoc.org](https://godoc.org/github.com/nochso/bytesort).

### Output example

```go
vv := []interface{}{
	"abc",
	int16(math.MinInt16),
	int16(0),
	int16(math.MaxInt16),
	false,
	true,
}
for _, v := range vv {
	b, err := bytesort.Encode(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("% 8X %-6T %#v\n", b, v, v)
}
// Output:
// 61 62 63 string "abc"
//    00 00 int16  -32768
//    80 00 int16  0
//    FF FF int16  32767
//       00 bool   false
//       01 bool   true
```

Use [bytes.Compare](https://godoc.org/bytes#Compare) and
[sort.Slice](https://godoc.org/sort#Slice) to sort a slice of `[]byte`:

```go
input := [][]byte{ ... }
sort.Slice(s, func(i, j int) bool {
	return bytes.Compare(s[i], s[j]) < 0
})
```

Using [sort.Sort](https://godoc.org/sort#Sort) on structs that implement
[sort.Interface](https://godoc.org/sort#Interface) might be faster.

[sort.Search](https://godoc.org/sort#Search) might also be of interest.

## Change log and versioning

This project adheres to [Semantic Versioning](http://semver.org/).

See the [CHANGELOG](CHANGELOG.md) for a full history of releases.

## License

[MIT](LICENSE).
