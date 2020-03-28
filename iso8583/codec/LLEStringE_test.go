package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLLEStringE_Encode(t *testing.T) {
	value := "1234567890A"
	expected := []byte{0xF1, 0xF1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xF0, 0xC1}
	codec := DefaultLLEStringE(11)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Encode_LeftPad(t *testing.T) {
	value := "12345678A"
	expected := []byte{0xF1, 0xF1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xC1}
	codec := DefaultLLEStringE(11)
	codec.Data.PaddingType = LeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Encode_RightPad(t *testing.T) {
	value := "1234567A"
	expected := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0x40, 0x40}
	codec := DefaultLLEStringE(10)
	codec.Data.PaddingType = RightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	codec := DefaultLLEStringE(10)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLLEStringE_Decode(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "1234567ABC"
	codec := DefaultLLEStringE(10)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Decode_InvalidLen(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	codec := DefaultLLEStringE(11)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF1, 0xF1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5}
	codec = DefaultLLEStringE(10)
	actual, _, err = codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLLEStringE_Decode_LeftPad(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0x40, 0x40, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "  34567ABC"
	codec := DefaultLLEStringE(10)
	codec.Data.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	codec := DefaultLLEStringE(10)
	codec.Data.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLLEStringE_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	codec := DefaultLLEStringE(10)
	codec.Data.PaddingType = RightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
