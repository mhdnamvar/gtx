package main

import (
	"reflect"
	"testing"
)

func Test_LLANumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35}
	codec := LLANumeric{Field{"", "Should be 3030303030303132333435", 11}, true}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLANumeric_2(t *testing.T) {
	value := "12345"
	expected := []byte{0x31, 0x32, 0x33, 0x34, 0x35}
	codec := LLANumeric{Field{"", "Should be 3132333435", 11}, false}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLANumeric_3(t *testing.T) {
	value := "123456789012"
	codec := LLANumeric{Field{"", "Should return error", 11}, false}
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

func Test_LLANumeric_4(t *testing.T) {
	value := "1234567890A"
	codec := LLANumeric{Field{"", "Should return nil, error", 11}, true}
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
