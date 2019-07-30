package main

import (
	"testing"
)

func Test_LLENumeric_Encode(t *testing.T) {
	value := "12345"
	expected := []byte{0xf0, 0xf5, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5}
	codec := LLENumericNew("", "Should be 0512345", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLENumeric_EncodePad(t *testing.T) {
	value := "12345"
	expected := []byte{0xf1, 0xf1, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5}
	codec := LLENumericNew("", "Should be 1100000012345", 11, true)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)

	expected = ASCIIToEbcdic("09000012345")
	// []byte{0xf0, 0xf9, 0xf0, 0xf0, 0xf0, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5}
	codec = LLENumericNew("", "Should be 0900000012345", 9, true)
	actual, err = codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLENumeric_EncodeInvalidLen(t *testing.T) {
	value := "123456789012"
	codec := LLENumericNew("", "Should return error", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLENumeric_EncodeInvalidFormat(t *testing.T) {
	value := "1234567890A"
	codec := LLENumericNew("", "Should return nil, error", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[NumberFormatError], err)
	assertEqual(t, nil, actual)
}

func Test_LLENumeric_Decode(t *testing.T) {
	value := []byte{0xf0, 0xf5, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5}
	expected := "12345"
	codec := LLENumericNew("", "Should be 12345", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLENumeric_DecodeInvalidLen(t *testing.T) {
	value := ASCIIToEbcdic("10123456789")
	codec := LLENumericNew("", "Should return error", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLENumeric_DecodeInvalidFormat(t *testing.T) {
	value := ASCIIToEbcdic("1012345678lmnop456783E4B")
	codec := LLENumericNew("", "Should return error", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[NumberFormatError], err)
	assertEqual(t, nil, actual)
}

func Test_LLENumeric_DecodePad(t *testing.T) {
	value := ASCIIToEbcdic("100001234567890123456783E4B")
	expected := "0001234567"
	codec := LLENumericNew("", "Should be 0001234567", 10, true)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
