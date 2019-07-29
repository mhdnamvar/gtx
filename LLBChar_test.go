package main

import (
	"testing"
)

func Test_LLBChar_Encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x04, 0x41, 0x42, 0x43, 0x44}
	codec := LLBChar{"", "Should be '04ABCD'", 11, false}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLBChar_EncodePad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x05, 0x41, 0x42, 0x43, 0x44, 0x20}
	codec := LLBChar{"", "Should be '05ABCD '", 5, true}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLBChar_InvalidLen(t *testing.T) {
	value := "ABCDEFGHIJKL"
	codec := LLBChar{"", "Should return error", 11, false}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLBChar_Decode(t *testing.T) {
	value := []byte{0x04, 0x41, 0x42, 0x43, 0x44, 0x45, 0x20}
	expected := "ABCD"
	codec := LLBChar{"", "Should be 'ABCD'", 10, false}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLBChar_DecodeInvalidLen(t *testing.T) {
	value := []byte{0x05, 0x41, 0x42, 0x43, 0x44}
	codec := LLBChar{"", "Should return error", 10, false}
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLBChar_DecodePad(t *testing.T) {
	value := []byte{0x05, 0x41, 0x42, 0x43, 0x44, 0x20}
	expected := "ABCD "
	codec := LLBChar{"", "Should be 'ABCD '", 5, true}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
