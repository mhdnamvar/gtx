package iso8583_test

import (
	. "./"
	"github.com/stretchr/testify/assert"
	"testing"
)

func StringB(size int) *IsoType{
	return &IsoType{
		Value: &IsoData{
			Encoding: IsoBinary,
			Min: size,
			Max: size,
			ContentType: IsoHexString,
			Padding: IsoNoPad,
		},
	}
}


func TestStringB_Encode(t *testing.T) {
	value := "2D2A98F12D2A9820"
	expected := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20}
	isoType := StringB(8)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Encode_LeftPad(t *testing.T) {
	value := "2D2A98F12D2A9820"
	expected := []byte{0x20, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20}
	isoType := StringB(9)
	isoType.Value.Padding = IsoLeftPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Encode_RightPad(t *testing.T) {
	value := "2D2A98F12D2A9820"
	expected := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x20}
	isoType := StringB(9)
	isoType.Value.Padding = IsoRightPad
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
func TestStringB_Encode_InvalidData(t *testing.T) {
	value := "This is an invalid binary text"
	isoType := StringB(15)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidData, err)
	assert.Equal(t, []byte(nil), actual)
}
func TestStringB_Encode_InvalidLen(t *testing.T) {
	value := "2D2A98F12D2A98"
	isoType := StringB(6)
	actual, err := isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)

	value = "2D2A98F12D2A98201"
	isoType = StringB(8)
	actual, err = isoType.Encode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, []byte(nil), actual)
}

func TestStringB_Decode(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x12, 0x34, 0x56}
	expected := "2D2A98F12D2A9820"
	isoType := StringB(8)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Decode_InvalidLen(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x12, 0x34, 0x56}
	isoType := StringB(12)
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestStringB_Decode_LeftPad(t *testing.T) {
	value := []byte{0x20, 0x20, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20}
	expected := "20202D2A98F12D2A9820"
	isoType := StringB(10)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Decode_RightPad(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0x20, 0x20, 0x20}
	expected := "2D2A98F12D2A98202020"
	isoType := StringB(10)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestStringB_Decode_LeftPad_InvalidLen(t *testing.T) {
	value := []byte{0x20, 0x20, 0x2D, 0x2A, 0x98}
	isoType := StringB(6)
	isoType.Value.Padding = IsoLeftPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}

func TestStringB_Decode_RightPad_InvalidLen(t *testing.T) {
	value := []byte{0x20, 0x20, 0x2D, 0x2A, 0x98}
	isoType := StringB(6)
	isoType.Value.Padding = IsoRightPad
	actual, _, err := isoType.Decode(value)
	assert.Equal(t, InvalidLength, err)
	assert.Equal(t, "", actual)
}
