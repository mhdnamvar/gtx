package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L1BStringB(size int) *IsoType{
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
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	}
}

func TestL1BStringBEncode(t *testing.T) {
	value := "2D34567890123456F1"
	expected := []byte{
		0x09,
		0x2D, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0xF1,
	}
	isoType := L1BStringB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BStringBEncodeLeftPad(t *testing.T) {
	value := "3D34567890123456"
	expected := []byte{
		0x09,
		0x20,
		0x3D, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	isoType := L1BStringB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BStringBEncodeRightPad(t *testing.T) {
	value := "B534567890123456"
	expected := []byte{
		0x09,
		0xB5, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
		0x20,
	}
	isoType := L1BStringB(9)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BStringBEncodeInvalidLen(t *testing.T) {
	value := "9F3456789012342D"
	isoType := L1BStringB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL1BStringBEncodeInvalidData(t *testing.T) {
	value := "D234567890123456MN"
	isoType := L1BStringB(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL1BStringBDecode(t *testing.T) {
	value := []byte{
		0x09,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x7F,
	}
	expected := "D2345678901234567F"
	isoType := L1BStringB(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BStringBDecodeInvalidLen(t *testing.T) {
	value := []byte{
		0x08,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x3E,
	}
	isoType := L1BStringB(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{
		0x11,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0xAB,
	}
	isoType = L1BStringB(9)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1BStringBDecodeLeftPad(t *testing.T) {
	value := []byte{
		0x09,
		0x20,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x9F,
	}
	expected := "20123456789012349F"
	isoType := L1BStringB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1BStringBDecodeLeftPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x41}
	isoType := L1BStringB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1BStringBDecodeRightPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x43}
	isoType := L1BStringB(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
