package iso8583_test

import (
	. "./"
	"github.com/stretchr/testify/assert"
	"testing"
)

func L2EStringE(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoEbcdic,
			Min: 2,
			Max: 2,
			ContentType: IsoNumeric,
			Padding: IsoLeftPad,
		},
		Value: &IsoData{
			Encoding: IsoEbcdic,
			Min: 0,
			Max: size,
			ContentType: IsoString,
			Padding: IsoNoPad,
		},
	}
}

func TestL2EStringE_Encode(t *testing.T) {
	value := "1234567890A"
	expected := []byte{0xF1, 0xF1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xF0, 0xC1}
	isoType := L2EStringE(11)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2EStringE_Encode_LeftPad(t *testing.T) {
	value := "12345678A"
	expected := []byte{0xF0, 0xF9, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xC1}
	isoType := L2EStringE(11)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2EStringE_Encode_RightPad(t *testing.T) {
	value := "1234567A"
	expected := []byte{0xF0, 0xF8, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0x40, 0x40}
	isoType := L2EStringE(10)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2EStringE_Encode_InvalidLen(t *testing.T) {
	value := "iso8583"
	isoType := L2EStringE(6)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestL2EStringE_Decode(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "1234567ABC"
	isoType := L2EStringE(10)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2EStringE_Decode_InvalidLen(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	isoType := L2EStringE(11)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)

	value = []byte{0xF1, 0xF1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5}
	isoType = L2EStringE(10)
	actual, _, err = isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2EStringE_Decode_LeftPad(t *testing.T) {
	value := []byte{0xF1, 0xF0, 0x40, 0x40, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0xC3}
	expected := "  34567ABC"
	isoType := L2EStringE(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestL2EStringE_Decode_LeftPad_InvalidData(t *testing.T) {
	value := []byte{0xC1, 0x40, 0x40, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2}
	isoType := L2EStringE(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestL2EStringE_Decode_RightPad_InvalidData(t *testing.T) {
	value := []byte{0xC1, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xC1, 0xC2, 0x40, 0x40}
	isoType := L2EStringE(10)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
