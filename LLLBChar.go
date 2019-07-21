package main

import (
	"strconv"
)

// LLLBChar ...
type LLLBChar struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LLLBChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 999 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	return append(StrToBcd(LeftPad2Len(strconv.Itoa(len(s)), "0", 3)), []byte(s)...), nil
}

// Decode ...
func (codec *LLLBChar) Decode(b []byte) (string, error) {
	if len(b) < 3 {
		return "", Errors[InvalidLengthError]
	}
	length := BcdToInt(b[:2])
	if length <= 0 || uint64(len(b)) < length+2 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[2 : length+2]), nil
}
