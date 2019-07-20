package main

import (
	"strconv"
)

// LLLAChar ...
type LLLAChar struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LLLAChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 999 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := []byte(LeftPad2Len(strconv.Itoa(len(s)), "0", 3))
	return append(length, []byte(s)...), nil
}

// Decode ...
func (codec *LLLAChar) Decode(b []byte) (string, error) {
	if len(b) < 4 {
		return "", Errors[InvalidLengthError]
	}
	length, err := strconv.Atoi(string(b[:3]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+3 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[3 : length+3]), nil
}
