package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLLBNumericB_Encode(t *testing.T) {
	value := "12345678901234567890"
	expected := []byte{
		0x10,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	c := codec.DefaultLLBNumericB(10)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBNumericB_Encode_LeftPad(t *testing.T) {
	value := "345678901234567890"
	expected := []byte{
		0x10,
		0x00,
		0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	c := codec.DefaultLLBNumericB(10)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBNumericB_Encode_RightPad(t *testing.T) {
	value := "1234567890123456"
	expected := []byte{
		0x09,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
		0x00,
	}
	c := codec.DefaultLLBNumericB(9)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBNumericB_Encode_InvalidLen(t *testing.T) {
	value := "1234567890123456"
	c := codec.DefaultLLBNumericB(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLLBNumericB_Encode_InvalidData(t *testing.T) {
	value := "1234567890123456MN"
	c := codec.DefaultLLBNumericB(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.NumberFormatError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLLBNumericB_Decode(t *testing.T) {
	value := []byte{
		0x09,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78,
	}
	expected := "123456789012345678"
	c := codec.DefaultLLBNumericB(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBNumericB_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0x08,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	c := codec.DefaultLLBNumericB(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{
		0x11,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	c = codec.DefaultLLBNumericB(9)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLLBNumericB_Decode_LeftPad(t *testing.T) {
	value := []byte{
		0x09,
		0x00,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	expected := "001234567890123456"
	c := codec.DefaultLLBNumericB(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLBNumericB_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	c := codec.DefaultLLBNumericB(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLLBNumericB_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	c := codec.DefaultLLBNumericB(9)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
