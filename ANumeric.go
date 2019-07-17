package main

import (
	"strconv"
)

// ANumeric ...
type ANumeric struct {
	Name        string
	Description string
	Length      int
}

// Encode ...
func (codec *ANumeric) Encode(s string) ([]byte, error) {
	_, err := strconv.Atoi(s)
	if err != nil {
		return nil, Errors[NumberFormatError]
	}
	return []byte(LeftPad2Len(s, "0", codec.Length)), nil
}

// Decode ...
func (codec *ANumeric) Decode(b []byte) (string, error) {
	if len(b) > codec.Length {
		return "", Errors[InvalidLengthError]
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return "", Errors[NumberFormatError]
	}
	return strconv.Itoa(i), nil
}
