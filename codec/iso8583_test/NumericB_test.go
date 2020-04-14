package iso8583_test

import (
	. "../iso8583"
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
	value := "0123456789"
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	isoType := NumericB(5)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBEncodeLeftPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	isoType := NumericB(5)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBEncodeRightPad(t *testing.T) {
	value := "123456789"
	expected := []byte{0x012, 0x34, 0x56, 0x78, 0x90}
	isoType := NumericB(5)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBEncodeInvalidLen(t *testing.T) {
	value := "123456789"
	isoType := NumericB(5)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)

	isoType = NumericB(4)
	actual, err = isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestNumericBDecode(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	expected := "0123456789"
	isoType := NumericB(5)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBDecodeInvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	isoType := NumericB(6)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestNumericBDecodeLeftPad(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	expected := "0123456789"
	isoType := NumericB(5)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestNumericBDecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	isoType := NumericB(6)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestNumericBDecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0x01, 0x23, 0x45, 0x67, 0x89}
	isoType := NumericB(6)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
