package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)


func L2BNumericB(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary,
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
		Value: &IsoData{
			Encoding: IsoBinary,
			Min: 0,
			Max: size,
			ContentType: IsoNumeric,
			Padding: IsoNoPad,
		},
	}
}

func TestL2BNumericBEncode(t *testing.T) {
	value := "1234567890123456789"
	expected := []byte{
		0x19,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	isoType := L2BNumericB(20)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BNumericBEncodeLeftPad(t *testing.T) {
	value := "345678901234567890"
	expected := []byte{
		0x18,
		0x00,
		0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	isoType := L2BNumericB(20)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BNumericBEncodeRightPad(t *testing.T) {
	value := "1234567890123456"
	expected := []byte{
		0x16,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
		0x00, 0x00,
	}
	isoType := L2BNumericB(20)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BNumericBEncodeInvalidLen(t *testing.T) {
	value := "12345678901234567"
	isoType := L2BNumericB(16)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2BNumericBEncodeInvalidData(t *testing.T) {
	value := "12345678901234MN"
	isoType := L2BNumericB(16)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidData, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2BNumericBDecode(t *testing.T) {
	value := []byte{
		0x18,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78,
	}
	expected := "123456789012345678"
	isoType := L2BNumericB(18)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BNumericBDecodeInvalidLen(t *testing.T) {
	value := []byte{
		0x18,
		0x12, 0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56,
	}
	isoType := L2BNumericB(18)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

}

func TestL2BNumericBDecodeLeftPad(t *testing.T) {
	value := []byte{
		0x18,
		0x00,
		0x34, 0x56, 0x78, 0x90,
		0x12, 0x34, 0x56, 0x78, 0x90,
	}
	expected := "345678901234567890"
	isoType := L2BNumericB(20)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2BNumericBDecodeLeftPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	isoType := L2BNumericB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestLLBNumericBDecodeRightPadInvalidData(t *testing.T) {
	value := []byte{0x4D, 0x4E, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	isoType := L2BNumericB(9)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
