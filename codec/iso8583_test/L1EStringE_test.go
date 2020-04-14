package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L1EStringE(size int) *IsoType{
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

func TestL1EStringEEncode(t *testing.T) {
	value := "1234567AB"
	expected := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	isoType := L1EStringE(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1EStringEEncodeLeftPad(t *testing.T) {
	value := "ABC3D"
	expected := []byte{0xF7, 0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xF3, 0xC4}
	isoType := L1EStringE(7)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1EStringEEncodeRightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	isoType := L1EStringE(7)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1EStringEEncodeInvalidLen(t *testing.T) {
	value := "iso8583"
	isoType := L1EStringE(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL1EStringEDecode(t *testing.T) {
	value := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	expected := "1234567AB"
	isoType := L1EStringE(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1EStringEDecodeInvalidLen(t *testing.T) {
	value := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1}
	isoType := L1EStringE(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1}
	isoType = L1EStringE(8)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1EStringEDecodeLeftPad(t *testing.T) {
	value := []byte{0xF9, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7}
	expected := "  1234567"
	isoType := L1EStringE(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL1EStringEDecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	isoType := L1EStringE(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL1EStringEDecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	isoType := L1EStringE(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
