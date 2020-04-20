package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L1AStringA(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoAscii,
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
		Value: &IsoData{
			Encoding: IsoAscii,
			Min: 0,
			Max: size,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	}
}


func TestL1AStringAEncode(t *testing.T) {
	value := "1234567AB"
	expected := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	isoType := L1AStringA(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1AStringAEncodeLeftPad(t *testing.T) {
	value := "ABC3D"
	expected := []byte{0x35, 0x20, 0x20, 0x41, 0x42, 0x43, 0x33, 0x44}
	isoType := L1AStringA(7)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1AStringAEncodeRightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x34, 0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	isoType := L1AStringA(7)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1AStringAEncodeInvalidLen(t *testing.T) {
	value := "8583"
	isoType := L1AStringA(3)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL1AStringADecode(t *testing.T) {
	value := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	expected := "1234567AB"
	isoType := L1AStringA(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1AStringADecodeInvalidLen(t *testing.T) {
	value := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41}
	isoType := L1AStringA(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41}
	isoType = L1AStringA(8)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1AStringADecodeLeftPad(t *testing.T) {
	value := []byte{0x39, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	expected := "  1234567"
	isoType := L1AStringA(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1AStringADecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	isoType := L1AStringA(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1AStringADecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
	isoType := L1AStringA(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
