package main

import (
	"testing"
)

func Test_ENumeric_encode(t *testing.T) {
	value := "12345"
	expected := []byte{0xF0, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5}
	codec := ENumeric{"", "Should be 0012345", 7}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_ENumeric_decode(t *testing.T) {
	value := []byte{0xF0, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5}
	expected := "12345"
	codec := ENumeric{"", "Should be 12345", 7}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
