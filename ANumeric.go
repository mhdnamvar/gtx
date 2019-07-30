package main

import (
	"math/big"
	"strconv"
)

// ANumeric ...
type ANumeric struct {
	Codec
}

// ANumericNew ...
func ANumericNew(name string, description string, length int) *ANumeric {
	return &ANumeric{Codec{name, description, length, true}}
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
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	n := new(big.Int)
	n, ok := n.SetString(string(b), 10)
	if !ok {
		return "", Errors[NumberFormatError]
	}
	return n.String(), nil
}
