package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringE_Encode(t *testing.T) {
	value := "0123456789  ABCD"
	expected := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	c := codec.DefaultStringE(16)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringE_Encode_LeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4}
	c := codec.DefaultStringE(6)
	c.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringE_Encode_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40}
	c := codec.DefaultStringE(6)
	c.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringE_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	c := codec.DefaultStringE(10)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	c = codec.DefaultStringE(5)
	actual, err = c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestStringE_Decode(t *testing.T) {
	value := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	expected := "0123456789  ABCD"
	c := codec.DefaultStringE(16)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringE_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	c := codec.DefaultStringE(20)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestStringE_Decode_LeftPad(t *testing.T) {
	value := []byte{
		0x40, 0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xC1, 0xC2, 0xC3, 0xC4,
	}
	expected := "  0123456789A"
	c := codec.DefaultStringE(13)
	c.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringE_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8}
	c := codec.DefaultStringE(11)
	c.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestStringE_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	c := codec.DefaultStringE(11)
	c.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
