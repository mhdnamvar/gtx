package iso8583_test

import (
	. "./"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L1ENumericE(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoEbcdic,
			Min: 1,
			Max: 1,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
		Value: &IsoData{
			Encoding: IsoEbcdic,
			Min: 0,
			Max: size,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	}
}

func TestL1ENumericEEncode(t *testing.T) {
	value := "123456789"
	expected := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	isoType := L1ENumericE(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1ENumericEEncodeLeftPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0xF8, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8}
	isoType := L1ENumericE(9)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1ENumericEEncodeRightPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0xF8, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF0}
	isoType := L1ENumericE(9)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1ENumericEEncodeInvalidLen(t *testing.T) {
	value := "iso8583"
	isoType := L1ENumericE(6)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL1ENumericEDecode(t *testing.T) {
	value := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	expected := "1234567AB"
	isoType := L1ENumericE(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1ENumericEDecodeInvalidLen(t *testing.T) {
	value := []byte{0xF4, 0xF1, 0xF2, 0xF3}
	isoType := L1ENumericE(5)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF4, 0xF1, 0xF2, 0xF3}
	isoType = L1ENumericE(3)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1ENumericEDecodeLeftPad(t *testing.T) {
	value := []byte{0xF9, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7}
	expected := "  1234567"
	isoType := L1ENumericE(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1ENumericEDecodeLeftPadInvalidData(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	isoType := L1ENumericE(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1ENumericEDecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	isoType := L1ENumericE(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
