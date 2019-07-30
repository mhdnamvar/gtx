package main

import (
	"testing"
)

func Test_LLAChar_Encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x30, 0x34, 0x41, 0x42, 0x43, 0x44}
	codec := LLACharNew("", "Should be '04ABCD'", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLAChar_EncodePad(t *testing.T) {
	value := "ABCD"
	expected := []byte("11ABCD       ")
	codec := LLACharNew("", "Should be '11ABCD       '", 11, true)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLAChar_InvalidLen(t *testing.T) {
	value := "ABCDEFGHIJKL"
	codec := LLACharNew("", "Should return error", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLAChar_Decode(t *testing.T) {
	value := []byte("04ABCD       ")
	expected := "ABCD"
	codec := LLACharNew("", "Should be 'ABCD'", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLAChar_DecodeInvalidLen(t *testing.T) {
	value := []byte("10ABCD")
	codec := LLACharNew("", "Should return error", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLAChar_DecodePad(t *testing.T) {
	value := []byte("10ABCD          ")
	expected := "ABCD      "
	codec := LLACharNew("", "Should be 'ABCD      '", 10, true)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}