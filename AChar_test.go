package main

import (
	"fmt"
	"testing"
)

func Test_AChar_encode(t *testing.T) {
	value := "ABCD"
	expected := []byte("ABCD   ")
	fmt.Printf("%x\n", expected)
	codec := AChar{"", "Should be 'ABCD   '", 7}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_AChar_decode(t *testing.T) {
	value := []byte{0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	expected := "ABCD   "
	codec := AChar{"", "Should be 'ABCD   '", 7}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
