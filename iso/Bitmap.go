package main

import (
	"encoding/hex"
	"strings"
)

// Bitmap ...
type Bitmap [16]byte

// IsSet ...
func (bits *Bitmap) IsSet(i int) bool {
	i--
	return bits[i/8]&(1<<uint(7-i%8)) != 0
}

// Set ...
func (bits *Bitmap) Set(i int) {
	i--
	bits[i/8] |= 1 << uint(7-i%8)
	if i > 64 {
		bits.Set(1)
	}
}

// Clear ...
func (bits *Bitmap) Clear(i int) {
	i--
	bits[i/8] &^= 1 << uint(7-i%8)
}

// Sets ...
func (bits *Bitmap) Sets(xs ...int) {
	for _, x := range xs {
		if x > 64 {
			bits.Set(1)
		}
		bits.Set(x)
	}
}

// Decode ...
func (bits *Bitmap) Decode(s string) ([]int, error) {
	if len(s)%2 != 0 {
		return nil, Errors[InvalidLengthError]
	}
	
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, Errors[InvalidDataError]
	}
	if len(b) > 16 {
		return nil, Errors[InvalidLengthError]
	}
	
	for i := 0; i < len(b); i++ {
		bits[i] = b[i]
	}
	return bits.Array(), nil
}

// Encode ...
func (bits *Bitmap) String() string {
	if bits.IsSet(1) {
		return strings.ToUpper(hex.EncodeToString(bits[:]))
	}
	return strings.ToUpper(hex.EncodeToString(bits[:8]))
}

// Array ...
func (bits *Bitmap) Array() []int {
	var array []int
	length := 64
	if bits.IsSet(1) {
		length = 128
	}
	for i := 1; i <= length; i++ {
		if bits.IsSet(i) {
			array = append(array, i)
		}
	}
	return array
}