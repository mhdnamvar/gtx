package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLLAStringA_Encode(t *testing.T) {
	value := "1234567890A"
	expected := []byte{0x31, 0x31, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x41}
	c := codec.DefaultLLAStringA(11)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLAStringA_Encode_LeftPad(t *testing.T) {
	value := "12345678A"
	expected := []byte{0x31, 0x31, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x41}
	c := codec.DefaultLLAStringA(11)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLAStringA_Encode_RightPad(t *testing.T) {
	value := "1234567A"
	expected := []byte{0x31, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x20, 0x20}
	c := codec.DefaultLLAStringA(10)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLAStringA_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	c := codec.DefaultLLAStringA(10)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLLAStringA_Decode(t *testing.T) {
	value := []byte{0x31, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x43}
	expected := "1234567ABC"
	c := codec.DefaultLLAStringA(10)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLAStringA_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x31, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x43}
	c := codec.DefaultLLAStringA(11)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0x31, 0x31, 0x31, 0x32, 0x33, 0x34, 0x35}
	c = codec.DefaultLLAStringA(10)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLLAStringA_Decode_LeftPad(t *testing.T) {
	value := []byte{0x31, 0x30, 0x20, 0x20, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x43}
	expected := "  34567ABC"
	c := codec.DefaultLLAStringA(10)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLLAStringA_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	c := codec.DefaultLLAStringA(10)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLLAStringA_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
	c := codec.DefaultLLAStringA(10)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
