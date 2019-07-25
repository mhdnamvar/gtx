package main

import (
	"testing"
)

func Test_Bitmap_Encode(t *testing.T) {
	var bits Bitmap
	bits.Sets(2, 3, 23, 36, 64, 65, 90, 128)
	assertEqual(t, "E0000200100000018000004000000001", bits.String())
}