package main

import (
	"testing"
)

func Test_LLLAChar_Encode(t *testing.T) {
	value := "ABCD"
	expected := []byte("004ABCD")
	codec := LLLACharNew("", "Should be '004ABCD'", 100, false)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLLAChar_EncodePad(t *testing.T) {
	value := "ABCD"
	expected := []byte(`100ABCD                         ` +
		`                                                    ` +
		`                   `)
	codec := LLLACharNew("", "Should be 100ABCD with 96 trailing spaces ", 100, true)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLLAChar_InvalidLen(t *testing.T) {
	value := "ABCDEFGHIJKL"
	codec := LLLACharNew("", "Should return error", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLLAChar_Decode(t *testing.T) {
	value := []byte("004ABCD       ")
	expected := "ABCD"
	codec := LLLACharNew("", "Should be 'ABCD'", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLLAChar_DecodeInvalidLen(t *testing.T) {
	value := []byte("100ABCD")
	codec := LLLACharNew("", "Should return error", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLLAChar_DecodePad(t *testing.T) {
	value := []byte("010ABCD      1234abcdextra")
	expected := "ABCD      "
	codec := LLLACharNew("", "Should be 'ABCD      '", 10, true)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)

	value = []byte("010ABCD123456extra")
	expected = "ABCD123456"
	codec = LLLACharNew("", "Should be 'ABCD123456'", 10, true)
	actual, err = codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
