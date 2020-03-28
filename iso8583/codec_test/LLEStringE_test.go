package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLLEStringE_Encode(t *testing.T) {
	value := "1234567890A"
	expected := []byte{0xF1, 0xF1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xF0, 0xC1}
	c := codec.DefaultLLEStringE(11)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Encode_LeftPad(t *testing.T) {
	value := "12345678A"
	expected := []byte{0xF1, 0xF1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xC1}
	c := codec.DefaultLLEStringE(11)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Encode_RightPad(t *testing.T) {
	value := "1234567A"
	expected := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0x40, 0x40}
	c := codec.DefaultLLEStringE(10)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	c := codec.DefaultLLEStringE(10)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLLEStringE_Decode(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "1234567ABC"
	c := codec.DefaultLLEStringE(10)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Decode_InvalidLen(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	c := codec.DefaultLLEStringE(11)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF1, 0xF1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5}
	c = codec.DefaultLLEStringE(10)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLLEStringE_Decode_LeftPad(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0x40, 0x40, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "  34567ABC"
	c := codec.DefaultLLEStringE(10)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLEStringE_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	c := codec.DefaultLLEStringE(10)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLLEStringE_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	c := codec.DefaultLLEStringE(10)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
