package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLLBStringB_Encode(t *testing.T) {
	value := "D2345678901234567890"
	expected := []byte{
		0x10,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	c := codec.DefaultLLBStringB(10)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBStringB_Encode_LeftPad(t *testing.T) {
	value := "D45678901234567890"
	expected := []byte{
		0x10,
		0x20,
		0xD4, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	c := codec.DefaultLLBStringB(10)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBStringB_Encode_RightPad(t *testing.T) {
	value := "E234567890123456"
	expected := []byte{
		0x09,
		0xE2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
		0x20,
	}
	c := codec.DefaultLLBStringB(9)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBStringB_Encode_InvalidLen(t *testing.T) {
	value := "D234567890123456"
	c := codec.DefaultLLBStringB(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLLBStringB_Encode_InvalidData(t *testing.T) {
	value := "D234567890123456MN"
	c := codec.DefaultLLBStringB(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLLBStringB_Decode(t *testing.T) {
	value := []byte{
		0x09,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78,
	}
	expected := "D23456789012345678"
	c := codec.DefaultLLBStringB(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBStringB_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0x08,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	c := codec.DefaultLLBStringB(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{
		0x11,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	c = codec.DefaultLLBStringB(9)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLLBStringB_Decode_LeftPad(t *testing.T) {
	value := []byte{
		0x09,
		0x20,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	expected := "20D234567890123456"
	c := codec.DefaultLLBStringB(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBStringB_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x41}
	c := codec.DefaultLLBStringB(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLLBStringB_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x41}
	c := codec.DefaultLLBStringB(9)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
