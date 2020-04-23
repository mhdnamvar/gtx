package iso8583_test

import (
	. "./"
	"github.com/stretchr/testify/assert"
	"testing"
)

func StringE(size int) *IsoType{
	return &IsoType{
		Value: &IsoData{
			Encoding: IsoEbcdic,
			Min: size,
			Max: size,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	}
}

func TestStringEEncode(t *testing.T) {
	value := "0123456789  ABCD"
	expected := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	isoType := StringE(16)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringEEncodeLeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4}
	isoType := StringE(6)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringEEncodeRightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40}
	isoType := StringE(6)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringEEncodeInvalidLen(t *testing.T) {
	value := "iso8583"
	isoType := StringE(6)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)

	isoType = StringE(5)
	actual, err = isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestStringEDecode(t *testing.T) {
	value := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	expected := "0123456789  ABCD"
	isoType := StringE(16)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringEDecodeInvalidLen(t *testing.T) {
	value := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	isoType := StringE(20)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestStringEDecodeLeftPad(t *testing.T) {
	value := []byte{
		0x40, 0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xC1, 0xC2, 0xC3, 0xC4,
	}
	expected := "  0123456789A"
	isoType := StringE(13)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringEDecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte{0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8}
	isoType := StringE(11)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestStringEDecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	isoType := StringE(11)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
