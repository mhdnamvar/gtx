package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLAStringA_Encode(t *testing.T) {
	value := "1234567AB"
	expected := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	c := codec.DefaultLAStringA(9)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLAStringA_Encode_LeftPad(t *testing.T) {
	value := "ABC3D"
	expected := []byte{0x37, 0x20, 0x20, 0x41, 0x42, 0x43, 0x33, 0x44}
	c := codec.DefaultLAStringA(7)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLAStringA_Encode_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x37, 0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	c := codec.DefaultLAStringA(7)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLAStringA_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	c := codec.DefaultLAStringA(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLAStringA_Decode(t *testing.T) {
	value := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	expected := "1234567AB"
	c := codec.DefaultLAStringA(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLAStringA_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41}
	c := codec.DefaultLAStringA(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41}
	c = codec.DefaultLAStringA(8)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLAStringA_Decode_LeftPad(t *testing.T) {
	value := []byte{0x39, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	expected := "  1234567"
	c := codec.DefaultLAStringA(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLAStringA_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	c := codec.DefaultLAStringA(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLAStringA_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
	c := codec.DefaultLAStringA(9)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
