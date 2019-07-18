package main

import (
	"testing"
)

func Test_BNumeric_Encode(t *testing.T) {
	value := "12345"
	expected := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	codec := BNumeric{"", "Should be 000012345", 9}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_BNumeric_InvalidFormat(t *testing.T) {
	value := "12345ABC"
	codec := BNumeric{"", "Should return nil, format error", 9}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, NumberFormatError)
}
func Test_BNumeric_InvalidLen(t *testing.T) {
	value := "12345"
	codec := BNumeric{"", "Should return invalid length error", 4}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}
func Test_BNumeric_Decode(t *testing.T) {
	value := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	expected := "12345"
	codec := BNumeric{"", "Should be 12345", 9}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)

	codec = BNumeric{"", "Should be 1", 4}
	actual, err = codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
