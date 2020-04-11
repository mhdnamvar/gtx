package isocodec_test

import (
	. "../isocodec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NumericA(size int) *IsoType{
	return &IsoType{
		Value: &IsoData{
			Encoding: IsoAscii,
			Min: size,
			Max: size,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	}
}

func TestNumericAEncode(t *testing.T) {
	value := "0123456789012345678901234567890123456789"
	expected := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
	}
	c := NumericA(40)
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericAEncodeLeftPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	c := NumericA(10)
	c.Value.Padding = IsoLeftPad
	actual, err := c.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericAEncodeRightPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30}
	isoType := NumericA(10)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericAEncode_InvalidLen(t *testing.T) {
	value := "123456789"
	isoType := NumericA(10)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)

	isoType = NumericA(8)
	actual, err = isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestNumericADecode(t *testing.T) {
	value := []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
	}
	expected := "0123456789012345678901234567890123456789"
	isoType := NumericA(40)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericADecodeInvalidLen(t *testing.T) {
	value := []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	isoType := NumericA(11)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestNumericADecodeLeftPad(t *testing.T) {
	value := []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	expected := "0123456789"
	isoType := NumericA(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericADecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	isoType := NumericA(11)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestNumericADecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	isoType := NumericA(11)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
