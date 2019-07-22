package main

import (
	"testing"
)

func Test_ABinary_encode(t *testing.T) {
	value := "31323334"
	expected := []byte{0x31, 0x32, 0x33, 0x34}
	codec := ABinary{"", "Should be [0x31, 0x32, 0x33, 0x34]", 4}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_ABinary_encodeMac(t *testing.T) {
	value := "2D2A98F12D2A98F1"
	expected := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1}
	codec := ABinary{"", "Should be [0x2d, 0x2a, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1]", 8}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_ABinary_decode(t *testing.T) {
	value := []byte{0x31, 0x32, 0x33, 0x34}
	expected := "31323334"
	codec := ABinary{"", "Should be 31323334", 4}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}

func Test_ABinary_decodeMac(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1}
	expected := "2D2A98F12D2A98F1"
	codec := ABinary{"", "Should be 2D2A98F12D2A98F1", 8}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}