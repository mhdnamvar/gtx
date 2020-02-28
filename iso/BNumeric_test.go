package main

import (
	"testing"
)

func Test_BNumeric_Encode(t *testing.T) {
	value := "12345"
	expected := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	codec := BNumericNew("", "Should be 000012345", 9)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BNumeric_InvalidFormat(t *testing.T) {
	value := "12345ABC"
	codec := BNumericNew("", "Should return nil, format error", 9)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, actual)
	assertEqual(t, Errors[NumberFormatError], err)
}
func Test_BNumeric_InvalidLen(t *testing.T) {
	value := "12345"
	codec := BNumericNew("", "Should return invalid length error", 4)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, actual)
	assertEqual(t, Errors[InvalidLengthError], err)
}
func Test_BNumeric_Decode(t *testing.T) {
	value := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	expected := "12345"
	codec := BNumericNew("", "Should be 12345", 9)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)

	codec = BNumericNew("", "Should be 1", 4)
	actual, err = codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
