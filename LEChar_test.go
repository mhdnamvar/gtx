package main

import (
	"testing"
)

func Test_LEChar_encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	codec := LECharNew("", "Should be '7ABCD   '", 7, true)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LEChar_decode(t *testing.T) {
	value := []byte{0xF7, 0xC1, 0xC2, 0xC3, 0xC4, 0x40, 0x40, 0x40}
	expected := "ABCD   "
	codec := LECharNew("", "Should be 'ABCD   '", 7, true)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
