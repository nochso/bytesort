// Package bytesort encodes common types as binary/byte slices that are bytewise sortable.
//
// The output is intended for binary/bytewise comparison and sorting.
// More specifically for creating the keys used in indexes in a bolt DB.
//
// Use bytes.Compare and sort.Slice to sort [][]byte:
//
//	sort.Slice(s, func(i, j int) bool {
//		return bytes.Compare(s[i], s[j]) < 0
//	})
//
// sort.Search might also be of interest.
package bytesort

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

// Encoder is the interface used for encoding custom types.
type Encoder interface {
	// EncodeSortable encodes the receiver as a byte slice in a byte sortable way.
	EncodeSortable() ([]byte, error)
}

// Encode a value as a byte slice that is bytewise/binary-sortable.
//
// Any results for the same type are sortable using a bytewise/binary
// comparison. A correct Sort order is not guaranteed when comparing different
// types.
//
// The length is the same for values of the same type. Except types string and
// []byte as they vary in length.
//
// Sortability is the only requirement. None of the encodings retain any type
// information because decoding of binary back into a value is out of scope.
//
// The following types are supported:
//
//	bool
//	float32 float64
//	int int8 int16 int32 int64
//	uint uint8 uint16 uint32 uint64
//	string    (case-sensitive)
//	time.Time (normalised to UTC)
//	[]byte    (copied)
func Encode(v interface{}) (b []byte, err error) {
	switch vv := v.(type) {
	case []byte:
		b := make([]byte, len(vv))
		copy(b, vv)
		return b, nil
	case string:
		return []byte(vv), nil
	case time.Time:
		return encodeTime(vv)
	case float64:
		return encodeFloat64(vv), nil
	case float32:
		return encodeFloat32(vv), nil
	case bool:
		if vv {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case int8:
		return []byte{byte(vv) ^ 0x80}, nil
	case uint8:
		return []byte{vv}, nil
	case int16:
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, uint16(vv))
		b[0] ^= 0x80
		return b, nil
	case uint16:
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, vv)
		return b, nil
	case int32:
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(vv))
		b[0] ^= 0x80
		return b, nil
	case uint32:
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, vv)
		return b, nil
	case int64:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(vv))
		b[0] ^= 0x80
		return b, nil
	case uint64:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, vv)
		return b, nil
	case int:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(vv))
		b[0] ^= 0x80
		return b, nil
	case uint:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(vv))
		return b, nil
	case Encoder:
		return vv.EncodeSortable()
	}
	return nil, fmt.Errorf("bytesort.Encode: unsupported type %T", v)
}

// http://stereopsis.com/radix.html
func encodeFloat64(v float64) []byte {
	bits := math.Float64bits(v)
	bits ^= -(bits >> 63) | (1 << 63)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, bits)
	return b
}

func encodeFloat32(v float32) []byte {
	bits := math.Float32bits(v)
	bits ^= -(bits >> 31) | (1 << 31)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, bits)
	return b
}

func encodeTime(v time.Time) ([]byte, error) {
	b, err := v.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("bytesort.Encode: %s", err)
	}
	// Strip version and time zone, leaving only the sort-relevant parts
	return b[1 : len(b)-2], nil
}
