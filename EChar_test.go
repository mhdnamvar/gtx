package main

import (
	"testing"
)

func Test_EChar_encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	codec := ECharNew("", "Should be 'ABCD   '", 7)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_EChar_decode(t *testing.T) {
	value := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	expected := "ABCD   "
	codec := ECharNew("", "Should be 'ABCD   '", 7)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
