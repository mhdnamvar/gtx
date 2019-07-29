package main

import (
	"testing"
)

func Test_BBinary_encode(t *testing.T) {
	value := "1234B1"
	expected := []byte{0x12, 0x34, 0xB1}
	codec := BBinary{"", "Should be [0x12, 0x34, 0xB1]", 3}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BBinary_encodeMac(t *testing.T) {
	value := "2D2A98F12D2A98F1"
	expected := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1}
	codec := BBinary{"", "Should be [0x2d, 0x2a, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1]", 8}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BBinary_encodeMacInvalidLen(t *testing.T) {
	value := "2D2A98F12D2A98F1"
	codec := BBinary{"", "Should return invalid length error", 16}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_BBinary_encodeMacInvalidHexLen(t *testing.T) {
	value := "2D2A98F12D2A98F"
	codec := BBinary{"", "Should return invalid length error", 8}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_BBinary_decode(t *testing.T) {
	value := []byte{0x12, 0x34, 0xB1}
	expected := "1234B1"
	codec := BBinary{"", "Should be 1234B1", 3}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BBinary_decodeMac(t *testing.T) {
	value := []byte{0x2D, 0x2A, 0x98, 0xF1, 0x2D, 0x2A, 0x98, 0xF1}
	expected := "2D2A98F12D2A98F1"
	codec := BBinary{"", "Should be 2D2A98F12D2A98F1", 8}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BBinary_decodeIgnoreExtraBytes(t *testing.T) {
	value := []byte{0x12, 0x34, 0xB1, 0x34, 0xB1, 0x34, 0xB1}
	expected := "1234B1"
	codec := BBinary{"", "Should be 1234B1", 3}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BBinary_decodeNotEnoughBytes(t *testing.T) {
	value := []byte{0x12, 0x34}
	codec := BBinary{"", "Should return invalid length error", 3}
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}