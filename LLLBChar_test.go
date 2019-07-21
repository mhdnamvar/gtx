package main

import (
	"testing"
)

func Test_LLLBChar_Encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x00, 0x04, 0x41, 0x42, 0x43, 0x44}
	codec := LLLBChar{"", "Should be '004ABCD'", 11, false}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLLBChar_EncodePad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x00, 0x05, 0x41, 0x42, 0x43, 0x44, 0x20}
	codec := LLLBChar{"", "Should be '005ABCD '", 5, true}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLLBChar_InvalidLen(t *testing.T) {
	value := "ABCDEFGHIJKL"
	codec := LLLBChar{"", "Should return error", 11, false}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLBChar_Decode(t *testing.T) {
	value := []byte{0x00, 0x04, 0x41, 0x42, 0x43, 0x44, 0x45, 0x20}
	expected := "ABCD"
	codec := LLLBChar{"", "Should be 'ABCD'", 10, false}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}

func Test_LLLBChar_DecodeInvalidLen(t *testing.T) {
	value := []byte{0x00, 0x05, 0x41, 0x42, 0x43, 0x44}
	codec := LLLBChar{"", "Should return error", 10, false}
	actual, err := codec.Decode(value)
	checkDecodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLBChar_DecodePad(t *testing.T) {
	value := []byte{0x00, 0x05, 0x41, 0x42, 0x43, 0x44, 0x20}
	expected := "ABCD "
	codec := LLLBChar{"", "Should be 'ABCD '", 5, true}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
