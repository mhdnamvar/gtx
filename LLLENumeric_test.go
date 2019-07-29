package main

import (
	"testing"
)

func Test_LLLENumeric_Encode(t *testing.T) {
	value := "12345"
	expected := []byte{
		0xf1, 0xf0, 0xf3,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf1, 0xf2, 0xf3, 0xf4, 0xf5}
	codec := LLLENumeric{"", `Should be 103000000000000000000000000000000000000000` +
		`0000000000000000000000000000000000000000000000000000000000012345`, 103, true}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)

	expected = ASCIIToEbcdic("00512345")
	codec = LLLENumeric{"", `Should be 00512345`, 103, false}
	actual, err = codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_LLLENumeric_EncodeInvalidLength(t *testing.T) {
	value := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	codec := LLLENumeric{"", "Should return LLL length error", 100, false}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLLENumeric_EncodeInvalidFormat(t *testing.T) {
	value := "12345ABC"
	codec := LLLENumeric{"", "Should return nil, error", 199, false}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[NumberFormatError], err)
	assertEqual(t, nil, actual)
}

func Test_LLLENumeric_EncodeInvalidLength2(t *testing.T) {
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
	codec := LLLENumeric{"", "Should return error", 999, false}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_LLLENumeric_Decode(t *testing.T) {
	value := []byte{0xf1, 0xf0, 0xf3,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0, 0xf0,
		0xf1, 0xf2, 0xf3, 0xf4, 0xf5}
	expected := `00000000000000000000000000000000000000000000000000` +
		`00000000000000000000000000000000000000000000000012345`
	codec := LLLENumeric{"", `Should be 000000000000000000000000000000000000000` +
		`0000000000000000000000000000000000000000000000000000000000012345`, 103, true}
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
