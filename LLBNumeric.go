package main

import (
	"math/big"
	"strconv"
)

// LLBNumeric ...
type LLBNumeric struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LLBNumeric) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 99 {
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
	// length := IntToBcd(uint64(len(s)))
	// b := []byte(s)
	// b = append(length, b...)

	return StrToBcd(s), nil
}

// Decode ...
func (codec *LLBNumeric) Decode(b []byte) (string, error) {
	if len(b) > codec.Length || len(b) > 99 {
		return "", Errors[InvalidLengthError]
	}
	s := strconv.FormatUint(BcdToInt(b), 10)
	return LeftPad2Len(s, "0", codec.Length), nil
}
