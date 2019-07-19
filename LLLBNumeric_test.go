package main

import (
	"reflect"
	"testing"
)

func Test_LLLBNumeric_Encode(t *testing.T) {
	value := "12345"
	expected := []byte{0x01, 0x03,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x23, 0x45}
	codec := LLLBNumeric{"", "Should be 1030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012345",
		103, true}
	actual, err := codec.Encode(value)
	checkEncodeResult(t, expected, actual, err)
}

func Test_LLLBNumeric_EncodeNoPad(t *testing.T) {
	value := "12345"
	expected := []byte{0x05, 0x01, 0x23, 0x45}
	codec := LLLBNumeric{"", "Should be 012345", 103, false}
	actual, err := codec.Encode(value)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %x, expected: %x\n", actual, expected)
	}
}

func Test_LLLBNumeric_InvalidLen(t *testing.T) {
	value := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	codec := LLLBNumeric{"", "Should return LLL length error", 100, false}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLBNumeric_InvalidNumber(t *testing.T) {
	value := "12345ABC"
	codec := LLLBNumeric{"", "Should return nil, error", 199, true}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, NumberFormatError)
}

func Test_LLLBNumeric_InvalidLen2(t *testing.T) {
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
	codec := LLLBNumeric{"", "Should return error", 999, false}
	actual, err := codec.Encode(value)
	checkEncodeError(t, actual, err, InvalidLengthError)
}

func Test_LLLBNumeric_DecodePdded(t *testing.T) {
	value := []byte{0x01, 0x03,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x23, 0x45}
	expected := "12345"
	codec := LLLBNumeric{"", "Should be 12345", 103, true}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
func Test_LLLBNumeric_DecodeNoPad(t *testing.T) {
	value := []byte{0x00, 0x05, 0x01, 0x23, 0x45}
	expected := "12345"
	codec := LLLBNumeric{"", "Should be 12345", 103, true}
	actual, err := codec.Decode(value)
	checkDecodeResult(t, expected, actual, err)
}
