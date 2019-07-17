package main

import (
	"reflect"
	"testing"
)

func Test_LLANumeric_encode(t *testing.T) {
	value := "12345"
	expected := []byte("0512345") //[]byte{0x30, 0x35, 0x31, 0x32, 0x33, 0x34, 0x35}
	codec := LLANumeric{"", "Should be 0512345", 11}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLANumeric_invalidLen(t *testing.T) {
	value := "123456789012"
	codec := LLANumeric{"", "Should return error", 11}
	actual, err := codec.Encode(value)
	if err == nil {
		t.Errorf("Should return error\n")
	}
	if !reflect.DeepEqual(err.Error(), Errors[InvalidLengthError].message) {
		t.Errorf("Should return error\n")
	}
	if actual != nil {
		t.Errorf("Should return nil\n")
	}
}

func Test_LLANumeric_InvalidNumber(t *testing.T) {
	value := "1234567890A"
	codec := LLANumeric{"", "Should return nil, error", 11}
	actual, err := codec.Encode(value)
	if err == nil {
		t.Errorf("Should return error\n")
	}
	if !reflect.DeepEqual(err.Error(), Errors[NumberFormatError].message) {
		t.Errorf("Should return error\n")
	}
	if actual != nil {
		t.Errorf("Should return nil\n")
	}
}
