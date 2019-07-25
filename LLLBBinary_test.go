package main

import (
	"testing"
)

func Test_LLLBBinary_EncodePad(t *testing.T) {
	value := "2D2A98F12D2A98"
	expected := []byte{0x00, 0x08, 0x00, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98}
	codec := LLLBBinary{"", "Should be [0x00, 0x08, 0x00, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98]", 8, true}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLLBBinary_EncodeNoPad(t *testing.T) {
	value := "2D2A98F12D2A98F1"
	expected := []byte{0x00, 0x08, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1}
	codec := LLLBBinary{"", "Should be [0x08, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1]", 11, false}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLLBBinary_EncodeInvalidLen(t *testing.T) {
	value := "2D2A98F12D"
	codec := LLLBBinary{"", "Should return error", 4, true}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLBBinary_EncodeWrongInput(t *testing.T) {
	value := "2D2A98F12"
	codec := LLLBBinary{"", "Should return error", 4, true}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLBBinary_Decode(t *testing.T) {
	value := []byte{0x00, 0x08, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1}
	expected := "2D2A98F12D2A98F1"
	codec := LLLBBinary{"", "Should be 2D2A98F12D2A98F1", 8, true}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}

func Test_LLLBBinary_DecodeError(t *testing.T) {
	value := []byte{0x00, 0x08, 0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98}
	codec := LLLBBinary{"", "Should return error", 10, true}
	actual, err := codec.Decode(value)
	checkDecodeError(t, actual, err, InvalidLengthError)
}
