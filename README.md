# bytesort
[![GoDoc](https://godoc.org/github.com/nochso/bytesort?status.svg)](https://godoc.org/github.com/nochso/bytesort)
[![Build Status](https://travis-ci.org/nochso/bytesort.svg?branch=master)](https://travis-ci.org/nochso/bytesort)
[![Coverage Status](https://coveralls.io/repos/github/nochso/bytesort/badge.svg?branch=master)](https://coveralls.io/github/nochso/bytesort?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nochso/bytesort)](https://goreportcard.com/report/github.com/nochso/bytesort)

Package bytesort encodes common Go types as binary/byte slices that are bytewise sortable.


The output is intended for binary/bytewise comparison and sorting.
More specifically for creating the keys used in indexes in a bolt DB.

Use `bytes.Compare` and `sort.Slice` to sort [][]byte:

```go
sort.Slice(s, func(i, j int) bool {
	return bytes.Compare(s[i], s[j]) < 0
})
```
`sort.Search` might also be of interest.

## Install

```
go get github.com/nochso/bytesort
```

## Change log and versioning
This project adheres to [Semantic Versioning](http://semver.org/).

See the [CHANGELOG](CHANGELOG.md) for a full history of releases.

## License

[MIT](LICENSE).
