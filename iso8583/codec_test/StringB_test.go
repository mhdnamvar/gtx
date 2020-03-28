package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringB_Encode(t *testing.T) {
	value := "2D2A98F12D2A9820"
	expected := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20}
	c := codec.DefaultStringB(8)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Encode_LeftPad(t *testing.T) {
	value := "2D2A98F12D2A9820"
	expected := []byte{0x20, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20}
	c := codec.DefaultStringB(9)
	c.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Encode_RightPad(t *testing.T) {
	value := "2D2A98F12D2A9820"
	expected := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x20}
	c := codec.DefaultStringB(9)
	c.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
func TestStringB_Encode_InvalidData(t *testing.T) {
	value := "This is an invalid binary text"
	c := codec.DefaultStringB(15)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, []byte(nil), actual)
}
func TestStringB_Encode_InvalidLen(t *testing.T) {
	value := "2D2A98F12D2A98"
	c := codec.DefaultStringB(8)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	c = codec.DefaultStringB(4)
	actual, err = c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	value = "2D2A98F12D2A98201"
	c = codec.DefaultStringB(8)
	actual, err = c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestStringB_Decode(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x12, 0x34, 0x56}
	expected := "2D2A98F12D2A9820"
	c := codec.DefaultStringB(8)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x12, 0x34, 0x56}
	c := codec.DefaultStringB(12)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestStringB_Decode_LeftPad(t *testing.T) {
	value := []byte{0x20, 0x20, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20}
	expected := "20202D2A98F12D2A9820"
	c := codec.DefaultStringB(10)
	c.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Decode_RightPad(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x20, 0x20}
	expected := "2D2A98F12D2A98202020"
	c := codec.DefaultStringB(10)
	c.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x20, 0x20, 0x2D, 0x2A, 0x98}
	c := codec.DefaultStringB(6)
	c.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestStringB_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0x20, 0x20, 0x2D, 0x2A, 0x98}
	c := codec.DefaultStringB(6)
	c.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
