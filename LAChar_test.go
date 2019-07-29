package main

import (
	"testing"
)

func Test_LAChar_encode(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x37, 0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	codec := LAChar{"", "Should be '7ABCD   '", 7, true}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LAChar_decode(t *testing.T) {
	value := []byte{0x37, 0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	expected := "ABCD   "
	codec := LAChar{"", "Should be 'ABCD   '", 7, true}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
