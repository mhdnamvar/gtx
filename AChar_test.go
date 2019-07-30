package main

import (
	"testing"	
)

func Test_AChar_Encode(t *testing.T) {
	value := "ABCD"
	expected := []byte("ABCD   ") // ]byte{0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	codec := ACharNew("", "Should be 'ABCD   '", 7)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_AChar_Decode(t *testing.T) {
	value := []byte{0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	expected := "ABCD   "
	codec := ACharNew("", "Should be 'ABCD   '", 7)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_AChar_Codec(t *testing.T) {
	codecName := "MTI"
	codec := ACharNew(codecName, "MESSAGE TYPE INDICATOR", 4)
	assertEqual(t, codecName, codec.Name)
}