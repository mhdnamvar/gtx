package main

import (
	"strconv"
)

// LLAChar ...
type LLAChar struct {
	Codec
}

// LLACharNew ...
func LLACharNew(name string, description string, length int, padding bool) *LLAChar {
	return &LLAChar{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LLAChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 99 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := []byte(LeftPad2Len(strconv.Itoa(len(s)), "0", 2))
	return append(length, []byte(s)...), nil
}

// Parse ...
func (codec *LLAChar) Decode(b []byte) (string, error) {
	if len(b) < 3 {
		return "", Errors[InvalidLengthError]
	}
	length, err := strconv.Atoi(string(b[:2]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+2 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[2 : length+2]), nil
}
