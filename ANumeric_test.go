package main

import (
	"testing"
)

func Test_ANumeric_encode(t *testing.T) {
	value := "12345"
	expected := []byte("0012345") //[]byte{0x30, 0x30, 0x31, 0x032, 0x33, 0x34, 0x35}
	codec := ANumericNew("", "Should be 30303132333435", 7)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_ANumeric_decode(t *testing.T) {
	value := []byte{0x30, 0x30, 0x31, 0x032, 0x33, 0x34, 0x35}
	expected := "0012345"
	codec := ANumericNew("", "Should be 12345", 7)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
