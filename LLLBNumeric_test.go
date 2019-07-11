package main

import (
	"reflect"
	"testing"
)

func Test_LLLBNumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x23, 0x45}
	codec := LLLBNumeric{Field{"",
		"Should be 0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012345",
		103}, true}
	actual, _ := codec.Encode(value)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}

}

func Test_LLLBNumeric_2(t *testing.T) {
	value := "12345"
	expected := []byte{0x01, 0x23, 0x45}
	codec := LLLBNumeric{Field{"", "Should be 012345", 103}, false}
	actual, _ := codec.Encode(value)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLLBNumeric_3(t *testing.T) {
	value := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	codec := LLLBNumeric{Field{"", "Should return LLL length error", 100}, false}
	expected, err := codec.Encode(value)

	if err == nil {
		t.Errorf("Should return error\n")
	}

	if !reflect.DeepEqual(err.Error(), Errors[InvalidLengthError].message) {
		t.Errorf("Should return invalid length error\n")
	}
	if expected != nil {
		t.Errorf("Should return nil\n")
	}
}

func Test_LLLBNumeric_4(t *testing.T) {
	value := "12345ABC"
	codec := LLLBNumeric{Field{"", "Should return nil, error", 199}, true}
	expected, err := codec.Encode(value)
	if !reflect.DeepEqual(err.Error(), Errors[NumberFormatError].message) {
		t.Errorf("Should return error\n")
	}
	if expected != nil {
		t.Errorf("Should return nil\n")
	}
}

func Test_LLLBNumeric_5(t *testing.T) {
	value := `1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890` +
		`1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890`
	codec := LLLBNumeric{Field{"", "Should return error", 999}, false}
	expected, err := codec.Encode(value)

	if err == nil {
		t.Errorf("Should return error\n")
	}

	if !reflect.DeepEqual(err.Error(), Errors[InvalidLengthError].message) {
		t.Errorf("Should return error\n")
	}
	if expected != nil {
		t.Errorf("Should return nil\n")
	}
}
