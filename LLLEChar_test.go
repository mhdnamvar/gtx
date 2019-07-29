package main

import (
	"testing"
)

func Test_LLLEChar_encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xF0, 0xF0, 0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	codec := LLLEChar{"", "Should be '007ABCD   '", 7, true}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLLEChar_decode(t *testing.T) {
	value := []byte{0xF0, 0xF0, 0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	expected := "ABCD   "
	codec := LLLEChar{"", "Should be 'ABCD   '", 7, true}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
