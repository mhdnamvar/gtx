package main

import (
	"encoding/hex"
	"fmt"
)
// BBinary ...
type BBinary struct {
	Name        string
	Description string
	Length      int
}

// Encode ...
func (codec *BBinary) Encode(s string) ([]byte, error) {
	if len(s)%2 != 0 || len(s)/2 != codec.Length {
		return nil, Errors[InvalidLengthError]
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, Errors[InvalidDataError]
	}
	return  b, nil
}

// Decode ...
func (codec *BBinary) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	return fmt.Sprintf("%X", b[:codec.Length]), nil
}
