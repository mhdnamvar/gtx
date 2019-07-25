package main

import (
	"encoding/hex"
	"strings"
)

type Bitmap [16]byte

func (bits *Bitmap) IsSet(i int) bool { i -= 1; return bits[i/8]&(1<<uint(7-i%8)) != 0 }
func (bits *Bitmap) Set(i int)        { i -= 1; bits[i/8] |= 1 << uint(7-i%8) }
func (bits *Bitmap) Clear(i int)      { i -= 1; bits[i/8] &^= 1 << uint(7-i%8) }

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
	} else {
		return strings.ToUpper(hex.EncodeToString(bits[8:]))
	}
	
}
