package iso8583

import (
	"encoding/hex"
	"log"
	"strings"
)

const (
	BitmapSize = 16
)

type Bitmap [BitmapSize]byte

func (bitmap *Bitmap) Get(i int) bool {
	i--
	return bitmap[i/8]&(1<<uint(7-i%8)) != 0
}

func set(bitmap *Bitmap, i int) {
	i--
	bitmap[i/8] |= 1 << uint(7-i%8)
	if i > 64 {
		set(bitmap, 1)
	}
}

func clear(bitmap *Bitmap, i int) {
	i--
	bitmap[i/8] &^= 1 << uint(7-i%8)
}

func (bitmap *Bitmap) Set(xs ...int) {
	for _, x := range xs {
		if x > 64 {
			set(bitmap, 1)
		}
		set(bitmap, x)
	}
}

func (bitmap *Bitmap) Clear(xs ...int) {
	for _, x := range xs {
		clear(bitmap, x)
	}

	found := false
	for _, x := range bitmap.Array() {
		if x > 64 {
			found = true
			break
		}
	}

	if found {
		set(bitmap, 1)
	} else {
		clear(bitmap, 1)
	}
}

func (bitmap *Bitmap) Parse(s string) error {
	if len(s)%2 != 0 {
		log.Fatal("bitmap length is wrong")
		return InvalidLength
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		log.Printf("s=%s, b=%x", s, b)
		return InvalidData
	}
	if len(b) > BitmapSize {
		return InvalidData
	}

	for i := 0; i < len(b); i++ {
		bitmap[i] = b[i]
	}

	return nil
}

func (bitmap *Bitmap) String() string {
	if bitmap.Get(1) {
		return strings.ToUpper(hex.EncodeToString(bitmap[:]))
	}
	return strings.ToUpper(hex.EncodeToString(bitmap[:8]))
}

func (bitmap *Bitmap) Array() []int {
	var array []int
	length := 64
	if bitmap.Get(1) {
		length = 128
	}
	for i := 1; i <= length; i++ {
		if bitmap.Get(i) {
			array = append(array, i)
		}
	}
	return array
}
