package main

import (
	"reflect"
	"testing"
)

func Test_BNumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	codec := BNumeric{Field{"", "Should be 000012345", 9}}
	actual, _ := codec.Encode(value)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_BNumeric_2(t *testing.T) {
	value := "12345ABC"
	codec := BNumeric{Field{"", "Should return nil, error", 9}}
	expected, err := codec.Encode(value)
	if !reflect.DeepEqual(err.Error(), Errors[NumberFormatError].message) {
		t.Errorf("Should return error\n")
	}
	if expected != nil {
		t.Errorf("Should return nil\n")
	}
}
