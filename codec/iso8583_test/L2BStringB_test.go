package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L2BStringB(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary,
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
		Value: &IsoData{
			Encoding: IsoBinary,
			Min: 0,
			Max: size,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	}
}

func TestL2BStringBEncode(t *testing.T) {
	value := "D2345678901234567890"
	expected := []byte{
		0x10,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	isoType := L2BStringB(10)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BStringBEncodeLeftPad(t *testing.T) {
	value := "D45678901234567890"
	expected := []byte{
		0x10,
		0x20,
		0xD4, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	isoType := L2BStringB(10)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BStringBEncodeRightPad(t *testing.T) {
	value := "E234567890123456"
	expected := []byte{
		0x09,
		0xE2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
		0x20,
	}
	isoType := L2BStringB(9)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BStringBEncodeInvalidLen(t *testing.T) {
	value := "D234567890123456"
	isoType := L2BStringB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2BStringBEncodeInvalidData(t *testing.T) {
	value := "D234567890123456MN"
	isoType := L2BStringB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2BStringBDecode(t *testing.T) {
	value := []byte{
		0x09,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78,
	}
	expected := "D23456789012345678"
	isoType := L2BStringB(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BStringBDecodeInvalidLen(t *testing.T) {
	value := []byte{
		0x08,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	isoType := L2BStringB(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{
		0x11,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	isoType = L2BStringB(9)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2BStringBDecodeLeftPad(t *testing.T) {
	value := []byte{
		0x09,
		0x20,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	expected := "20D234567890123456"
	isoType := L2BStringB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BStringBDecodeLeftPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x41}
	isoType := L2BStringB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2BStringBDecodeRightPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x41}
	isoType := L2BStringB(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
