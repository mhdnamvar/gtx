package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericE_Encode(t *testing.T) {
	value := "0123456789012345678901234567890123456789"
	expected := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
	}
	codec := DefaultNumericE(40)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericE_Encode_LeftPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	codec := DefaultNumericE(10)
	codec.PaddingType = LeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericE_Encode_RightPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xF0}
	codec := DefaultNumericE(10)
	codec.PaddingType = RightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericE_Encode_InvalidLen(t *testing.T) {
	value := "123456789"
	codec := DefaultNumericE(10)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	codec = DefaultNumericE(8)
	actual, err = codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestNumericE_Decode(t *testing.T) {
	value := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
	}
	expected := "0123456789012345678901234567890123456789"
	codec := DefaultNumericE(40)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericE_Decode_InvalidLen(t *testing.T) {
	value := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	codec := DefaultNumericE(11)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestNumericE_Decode_LeftPad(t *testing.T) {
	value := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	expected := "0123456789"
	codec := DefaultNumericE(10)
	codec.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericE_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	codec := DefaultNumericE(11)
	codec.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestNumericE_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	codec := DefaultNumericE(11)
	codec.PaddingType = RightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
