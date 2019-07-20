package main

import (
	"math/big"
	"strconv"
)

// ENumeric ...
type ENumeric struct {
	Name        string
	Description string
	Length      int
}

// Encode ...
func (codec *ENumeric) Encode(s string) ([]byte, error) {
	_, err := strconv.Atoi(s)
	if err != nil {
		return nil, Errors[NumberFormatError]
	}
	return ASCIIToEbcdic(LeftPad2Len(s, "0", codec.Length)), nil
}

// Decode ...
func (codec *ENumeric) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToASCII(string(b))
	n := new(big.Int)
	n, ok := n.SetString(string(b), 10)
	if !ok {
		return "", Errors[NumberFormatError]
	}
	return n.String(), nil
}
