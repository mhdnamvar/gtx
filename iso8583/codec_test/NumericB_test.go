package codec_test

import (
	"../../iso8583"
	"../codec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericB_Encode(t *testing.T) {
	value := "0123456789"
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	c := codec.DefaultNumericB(5)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Encode_LeftPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	c := codec.DefaultNumericB(5)
	c.PaddingType = codec.LeftPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Encode_RightPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x012, 0x34, 0x56, 0x78, 0x90}
	c := codec.DefaultNumericB(5)
	c.PaddingType = codec.RightPadding
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Encode_InvalidLen(t *testing.T) {
	value := "123456789"
	c := codec.DefaultNumericB(5)
	actual, err := c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	c = codec.DefaultNumericB(4)
	actual, err = c.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestNumericB_Decode(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	expected := "0123456789"
	c := codec.DefaultNumericB(5)
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	c := codec.DefaultNumericB(6)
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestNumericB_Decode_LeftPad(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	expected := "0123456789"
	c := codec.DefaultNumericB(5)
	c.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	c := codec.DefaultNumericB(6)
	c.PaddingType = codec.LeftPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestNumericB_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	c := codec.DefaultNumericB(6)
	c.PaddingType = codec.RightPadding
	actual, _, err := c.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
