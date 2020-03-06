package iso8583

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsoBitmapStringPrimary(t *testing.T) {
	var bitmap IsoBitmap
	bitmap.Set(2, 12)
	assert.Equal(t, "4010000000000000", bitmap.String())
}

func TestIsoBitmapArrayPrimary(t *testing.T) {
	var bitmap IsoBitmap
	bitmap.Set(2, 12)
	assert.Equal(t, []int{2, 12}, bitmap.Array())
}

func TestIsoBitmapParsePrimary(t *testing.T) {
	var bitmap IsoBitmap
	err := bitmap.Parse("767F4601A8E1A20A")
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{2, 3, 4, 6, 7, 10, 11, 12, 13, 14, 15, 16, 18,
		22, 23, 32, 33, 35, 37, 41, 42, 43, 48, 49, 51, 55, 61, 63}, bitmap.Array())
}

func TestIsoBitmapStringSecondary(t *testing.T) {
	var bitmap IsoBitmap
	bitmap.Set(2, 3, 23, 36, 64, 65, 90, 128)
	assert.Equal(t, "E0000200100000018000004000000001", bitmap.String())
}

func TestIsoBitmapArraySecondary(t *testing.T) {
	var bitmap IsoBitmap
	bitmap.Set(2, 3, 23, 36, 64, 65, 90, 128)
	assert.Equal(t, []int{1, 2, 3, 23, 36, 64, 65, 90, 128}, bitmap.Array())
}

func TestIsoBitmapParseSecondary(t *testing.T) {
	var bitmap IsoBitmap
	err := bitmap.Parse("E0000200100000018000004000000001")
	assert.Equal(t, nil, err)
	assert.Equal(t, []int{1, 2, 3, 23, 36, 64, 65, 90, 128}, bitmap.Array())
}

func TestIsoBitmapClear(t *testing.T) {
	var bitmap IsoBitmap
	bitmap.Set(2, 3, 64, 65, 128)
	bitmap.Clear(64, 65, 128)
	assert.Equal(t, "6000000000000000", bitmap.String())
	bitmap.Clear(2, 3)
	assert.Equal(t, "0000000000000000", bitmap.String())
}
