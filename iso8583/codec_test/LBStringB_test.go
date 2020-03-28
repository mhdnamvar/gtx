package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLBStringB_Encode(t *testing.T) {
	value := "2D34567890123456F1"
	expected := []byte{
		0x09,
		0x2D, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0xF1,
	}
	c := codec.DefaultLBStringB(9)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBStringB_Encode_LeftPad(t *testing.T) {
	value := "3D34567890123456"
	expected := []byte{
		0x09,
		0x20,
		0x3D, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	c := codec.DefaultLBStringB(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBStringB_Encode_RightPad(t *testing.T) {
	value := "B534567890123456"
	expected := []byte{
		0x09,
		0xB5, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
		0x20,
	}
	c := codec.DefaultLBStringB(9)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBStringB_Encode_InvalidLen(t *testing.T) {
	value := "9F3456789012342D"
	c := codec.DefaultLBStringB(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLBStringB_Encode_InvalidData(t *testing.T) {
	value := "D234567890123456MN"
	c := codec.DefaultLBStringB(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLBStringB_Decode(t *testing.T) {
	value := []byte{
		0x09,
		0xD2, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x7F,
	}
	expected := "D2345678901234567F"
	c := codec.DefaultLBStringB(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBStringB_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0x08,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x3E,
	}
	c := codec.DefaultLBStringB(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{
		0x11,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0xAB,
	}
	c = codec.DefaultLBStringB(9)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLBStringB_Decode_LeftPad(t *testing.T) {
	value := []byte{
		0x09,
		0x20,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x9F,
	}
	expected := "20123456789012349F"
	c := codec.DefaultLBStringB(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBStringB_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x41}
	c := codec.DefaultLBStringB(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLBStringB_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x43}
	c := codec.DefaultLBStringB(9)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
