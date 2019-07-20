package main

import (
	"testing"
)

func Test_LLAChar_Encode(t *testing.T) {
	value := "ABCD"
	expected := []byte("04ABCD")
	codec := LLAChar{"", "Should be '04ABCD'", 11, false}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLAChar_EncodePad(t *testing.T) {
	value := "ABCD"
	expected := []byte("11ABCD       ")
	codec := LLAChar{"", "Should be '11ABCD       '", 11, true}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLAChar_InvalidLen(t *testing.T) {
	value := "ABCDEFGHIJKL"
	codec := LLAChar{"", "Should return error", 11, false}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLAChar_Decode(t *testing.T) {
	value := []byte("04ABCD       ")
	expected := "ABCD"
	codec := LLAChar{"", "Should be 'ABCD'", 10, false}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}

func Test_LLAChar_DecodeInvalidLen(t *testing.T) {
	value := []byte("10ABCD")
	codec := LLAChar{"", "Should return error", 10, false}
	actual, err := codec.Decode(value)
	checkDecodeError(t, actual, err, InvalidLengthError)
}

func Test_LLAChar_DecodePad(t *testing.T) {
	value := []byte("10ABCD          ")
	expected := "ABCD      "
	codec := LLAChar{"", "Should be 'ABCD      '", 10, true}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
