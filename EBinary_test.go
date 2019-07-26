package main

import (
	"testing"
)

func Test_EBinary_encode(t *testing.T) {
	value := "F1F2F3F4"
	expected := []byte{0xF1, 0xF2, 0xF3, 0xF4}
	codec := EBinary{"", "Should be [0xF1, 0xF2, 0xF3, 0xF4]", 4}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_EBinary_decodeMac(t *testing.T) {
	value := []byte{0xF1, 0xF2, 0xF3, 0xF4}
	expected := "F1F2F3F4"
	codec := EBinary{"", "Should be F1F2F3F4", 4}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}