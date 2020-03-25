package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericB_Encode(t *testing.T) {
	value := "0123456789"
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	codec := DefaultNumericB(5)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Encode_LeftPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	codec := DefaultNumericB(5)
	codec.PaddingType = LeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Encode_RightPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x012, 0x34, 0x56, 0x78, 0x90}
	codec := DefaultNumericB(5)
	codec.PaddingType = RightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Encode_InvalidLen(t *testing.T) {
	value := "123456789"
	codec := DefaultNumericB(5)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	codec = DefaultNumericB(4)
	actual, err = codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestNumericB_Decode(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	expected := "0123456789"
	codec := DefaultNumericB(5)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	codec := DefaultNumericB(6)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestNumericB_Decode_LeftPad(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	expected := "0123456789"
	codec := DefaultNumericB(5)
	codec.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericB_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	codec := DefaultNumericB(6)
	codec.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestNumericB_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	codec := DefaultNumericB(6)
	codec.PaddingType = RightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
