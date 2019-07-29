package main

import (
	"testing"
)

func Test_LLANumeric_Encode(t *testing.T) {
	value := "12345"
	expected := []byte("0512345") //[]byte{0x30, 0x35, 0x31, 0x32, 0x33, 0x34, 0x35}
	codec := LLANumeric{"", "Should be 0512345", 11, false}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLANumeric_EncodePad(t *testing.T) {
	value := "12345"
	expected := []byte("1100000012345")
	codec := LLANumeric{"", "Should be 1100000012345", 11, true}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)

	expected = []byte("09000012345")
	codec = LLANumeric{"", "Should be 1100000012345", 9, true}
	actual, err = codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLANumeric_InvalidLen(t *testing.T) {
	value := "123456789012"
	codec := LLANumeric{"", "Should return error", 11, false}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLANumeric_InvalidFormat(t *testing.T) {
	value := "1234567890A"
	codec := LLANumeric{"", "Should return nil, error", 11, false}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[NumberFormatError], err)
	assertEqual(t, nil, actual)
}

func Test_LLANumeric_Decode(t *testing.T) {
	value := []byte("101234567890123456783E4B")
	expected := "1234567890"
	codec := LLANumeric{"", "Should be 1234567890", 10, false}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLANumeric_DecodeInvalidLen(t *testing.T) {
	value := []byte("10123456789")
	codec := LLANumeric{"", "Should return error", 10, false}
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLANumeric_DecodeInvalidFormat(t *testing.T) {
	value := []byte("1012345678lmnop456783E4B")
	codec := LLANumeric{"", "Should return error", 10, false}
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[NumberFormatError], err)
	assertEqual(t, nil, actual)
}

func Test_LLANumeric_DecodePad(t *testing.T) {
	value := []byte("100001234567890123456783E4B")
	expected := "0001234567"
	codec := LLANumeric{"", "Should be 0001234567", 10, true}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
