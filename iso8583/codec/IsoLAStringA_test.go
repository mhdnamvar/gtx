package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsoLAStringA_Encode(t *testing.T) {
	value := "1234567AB"
	expected := []byte{0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	codec := DefaultIsoLAStringA(9)
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLAStringA_Encode_LeftPad(t *testing.T) {
	value := "ABC3D"
	expected := []byte{0x35, 0x20, 0x20, 0x41, 0x42, 0x43, 0x33, 0x44}
	codec := DefaultIsoLAStringA(7)
	codec.Data.PaddingType = IsoLeftPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

/*
func TestIsoLAStringA_Encode_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte("ABCD      ")
	codec := DefaultIsoLAStringA(10)
	codec.PaddingType = IsoRightPadding
	actual, err := codec.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLAStringA_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	codec := DefaultIsoLAStringA(10)
	actual, err := codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)

	codec = DefaultIsoLAStringA(5)
	actual, err = codec.Encode(value)
	assert.Equal(t, iso8583.Errors[iso8583.InvalidLengthError], err)
	assert.Equal(t, []byte(nil), actual)
}

func TestIsoLAStringA_Decode(t *testing.T) {
	value := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x5F, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44,
	}
	expected := "0123456789`~!#$%^*()_+-=  ABCD"
	codec := DefaultIsoLAStringA(30)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLAStringA_Decode_InvalidLen(t *testing.T) {
	value := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x5F, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44,
	}
	codec := DefaultIsoLAStringA(31)
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestIsoLAStringA_Decode_LeftPad(t *testing.T) {
	value := []byte("   ABCDE01234--extra data--")
	expected := "   ABCDE01234"
	codec := DefaultIsoLAStringA(13)
	codec.PaddingType = IsoLeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestIsoLAStringA_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte("   ABCDE10")
	codec := DefaultIsoLAStringA(11)
	codec.PaddingType = IsoLeftPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}

func TestIsoLAStringA_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte("ABCDE10   ")
	codec := DefaultIsoLAStringA(11)
	codec.PaddingType = IsoRightPadding
	actual, _, err := codec.Decode(value)
	assert.Equal(t, iso8583.NotEnoughData, err)
	assert.Equal(t, "", actual)
}
*/
