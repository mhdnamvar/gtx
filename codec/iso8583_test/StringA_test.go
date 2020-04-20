package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func StringA(size int) *IsoType{
	return &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii,
			Min: size,
			Max: size,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	}
}

func TestStringAEncode(t *testing.T) {
	value := "0123456789`~!#$%^*()+-=  ABCDE"
	expected := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44, 0x45,
	}
	isoType := StringA(30)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringAEncodeLeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte("   ABCD")
	isoType := StringA(7)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringAEncodeRightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte("ABCD      ")
	isoType := StringA(10)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringAEncodeInvalidLen(t *testing.T) {
	value := "iso8583"
	isoType := StringA(6)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)

	isoType = StringA(5)
	actual, err = isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestStringADecode(t *testing.T) {
	value := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x5F, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44,
	}
	expected := "0123456789`~!#$%^*()_+-=  ABCD"
	isoType := StringA(30)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringADecodeInvalidLen(t *testing.T) {
	value := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x60, 0x7E, 0x21, 0x23, 0x24, 0x25, 0x5E, 0x2A, 0x28, 0x29, 0x5F, 0x2B, 0x2D, 0x3D,
		0x20, 0x20, 0x41, 0x42, 0x43, 0x44,
	}
	isoType := StringA(31)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestStringADecodeLeftPad(t *testing.T) {
	value := []byte("   ABCDE01234--extra data--")
	expected := "   ABCDE01234"
	isoType := StringA(13)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringADecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte("   ABCDE10")
	isoType := StringA(11)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestStringADecodeRightPadInvalidLen(t *testing.T) {
	value := []byte("ABCDE10   ")
	isoType := StringA(11)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
