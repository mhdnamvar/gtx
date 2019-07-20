package main

import (
	"math/big"
	"strconv"
)

// LLLENumeric ...
type LLLENumeric struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LLLENumeric) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 999 {
		return nil, Errors[InvalidLengthError]
	}
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, Errors[NumberFormatError]
	}
	if codec.Padding {
		s = LeftPad2Len(s, "0", codec.Length)
	}
	length := ASCIIToEbcdic(LeftPad2Len(strconv.Itoa(len(s)), "0", 3))
	return append(length, ASCIIToEbcdic(s)...), nil
}

// Decode ...
func (codec *LLLENumeric) Decode(b []byte) (string, error) {
	if len(b) < 4 {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToASCII(string(b))
	length, err := strconv.Atoi(string(b[:3]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+3 {
		return "", Errors[InvalidLengthError]
	}
	n := new(big.Int)
	n, ok := n.SetString(string(b[3:length+3]), 10)
	if !ok {
		return "", Errors[NumberFormatError]
	}
	s := n.String()
	if codec.Padding {
		s = LeftPad2Len(s, "0", codec.Length)
	}
	return s, nil
}
