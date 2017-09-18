# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

### Added

- Interface `Encoder` `EncodeSortable() ([]byte, error)`. It allows you to
  define the sortable encoding of your own types.
- Add godoc example of `Encode()` with output.

### Changed

- **BREAKING** `Encode()` no longer guarantees `len(x) > 0`.  
  Up to this point it was only relevant for strings. Empty strings used to be
  encoded as `0x00` to avoid empty slices for direct use as Bolt keys. Now an
  empty string is encoded as an empty byte slice.
- Improve speed by 30% on average by inlining code. See commit ab2cdb70 for
  details.

## 1.0.0 - 2017-06-03

### Added

- Split github.com/nochso/bolster/bytesort into a new separate repository.

[Unreleased]: https://github.com/nochso/bytesort/compare/1.0.0...HEAD