package iso8583_test

import (
	. "./"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NumericB(size int) *IsoType{
	return &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary,
			Min: size,
			Max: size,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	}
}

func TestNumericBEncode(t *testing.T) {
	value := "12345"
	expected := []byte{0x12, 0x34, 0x50}
	isoType := NumericB(5)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBEncodeLeftPad(t *testing.T) {
	value := "1234567"
	expected := []byte{0x00, 0x01, 0x23, 0x45, 0x67}
	isoType := NumericB(10)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)

	value = "210"
	expected = []byte{0x02, 0x10}
	isoType = NumericB(4)
	isoType.Value.Padding = IsoLeftPad
	actual, err = isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBEncodeRightPad(t *testing.T) {
	value := "1234567"
	expected := []byte{0x12, 0x34, 0x56, 0x70, 0x00}
	isoType := NumericB(10)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBEncodeInvalidLen(t *testing.T) {
	value := "1234567"
	isoType := NumericB(6)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestNumericBDecode(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0x11, 0x22}
	expected := "0123456"
	isoType := NumericB(7)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)

	expected = "01234567"
	isoType = NumericB(8)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBDecodeInvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0x11, 0x22}
	isoType := NumericB(15)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestNumericBDecodeLeftPad(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	expected := "12345"
	isoType := NumericB(5)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBDecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte{0x00, 0x00, 0x00, 0x67, 0x89, 0x11, 0x22}
	isoType := NumericB(15)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestNumericBDecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23}
	isoType := NumericB(6)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
