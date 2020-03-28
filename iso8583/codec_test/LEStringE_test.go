package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLEStringE_Encode(t *testing.T) {
	value := "1234567AB"
	expected := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	c := codec.DefaultLEStringE(9)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLEStringE_Encode_LeftPad(t *testing.T) {
	value := "ABC3D"
	expected := []byte{0xF7, 0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xF3, 0xC4}
	c := codec.DefaultLEStringE(7)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLEStringE_Encode_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	c := codec.DefaultLEStringE(7)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLEStringE_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	c := codec.DefaultLEStringE(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLEStringE_Decode(t *testing.T) {
	value := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	expected := "1234567AB"
	c := codec.DefaultLEStringE(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLEStringE_Decode_InvalidLen(t *testing.T) {
	value := []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1}
	c := codec.DefaultLEStringE(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF9, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1}
	c = codec.DefaultLEStringE(8)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLEStringE_Decode_LeftPad(t *testing.T) {
	value := []byte{0xF9, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7}
	expected := "  1234567"
	c := codec.DefaultLEStringE(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLEStringE_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	c := codec.DefaultLEStringE(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLEStringE_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	c := codec.DefaultLEStringE(9)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
