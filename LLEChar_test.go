package main

import (
	"testing"
)

func Test_LLEChar_encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xF0, 0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	codec := LLECharNew("", "Should be '07ABCD   '", 7, true)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLEChar_decode(t *testing.T) {
	value := []byte{0xF0, 0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	expected := "ABCD   "
	codec := LLECharNew("", "Should be 'ABCD   '", 7, true)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
