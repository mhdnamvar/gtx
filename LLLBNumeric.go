package main

import (
	"math/big"
	"strconv"
)

// LLLBNumeric ...
type LLLBNumeric struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LLLBNumeric) Encode(s string) ([]byte, error) {
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

	return StrToBcd(s), nil
}

// Decode ...
func (codec *LLLBNumeric) Decode(b []byte) (string, error) {
	if len(b) > codec.Length || len(b) > 999 {
		return "", Errors[InvalidLengthError]
	}
	s := strconv.FormatUint(BcdToInt(b), 10)
	return LeftPad2Len(s, "0", codec.Length), nil
}
