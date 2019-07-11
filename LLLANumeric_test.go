package main

import (
	"reflect"
	"testing"
)

func Test_LLLANumeric_1(t *testing.T) {
	value := "12345"
	expected := []byte{
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x31, 0x32, 0x33, 0x34, 0x35}
	codec := LLLANumeric{Field{"",
		`Should be 3030303030303030303030303030303030303030303030303030303030303030` +
			`30303030303030303030303030303030303030303030303030303030303030303030303030` +
			`30303030303030303030303030303030303030303030303030303030303132333435`, 103}, true}
	actual, _ := codec.Encode(value)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}

}

func Test_LLLANumeric_2(t *testing.T) {
	value := "12345"
	expected := []byte{0x31, 0x32, 0x33, 0x34, 0x35}
	codec := LLLANumeric{Field{"", "Should be 3132333435", 103}, false}
	actual, _ := codec.Encode(value)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLLANumeric_3(t *testing.T) {
	value := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	codec := LLLANumeric{Field{"", "Should return LLL length error", 100}, false}
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

func Test_LLLANumeric_4(t *testing.T) {
	value := "12345ABC"
	codec := LLLANumeric{Field{"", "Should return nil, error", 199}, true}
	expected, err := codec.Encode(value)
	if !reflect.DeepEqual(err.Error(), Errors[NumberFormatError].message) {
		t.Errorf("Should return error\n")
	}
	if expected != nil {
		t.Errorf("Should return nil\n")
	}
}

func Test_LLLANumeric_5(t *testing.T) {
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
	codec := LLLANumeric{Field{"", "Should return error", 999}, false}
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
