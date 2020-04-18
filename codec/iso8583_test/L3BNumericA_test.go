package iso8583_test

import (
	. "../iso8583"
	"github.com/stretchr/testify/assert"
	"testing"
)

//IFB_LLLCHAR
func L3BStringA(size int) *IsoType{
	return &IsoType{
		Len: &IsoData{
			Encoding: IsoBinary,
			Min: 3,
			Max: 3,
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

func TestL3BStringAEncode(t *testing.T) {
	value := "673901002MN"
	expected := []byte{
		0x00, 0x11,
		0x36, 0x37, 0x33, 0x39, 0x30, 0x31, 0x30, 0x30, 0x32, 0x4D, 0x4E,
	}
	isoType := L3BStringA(104)
	actual, err := isoType.Encode(value)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}