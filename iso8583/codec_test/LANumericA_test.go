package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLANumericA_Encode(t *testing.T) {
	value := "123456789"
	expected := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	c := codec.DefaultLANumericA(9)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLANumericA_Encode_LeftPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
	c := codec.DefaultLANumericA(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLANumericA_Encode_RightPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x30}
	c := codec.DefaultLANumericA(9)
	c.Data.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLANumericA_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	c := codec.DefaultLANumericA(9)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLANumericA_Decode(t *testing.T) {
	value := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	expected := "1234567AB"
	c := codec.DefaultLANumericA(9)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLANumericA_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x34, 0x31, 0x32, 0x33, 0x34}
	c := codec.DefaultLANumericA(5)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0x34, 0x31, 0x32, 0x33}
	c = codec.DefaultLANumericA(3)
	actual, _, err = c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLANumericA_Decode_LeftPad(t *testing.T) {
	value := []byte{0x39, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	expected := "  1234567"
	c := codec.DefaultLANumericA(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLANumericA_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	c := codec.DefaultLANumericA(9)
	c.Data.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLANumericA_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
	c := codec.DefaultLANumericA(9)
	c.Data.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
