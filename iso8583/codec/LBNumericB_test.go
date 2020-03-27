package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLBNumericB_Encode(t *testing.T) {
	value := "12345678901234567890"
	expected := []byte{
		0x10,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	codec := DefaultLBNumericB(10)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBNumericB_Encode_LeftPad(t *testing.T) {
	value := "12345678901234567890"
	expected := []byte{
		0x11,
		0x00,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	codec := DefaultLBNumericB(11)
	codec.Data.PaddingType = LeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBNumericB_Encode_RightPad(t *testing.T) {
	value := "12345678901234567890"
	expected := []byte{
		0x11,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x00,
	}
	codec := DefaultLBNumericB(11)
	codec.Data.PaddingType = RightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBNumericB_Encode_InvalidLen(t *testing.T) {
	value := "1234567890123456789"
	codec := DefaultLBNumericB(11)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	value = "1234567890123456789012345678901234567890123456789" +
		"012345678901234567890123456789012345678901234567890"
	codec = DefaultLBNumericB(99)
	actual, err = codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLBNumericB_Encode_InvalidData(t *testing.T) {
	value := "123456789012345678MN"
	codec := DefaultLBNumericB(10)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.NumberFormatError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestLBNumericB_Decode(t *testing.T) {
	value := []byte{
		0x10,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	expected := "12345678901234567890"
	codec := DefaultLBNumericB(10)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBNumericB_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0x09,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78,
	}
	codec := DefaultLBNumericB(10)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{
		0x11,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	codec = DefaultLBNumericB(10)
	actual, _, err = codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestLBNumericB_Decode_LeftPad(t *testing.T) {
	value := []byte{
		0x11,
		0x00,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	expected := "0012345678901234567890"
	codec := DefaultLBNumericB(11)
	codec.Data.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestLBNumericB_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	codec := DefaultLBNumericB(9)
	codec.Data.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestLBNumericB_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	codec := DefaultLBNumericB(9)
	codec.Data.PaddingType = RightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
