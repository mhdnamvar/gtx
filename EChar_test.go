package main

import (
	"testing"
)

func Test_EChar_encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	codec := EChar{"", "Should be 'ABCD   '", 7}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_EChar_decode(t *testing.T) {
	value := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	expected := "ABCD   "
	codec := EChar{"", "Should be 'ABCD   '", 7}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
