package main

import (
	"testing"
)

func Test_AsciiEncodeNoPad(t *testing.T) {
	value := "0320"
	expected := []byte{0x30, 0x33, 0x32, 0x30}
	codec := &IsoText{ASCII, "MTI", "MESSAGE TYPE INDICATOR",
		&IsoLength{ASCII, FIXED, 4}, NONE}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
func Test_AsciiEncodeLeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte("   ABCD") // ]byte{0x20, 0x20, 0x20, 0x41, 0x42, 0x43, 0x44}
	codec := &IsoText{ASCII, "", "Should be '   ABCD'",
		&IsoLength{ASCII, FIXED, 7}, LEFT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_AsciiEncodeRightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte("ABCD   ") // ]byte{0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	codec := &IsoText{ASCII, "", "Should be 'ABCD   '",
		&IsoLength{ASCII, FIXED, 7}, RIGHT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BinaryEncodeNoPad(t *testing.T) {
	value := "0320"
	expected := []byte{0x03, 0x20}
	codec := &IsoText{BCD, "MTI", "MESSAGE TYPE INDICATOR",
		&IsoLength{BCD, FIXED, 4}, NONE}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BinaryEncodeLeftPad(t *testing.T) {
	value := "1234"
	expected := []byte{0x0, 0x12, 0x34}
	codec := &IsoText{BCD, "", "",
		&IsoLength{BCD, FIXED, 5}, LEFT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_BinaryEncodeRightPad(t *testing.T) {
	value := "1234"
	codec := &IsoText{BCD, "", "",
		&IsoLength{BCD, FIXED, 7}, RIGHT}
	actual, err := codec.Encode(value)
	assertEqual(t, InvalidPaddingError, err)
	assertEqual(t, nil, actual)
}

func Test_EbcdicEncodeNoPad(t *testing.T) {
	value := "0320"
	expected := []byte{0xF0, 0xF3, 0xF2, 0xF0}
	codec := &IsoText{EBCDIC, "MTI", "MESSAGE TYPE INDICATOR",
		&IsoLength{EBCDIC, FIXED, 4}, NONE}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_EbcdicEncodeLeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x40, 0xC1, 0xC2, 0xC3, 0xC4}
	codec := &IsoText{EBCDIC, "", "Should be ' ABCD'",
		&IsoLength{EBCDIC, FIXED, 5}, LEFT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_EbcdicEncodeRightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0xC1, 0xC2, 0xC3, 0xC4, 0x40}
	codec := &IsoText{EBCDIC, "", "Should be 'ABCD   '",
		&IsoLength{EBCDIC, FIXED, 5}, RIGHT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Encode_NoPad(t *testing.T) {
	value := "0320"
	expected := []byte(value)
	codec := &IsoText{ASCII, "MTI", "MESSAGE TYPE INDICATOR", &IsoLength{ASCII, FIXED, 4}, NONE}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Encode_LenErr(t *testing.T) {
	value := "032"
	codec := &IsoText{ASCII, "MTI", "MESSAGE TYPE INDICATOR",
		&IsoLength{ASCII, FIXED, 4}, NONE}
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_IsoText_Encode_LLVAR_NoPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x30, 0x34, 0x41, 0x42, 0x43, 0x44}
	codec := &IsoText{ASCII, "", "Should be '04ABCD'",
		&IsoLength{ASCII, LLVAR, 4}, NONE}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Encode_LLVAR_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x30, 0x37, 0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	codec := &IsoText{ASCII, "", "Should be '07ABCD   '",
		&IsoLength{ASCII, LLVAR, 7}, RIGHT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Encode_LLVAR_LeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{0x30, 0x37, 0x20, 0x20, 0x20, 0x41, 0x42, 0x43, 0x44}
	codec := &IsoText{ASCII, "", "Should be '07   ABCD'",
		&IsoLength{ASCII, LLVAR, 7}, LEFT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Encode_LLLVAR_NONE(t *testing.T) {
	value := "ABCD                                                                                                    "
	expected := []byte{
		0x31, 0x30, 0x34,
		0x41, 0x42, 0x43, 0x44,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	}
	codec := &IsoText{ASCII, "", "",
		&IsoLength{ASCII, LLLVAR, 104}, NONE}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Encode_LLLVAR_RightPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{
		0x31, 0x30, 0x34,
		0x41, 0x42, 0x43, 0x44,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	}
	codec := &IsoText{ASCII, "", "",
		&IsoLength{ASCII, LLLVAR, 104}, RIGHT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Encode_LLLVAR_LeftPad(t *testing.T) {
	value := "ABCD"
	expected := []byte{
		0x31, 0x30, 0x34,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x41, 0x42, 0x43, 0x44,
	}
	codec := &IsoText{ASCII, "", "",
		&IsoLength{ASCII, LLLVAR, 104}, LEFT}
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

/*
func Test_IsoText_EncodePad(t *testing.T) {
	value := "ABCD"
	expected := []byte("11ABCD       ")
	codec := &IsoTextNew{"", "Should be '11ABCD       '", 11, true)
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_InvalidLen(t *testing.T) {
	value := "ABCDEFGHIJKL"
	codec := &IsoTextNew{"", "Should return error", 11, false)
	actual, err := codec.Encode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_IsoText_Decode(t *testing.T) {
	value := []byte("04ABCD       ")
	expected := "ABCD"
	codec := &IsoTextNew{"", "Should be 'ABCD'", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_DecodeInvalidLen(t *testing.T) {
	value := []byte("10ABCD")
	codec := &IsoTextNew{"", "Should return error", 10, false)
	actual, err := codec.Decode(value)
	assertEqual(t, Errors[InvalidLengthError], err)
	assertEqual(t, nil, actual)
}

func Test_IsoText_DecodePad(t *testing.T) {
	value := []byte("10ABCD          ")
	expected := "ABCD      "
	codec := &IsoTextNew{"", "Should be 'ABCD      '", 10, true)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}

func Test_IsoText_Decode(t *testing.T) {
	value := []byte{0x41, 0x42, 0x43, 0x44, 0x20, 0x20, 0x20}
	expected := "ABCD   "
	codec := IsoText("", "Should be 'ABCD   '", 7)
	actual, err := codec.Decode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
*/
