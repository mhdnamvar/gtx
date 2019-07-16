package main

import (
	"math/big"
	"strconv"
)

// LLBNumeric ...
type LLBNumeric struct {
	Fields  Field
	Padding bool
}

// Encode ...
func (codec *LLBNumeric) Encode(s string) ([]byte, error) {
	if len(s) > codec.Fields.Length || len(s) > 99 {
		return nil, Errors[InvalidLengthError]
	}

	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, Errors[NumberFormatError]
	}

	if codec.Padding {
		s = LeftPad2Len(s, "0", codec.Fields.Length)
	}

	return StrToBcd(s), nil
}

// Decode ...
func (codec *LLBNumeric) Decode(b []byte) (string, error) {
	if len(b) > codec.Fields.Length || len(b) > 99 {
		return "", Errors[InvalidLengthError]
	}
	s := strconv.FormatUint(BcdToInt(b), 10)
	return LeftPad2Len(s, "0", codec.Fields.Length), nil
}
