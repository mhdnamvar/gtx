package main

import (
	"reflect"
	"testing"
)

func Test_ANumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{0x30, 0x30, 0x31, 0x032, 0x33, 0x34, 0x35}
	codec := ANumeric{Field{"", "Should be 30303132333435", 7}}
	actual, _ := codec.Encode(value)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}
