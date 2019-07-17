package main

import (
	"reflect"
	"testing"
)

func Test_LLBNumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{0x00, 0x00, 0x00, 0x01, 0x23, 0x45}
	codec := LLBNumeric{"", "Should be 00000012345", 11, true}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLBNumeric_2(t *testing.T) {
	value := "12345"
	expected := []byte{0x01, 0x23, 0x45}
	codec := LLBNumeric{"", "Should be 012345", 11, false}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLBNumeric_3(t *testing.T) {
	value := "12345"
	codec := LLBNumeric{"", "Should return error", 4, false}
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

func Test_LLBNumeric_4(t *testing.T) {
	value := "12345ABC"
	codec := LLBNumeric{"", "Should return nil, error", 9, true}
	actual, err := codec.Encode(value)
	if err == nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(err.Error(), Errors[NumberFormatError].message) {
		t.Errorf("Should return error\n")
	}
	if actual != nil {
		t.Errorf("Should return nil\n")
	}
}

func Test_LLBNumeric_5(t *testing.T) {
	value := []byte{0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x23, 0x45}
	expected := "00000012345"
	codec := LLBNumeric{"", "Should be 00000012345", 11, true}
	actual, err := codec.Decode(value)
	if err != nil {
		t.Errorf("Should return value\n")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}
