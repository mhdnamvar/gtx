package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L2ANumericA(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoAscii,
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
		Value: &IsoData{
			Encoding: IsoAscii,
			Min: 0,
			Max: size,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	}
}

func TestL2NumericAEncode(t *testing.T) {
	value := "123456789"
	expected := []byte{0x30, 0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	isoType := L2ANumericA(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ANumericAEncodeLeftPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x30, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
	isoType := L2ANumericA(9)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ANumericAEncodeRightPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x30, 0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x30}
	isoType := L2ANumericA(9)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ANumericAEncodeInvalidData(t *testing.T) {
	value := "12iso8583"
	isoType := L2ANumericA(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidData, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2ANumericAEncodeInvalidLen(t *testing.T) {
	value := "8583"
	isoType := L2ANumericA(3)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2ANumericADecode(t *testing.T) {
	value := []byte{0x30, 0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	expected := "1234567AB"
	isoType := L2ANumericA(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ANumericADecodeInvalidLen(t *testing.T) {
	value := []byte{0x30, 0x36, 0x31, 0x32, 0x33, 0x34}
	isoType := L2ANumericA(6)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2ANumericADecodeLeftPad(t *testing.T) {
	value := []byte{0x30, 0x39, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	expected := "  1234567"
	isoType := L2ANumericA(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2ANumericADecodeLeftPadInvalidLen(t *testing.T) {
	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	isoType := L2ANumericA(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2ANumericADecodeRightPadInvalidLen(t *testing.T) {
	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
	isoType := L2ANumericA(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}