package main

import (
	"testing"
)

func Test_LLBNumeric_EncodePad(t *testing.T) {
	value := "12345"
	expected := []byte{0x11, 0x00, 0x00, 0x00, 0x01, 0x23, 0x45}
	codec := LLBNumericNew("", "Should be 1100000012345", 11, true)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLBNumeric_EncodeNoPad(t *testing.T) {
	value := "12345"
	expected := []byte{0x05, 0x01, 0x23, 0x45}
	codec := LLBNumericNew("", "Should be 05012345", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLBNumeric_EncodeInvalidLen(t *testing.T) {
	value := "12345"
	codec := LLBNumericNew("", "Should return error", 4, false)
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLBNumeric_EncodeWrongFormat(t *testing.T) {
	value := "12345ABC"
	codec := LLBNumericNew("", "Should return nil, error", 9, true)
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[NumberFormatError], err)
	assertEqual(t, nil, actual)
}

func Test_LLBNumeric_Decode(t *testing.T) {
	value := []byte{0x11, 0x00, 0x00, 0x00, 0x01, 0x23, 0x45}
	expected := "12345"
	codec := LLBNumericNew("", "Should be 12345", 11, true)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLBNumeric_DecodeError(t *testing.T) {
	value := []byte{0x10}
	codec := LLBNumericNew("", "Should return error", 10, true)
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)

	value = []byte{0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	expected := "1234567890"
	actual, err = codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)

	value = []byte{0x11, 0x12, 0x34, 0x56, 0x78, 0x90}
	actual, err = codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}
