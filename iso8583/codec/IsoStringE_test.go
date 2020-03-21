package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsoStringE_Encode(t *testing.T) {
	value := "0123456789  ABCD"
	expected := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	codec := DefaultIsoStringE(16)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringE_Encode_LeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4}
	codec := DefaultIsoStringE(6)
	codec.PaddingType = IsoLeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringE_Encode_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40}
	codec := DefaultIsoStringE(6)
	codec.PaddingType = IsoRightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringE_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	codec := DefaultIsoStringE(10)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	codec = DefaultIsoStringE(5)
	actual, err = codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestIsoStringE_Decode(t *testing.T) {
	value := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	expected := "0123456789  ABCD"
	codec := DefaultIsoStringE(16)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringE_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0x40, 0x40, 0xC1, 0xC2, 0xC3, 0xC4,
	}
	codec := DefaultIsoStringE(20)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestIsoStringE_Decode_LeftPad(t *testing.T) {
	value := []byte{
		0x40, 0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9,
		0xC1, 0xC2, 0xC3, 0xC4,
	}
	expected := "  0123456789A"
	codec := DefaultIsoStringE(13)
	codec.PaddingType = IsoLeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringE_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8}
	codec := DefaultIsoStringE(11)
	codec.PaddingType = IsoLeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestIsoStringE_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
	codec := DefaultIsoStringE(11)
	codec.PaddingType = IsoRightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
