package codec

import (
	"../../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsoStringA_Encode(t *testing.T) {
	value := "0123456789`~!#$%^*()_+-=  ABCD"
	expected := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x5F, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44,
	}
	codec := DefaultStringA(30)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringA_Encode_LeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte("   ABCD")
	codec := DefaultStringA(7)
	codec.PaddingType = LeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringA_Encode_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte("ABCD      ")
	codec := DefaultStringA(10)
	codec.PaddingType = RightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringA_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	codec := DefaultStringA(10)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	codec = DefaultStringA(5)
	actual, err = codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestIsoStringA_Decode(t *testing.T) {
	value := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x5F, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44,
	}
	expected := "0123456789`~!#$%^*()_+-=  ABCD"
	codec := DefaultStringA(30)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringA_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x5F, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44,
	}
	codec := DefaultStringA(31)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestIsoStringA_Decode_LeftPad(t *testing.T) {
	value := []byte("   ABCDE01234--extra data--")
	expected := "   ABCDE01234"
	codec := DefaultStringA(13)
	codec.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoStringA_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte("   ABCDE10")
	codec := DefaultStringA(11)
	codec.PaddingType = LeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestIsoStringA_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte("ABCDE10   ")
	codec := DefaultStringA(11)
	codec.PaddingType = RightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
