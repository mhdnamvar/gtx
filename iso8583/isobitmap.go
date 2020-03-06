package iso8583

import (
	"encoding/hex"
	"strings"
)

const (
	BitmapSize = 16
)

type IsoBitmap [BitmapSize]byte

func (isoBitmap *IsoBitmap) Get(i int) bool {
	i--
	return isoBitmap[i/8]&(1<<uint(7-i%8)) != 0
}

func set(isoBitmap *IsoBitmap, i int) {
	i--
	isoBitmap[i/8] |= 1 << uint(7-i%8)
	if i > 64 {
		set(isoBitmap, 1)
	}
}

func clear(isoBitmap *IsoBitmap, i int) {
	i--
	isoBitmap[i/8] &^= 1 << uint(7-i%8)
}

func (isoBitmap *IsoBitmap) Set(xs ...int) {
	for _, x := range xs {
		if x > 64 {
			set(isoBitmap, 1)
		}
		set(isoBitmap, x)
	}
}
func (isoBitmap *IsoBitmap) Clear(xs ...int) {
	for _, x := range xs {
		clear(isoBitmap, x)
	}

	found := false
	for _, x := range isoBitmap.Array() {
		if x > 64 {
			found = true
			break
		}
	}

	if found {
		set(isoBitmap, 1)
	} else {
		clear(isoBitmap, 1)
	}
}

func (isoBitmap *IsoBitmap) Parse(s string) error {
	if len(s)%2 != 0 {
		return Errors[InvalidLengthError]
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return Errors[InvalidDataError]
	}
	if len(b) > BitmapSize {
		return Errors[InvalidLengthError]
	}

	for i := 0; i < len(b); i++ {
		isoBitmap[i] = b[i]
	}

	return nil
}

func (isoBitmap *IsoBitmap) String() string {
	if isoBitmap.Get(1) {
		return strings.ToUpper(hex.EncodeToString(isoBitmap[:]))
	}
	return strings.ToUpper(hex.EncodeToString(isoBitmap[:8]))
}

func (isoBitmap *IsoBitmap) Array() []int {
	var array []int
	length := 64
	if isoBitmap.Get(1) {
		length = 128
	}
	for i := 1; i <= length; i++ {
		if isoBitmap.Get(i) {
			array = append(array, i)
		}
	}
	return array
}
