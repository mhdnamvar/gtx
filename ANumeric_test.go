package main

import (
	"reflect"
	"testing"
)

func Test_ANumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{0x30, 0x30, 0x31, 0x032, 0x33, 0x34, 0x35}
	codec := ANumeric{Field{"", "Should be 30303132333435", 7}}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_ANumeric_2(t *testing.T) {
	value := []byte{0x30, 0x30, 0x31, 0x032, 0x33, 0x34, 0x35}
	expected := "12345"
	codec := ANumeric{Field{"", "Should be 12345", 7}}
	actual, err := codec.Decode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}
