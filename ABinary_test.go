package main

import (
	"testing"
)

func Test_ABinary_Encode(t *testing.T) {
	value := "313233"
	expected := []byte{0x33, 0x31, 0x33, 0x32, 0x33, 0x33}
	codec := ABinaryNew("", "Should be [0x33, 0x31, 0x33, 0x32]", 3)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_ABinary_EncodeMac2(t *testing.T) {
	value := "2D2A98F12D2A98F1"
	expected := []byte{0x32, 0x44, 0x32, 0x41, 0x39, 0x38, 0x46, 0x31, 0x32, 0x44, 0x32, 0x41, 0x39, 0x38, 0x46, 0x31}
	codec := ABinaryNew("", "Should be [0x32, 0x44, 0x32, 0x41, 0x39, 0x38, 0x46, 0x31, 0x32, 0x44, 0x32, 0x41, 0x39, 0x38, 0x46, 0x31]", 8)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_ABinary_DecodeMac2(t *testing.T) {
	value := []byte{0x32, 0x44, 0x32, 0x41, 0x39, 0x38, 0x46, 0x31, 0x32, 0x44, 0x32, 0x41, 0x39, 0x38, 0x46, 0x31}
	expected := "2D2A98F12D2A98F1"
	codec := ABinaryNew("", "Should be 2D2A98F12D2A98F1", 4)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_ABinary_EncodePrimaryBitmap(t *testing.T) {
	value := "4000000000000000"
	expected := []byte("00000000000000004000000000000000")
	codec := ABinaryNew("", "Should be []byte(00000000000000004000000000000000)", 16)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_ABinary_EncodeSecondaryBitmap(t *testing.T) {
	value := "C0100000000000000000000000000001"
	expected := []byte("C0100000000000000000000000000001")
	codec := ABinaryNew("", "Should be []byte(C0100000000000000000000000000001)", 16)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_ABinary_DecodeSecondaryBitmap(t *testing.T) {
	value := []byte("C0100000000000000000000000000001")
	expected := "C0100000000000000000000000000001"
	codec := ABinaryNew("", "Should be C0100000000000000000000000000001", 16)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}