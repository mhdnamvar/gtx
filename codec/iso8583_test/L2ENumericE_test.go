package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L2ENumericE(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoEbcdic,
			Min: 2,
			Max: 2,
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

func TestL2ENumericEEncode(t *testing.T) {
	value := "1234567890"
	expected := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xF0}
	isoType := L2ENumericE(10)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ENumericEEncodeLeftPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0xF0, 0xF8, 0xF0, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8}
	isoType := L2ENumericE(10)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ENumericEEncodeRightPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0xF0, 0xF8, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF0, 0xF0}
	isoType := L2ENumericE(10)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ENumericEEncodeInvalidLen(t *testing.T) {
	value := "iso8583"
	isoType := L2ENumericE(6)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2ENumericEDecode(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "1234567ABC"
	isoType := L2ENumericE(10)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ENumericEDecodeInvalidLen(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	isoType := L2ENumericE(11)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF1, 0xF1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5}
	isoType = L2ENumericE(10)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2ENumericEDecodeLeftPad(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0x40, 0x40, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "  34567ABC"
	isoType := L2ENumericE(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ENumericEDecodeLeftPadInvalidData(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	isoType := L2ENumericE(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2ENumericEDecodeRightPadInvalidData(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	isoType := L2ENumericE(10)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
