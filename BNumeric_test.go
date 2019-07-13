package main

import (
	"reflect"
	"testing"
)

func Test_BNumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	codec := BNumeric{Field{"", "Should be 000012345", 9}}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_BNumeric_2(t *testing.T) {
	value := "12345ABC"
	codec := BNumeric{Field{"", "Should return nil, error", 9}}
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
func Test_BNumeric_3(t *testing.T) {
	value := "12345"
	codec := BNumeric{Field{"", "Should return invalid length error", 4}}
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
func Test_BNumeric_4(t *testing.T) {
	value := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	expected := "12345"
	codec := BNumeric{Field{"", "Should be 12345", 9}}
	actual, err := codec.Decode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_BNumeric_5(t *testing.T) {
	value := []byte{0x00, 0x00, 0x01, 0x23, 0x45}
	expected := ""
	codec := BNumeric{Field{"", "Should return error", 4}}
	actual, err := codec.Decode(value)
	if err == nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(err.Error(), Errors[InvalidLengthError].message) {
		t.Errorf("Should return error\n")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}
