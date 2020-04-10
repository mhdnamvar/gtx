package isocodec_test

import (
	. "../isocodec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LaNumericA_Encode(t *testing.T) {
	value := "123456789"
	expected := []byte{0x30, 0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	isoType := LaNumericA(9)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_LaNumericA_Encode_LeftPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x30, 0x39, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
	isoType := LaNumericA(9)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_LaNumericA_Encode_RightPad(t *testing.T) {
	value := "12345678"
	expected := []byte{0x30, 0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x30}
	isoType := LaNumericA(9)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

//func Test_LaNumericA_Encode_1_InvalidLen(t *testing.T) {
//	value := "iso8583"
//	isoType := LaNumericA(9)
//	actual, err := isoType.Encode(value)
//	assert.Equal(t, InvalidLength, err)
//	assert.Equal(t, []byte(nil), actual)
//}

func Test_LaNumericA_Decode(t *testing.T) {
	value := []byte{0x30, 0x39, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	expected := "1234567AB"
	isoType := LaNumericA(9)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

//func Test_LaNumericA_Decode_InvalidLen(t *testing.T) {
//	value := []byte{0x34, 0x31, 0x32, 0x33, 0x34}
//	isoType := LaNumericA(5)
//	actual, _, err := isoType.Decode(value)
//	assert.Equal(t, NotEnoughData, err)
//	assert.Equal(t, "", actual)
//
//	value = []byte{0x34, 0x31, 0x32, 0x33}
//	isoType = LaNumericA(3)
//	actual, _, err = isoType.Decode(value)
//	assert.Equal(t, NotEnoughData, err)
//	assert.Equal(t, "", actual)
//}

func Test_LaNumericA_Decode_LeftPad(t *testing.T) {
	value := []byte{0x30, 0x39, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37}
	expected := "  1234567"
	isoType := LaNumericA(9)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

//func Test_LaNumericA_Decode_LeftPad_InvalidData(t *testing.T) {
//	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
//	isoType := LaNumericA(9)
//	isoType.Value.Padding = IsoLeftPad
//	actual, _, err := isoType.Decode(value)
//	assert.Equal(t, InvalidData, err)
//	assert.Equal(t, "", actual)
//}

//func Test_LaNumericA_Decode_RightPad_InvalidLen(t *testing.T) {
//	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
//	isoType := LaNumericA(9)
//	isoType.Value.Padding = IsoRightPad
//	actual, _, err := isoType.Decode(value)
//	assert.Equal(t, InvalidData, err)
//	assert.Equal(t, "", actual)
//}

func LaNumericA(size int) *IsoType{
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
