package codec_test

import (
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitmapStringPrimary(t *testing.T) {
	var bitmap codec.IsoBitmap
	bitmap.Set(2, 12)
	assert.Equal(t, "4010000000000000", bitmap.String())
}

func TestBitmapArrayPrimary(t *testing.T) {
	var bitmap codec.IsoBitmap
	bitmap.Set(2, 12)
	assert.Equal(t, []int{2, 12}, bitmap.Array())
}

func TestBitmapParsePrimary(t *testing.T) {
	var bitmap codec.IsoBitmap
	err := bitmap.Parse("767F4601A8E1A20A")
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{2, 3, 4, 6, 7, 10, 11, 12, 13, 14, 15, 16, 18,
		22, 23, 32, 33, 35, 37, 41, 42, 43, 48, 49, 51, 55, 61, 63}, bitmap.Array())
}

func TestBitmapStringSecondary(t *testing.T) {
	var bitmap codec.IsoBitmap
	bitmap.Set(2, 3, 23, 36, 64, 65, 90, 128)
	assert.Equal(t, "E0000200100000018000004000000001", bitmap.String())
}

func TestBitmapArraySecondary(t *testing.T) {
	var bitmap codec.IsoBitmap
	bitmap.Set(2, 3, 23, 36, 64, 65, 90, 128)
	assert.Equal(t, []int{1, 2, 3, 23, 36, 64, 65, 90, 128}, bitmap.Array())
}

func TestBitmapParseSecondary(t *testing.T) {
	var bitmap codec.IsoBitmap
	err := bitmap.Parse("E0000200100000018000004000000001")
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{1, 2, 3, 23, 36, 64, 65, 90, 128}, bitmap.Array())
}

func TestBitmapClear(t *testing.T) {
	var bitmap codec.IsoBitmap
	bitmap.Set(2, 3, 64, 65, 128)
	bitmap.Clear(64, 65, 128)
	assert.Equal(t, "6000000000000000", bitmap.String())
	bitmap.Clear(2, 3)
	assert.Equal(t, "0000000000000000", bitmap.String())
}
