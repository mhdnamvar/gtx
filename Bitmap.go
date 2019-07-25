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

func (bits Bitmap) String() string {
	if bits.IsSet(1) {
		return strings.ToUpper(hex.EncodeToString(bits[:]))
	}
	return strings.ToUpper(hex.EncodeToString(bits[8:]))
}
