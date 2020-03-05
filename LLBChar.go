package main

import (
	"strconv"
)

// LLBChar ...
type LLBChar struct {
	Codec
}

// LLBCharNew ...
func LLBCharNew(name string, description string, length int, padding bool) *LLBChar {
	return &LLBChar{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LLBChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 99 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	return append(StrToBcd(LeftPad2Len(strconv.Itoa(len(s)), "0", 2)), []byte(s)...), nil
}

// Decode ...
func (codec *LLBChar) Decode(b []byte) (string, error) {
	if len(b) < 2 {
		return "", Errors[InvalidLengthError]
	}
	length := BcdToInt(b[:1])
	if length <= 0 || uint64(len(b)) < length+1 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[1 : length+1]), nil
}
