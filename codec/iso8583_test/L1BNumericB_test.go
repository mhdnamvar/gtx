package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L1BNumericB(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary,
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
		Value: &IsoData{
			Encoding: IsoBinary,
			Min: 0,
			Max: size,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	}
}

func TestL1BNumericBEncode(t *testing.T) {
	value := "123456789012345678"
	expected := []byte{
		0x09,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78,
	}
	isoType := L1BNumericB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BNumericBEncodeLeftPad(t *testing.T) {
	value := "1234567890123456"
	expected := []byte{
		0x09,
		0x00,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	isoType := L1BNumericB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BNumericBEncodeRightPad(t *testing.T) {
	value := "1234567890123456"
	expected := []byte{
		0x09,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
		0x00,
	}
	isoType := L1BNumericB(9)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BNumericBEncodeInvalidLen(t *testing.T) {
	value := "1234567890123456"
	isoType := L1BNumericB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL1BNumericBEncodeInvalidData(t *testing.T) {
	value := "1234567890123456MN"
	isoType := L1BNumericB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidData, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL1BNumericBDecode(t *testing.T) {
	value := []byte{
		0x09,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78,
	}
	expected := "123456789012345678"
	isoType := L1BNumericB(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BNumericBDecodeInvalidLen(t *testing.T) {
	value := []byte{
		0x08,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	isoType := L1BNumericB(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{
		0x11,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	isoType = L1BNumericB(9)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1BNumericBDecodeLeftPad(t *testing.T) {
	value := []byte{
		0x09,
		0x00,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	expected := "001234567890123456"
	isoType := L1BNumericB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BNumericBDecodeLeftPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	isoType := L1BNumericB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1BNumericBDecodeRightPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	isoType := L1BNumericB(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
