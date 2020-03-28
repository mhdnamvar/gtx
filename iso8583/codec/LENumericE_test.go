package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLENumericE_Encode(t *testing.T) {
	value := "123456789"
	expected := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	codec := DefaultLENumericE(9)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLENumericE_Encode_LeftPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0xF9, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8}
	codec := DefaultLENumericE(9)
	codec.Data.PaddingType = LeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLENumericE_Encode_RightPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF0}
	codec := DefaultLENumericE(9)
	codec.Data.PaddingType = RightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLENumericE_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	codec := DefaultLENumericE(9)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLENumericE_Decode(t *testing.T) {
	value := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	expected := "1234567AB"
	codec := DefaultLENumericE(9)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLENumericE_Decode_InvalidLen(t *testing.T) {
	value := []byte{0xF4, 0xF1, 0xF2, 0xF3, 0xF4}
	codec := DefaultLENumericE(5)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF4, 0xF1, 0xF2, 0xF3}
	codec = DefaultLENumericE(3)
	actual, _, err = codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLENumericE_Decode_LeftPad(t *testing.T) {
	value := []byte{0xF9, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7}
	expected := "  1234567"
	codec := DefaultLENumericE(9)
	codec.Data.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLENumericE_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	codec := DefaultLENumericE(9)
	codec.Data.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLENumericE_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	codec := DefaultLENumericE(9)
	codec.Data.PaddingType = RightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
