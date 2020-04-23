package iso8583_test

import (
	. "./"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L2AStringA(size int) *IsoType{
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
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	}
}

func TestL2AStringA_Encode(t *testing.T) {
	value := "1234567890A"
	expected := []byte{0x31, 0x31, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30, 0x41}
	isoType := L2AStringA(11)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2AStringA_Encode_LeftPad(t *testing.T) {
	value := "12345678A"
	expected := []byte{0x30, 0x39, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x41}
	isoType := L2AStringA(11)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2AStringA_Encode_RightPad(t *testing.T) {
	value := "1234567A"
	expected := []byte{0x30, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x20, 0x20}
	isoType := L2AStringA(10)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2AStringA_Encode_InvalidLen(t *testing.T) {
	value := "8583"
	isoType := L2AStringA(3)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2AStringA_Decode(t *testing.T) {
	value := []byte{0x31, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x43}
	expected := "1234567ABC"
	isoType := L2AStringA(10)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2AStringA_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x31, 0x31, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x43}
	isoType := L2AStringA(11)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2AStringA_Decode_LeftPad(t *testing.T) {
	value := []byte{0x31, 0x30, 0x20, 0x20, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x43}
	expected := "  34567ABC"
	isoType := L2AStringA(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2AStringA_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0x41, 0x20, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42}
	isoType := L2AStringA(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2AStringA_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0x41, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x41, 0x42, 0x20, 0x20}
	isoType := L2AStringA(10)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
