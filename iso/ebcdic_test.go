package iso

import (
	"testing"
)

func Test_EncodeEbcdic(t *testing.T) {
	expected := []byte{0xd5, 0x81, 0x94, 0xa5, 0x81, 0x99}
	actual := AsciiToEbcdic("Namvar")
	assertEqual(t, expected, actual)
}

func Test_EncodeAscii(t *testing.T) {
	value := string([]byte{0xd5, 0x81, 0x94, 0xa5, 0x81, 0x99})
	expected := "Namvar"
	actual := string(EbcdicToAscii(value))
	assertEqual(t, expected, actual)
}
