package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsoLANumericA_Encode_1(t *testing.T) {
	value := "123456789"
	expected := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	codec := DefaultIsoLANumericA(9)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLANumericA_Encode_LeftPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
	codec := DefaultIsoLANumericA(9)
	codec.Data.PaddingType = IsoLeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLANumericA_Encode_RightPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x30}
	codec := DefaultIsoLANumericA(9)
	codec.Data.PaddingType = IsoRightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLANumericA_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	codec := DefaultIsoLANumericA(9)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestIsoLANumericA_Decode(t *testing.T) {
	value := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	expected := "1234567AB"
	codec := DefaultIsoLANumericA(9)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLANumericA_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x34, 0x31, 0x32, 0x33, 0x34}
	codec := DefaultIsoLANumericA(5)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)

	value = []byte{0x34, 0x31, 0x32, 0x33}
	codec = DefaultIsoLANumericA(3)
	actual, _, err = codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestIsoLANumericA_Decode_LeftPad(t *testing.T) {
	value := []byte{0x39, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	expected := "  1234567"
	codec := DefaultIsoLANumericA(9)
	codec.Data.PaddingType = IsoLeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLANumericA_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	codec := DefaultIsoLANumericA(9)
	codec.Data.PaddingType = IsoLeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}

func TestIsoLANumericA_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
	codec := DefaultIsoLANumericA(9)
	codec.Data.PaddingType = IsoRightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidDataError], err)
	assert.Equal(t, "", actual)
}
