package main

import (
	"testing"
)

func Test_LLLAChar_Encode(t *testing.T) {
	value := "ABCD"
	expected := []byte("004ABCD")
	codec := LLLAChar{"", "Should be '004ABCD'", 100, false}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLLAChar_EncodePad(t *testing.T) {
	value := "ABCD"
	expected := []byte(`100ABCD                         ` +
		`                                                    ` +
		`                   `)
	codec := LLLAChar{"", "Should be 100ABCD with 96 trailing spaces ", 100, true}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLLAChar_InvalidLen(t *testing.T) {
	value := "ABCDEFGHIJKL"
	codec := LLLAChar{"", "Should return error", 11, false}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLAChar_Decode(t *testing.T) {
	value := []byte("004ABCD       ")
	expected := "ABCD"
	codec := LLLAChar{"", "Should be 'ABCD'", 10, false}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}

func Test_LLLAChar_DecodeInvalidLen(t *testing.T) {
	value := []byte("100ABCD")
	codec := LLLAChar{"", "Should return error", 10, false}
	actual, err := codec.Decode(value)
	checkDecodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLAChar_DecodePad(t *testing.T) {
	value := []byte("010ABCD      1234abcdextra")
	expected := "ABCD      "
	codec := LLLAChar{"", "Should be 'ABCD      '", 10, true}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)

	value = []byte("010ABCD123456extra")
	expected = "ABCD123456"
	codec = LLLAChar{"", "Should be 'ABCD123456'", 10, true}
	actual, err = codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
