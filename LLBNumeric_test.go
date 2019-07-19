package main

import (
	"testing"
)

func Test_LLBNumeric_EncodePad(t *testing.T) {
	value := "12345"
	expected := []byte{0x11, 0x00, 0x00, 0x00, 0x01, 0x23, 0x45}
	codec := LLBNumeric{"", "Should be 1100000012345", 11, true}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLBNumeric_EncodeNoPad(t *testing.T) {
	value := "12345"
	expected := []byte{0x05, 0x01, 0x23, 0x45}
	codec := LLBNumeric{"", "Should be 05012345", 11, false}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLBNumeric_EncodeInvalidLen(t *testing.T) {
	value := "12345"
	codec := LLBNumeric{"", "Should return error", 4, false}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLBNumeric_EncodeWrongFormat(t *testing.T) {
	value := "12345ABC"
	codec := LLBNumeric{"", "Should return nil, error", 9, true}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, NumberFormatError)
}

func Test_LLBNumeric_Decode(t *testing.T) {
	value := []byte{0x11, 0x00, 0x00, 0x00, 0x01, 0x23, 0x45}
	expected := "12345"
	codec := LLBNumeric{"", "Should be 12345", 11, true}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
