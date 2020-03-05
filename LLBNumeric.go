package main

import (
	"math/big"
	"strconv"
)

// LLBNumeric ...
type LLBNumeric struct {
	Codec
}

// LLBNumericNew ...
func LLBNumericNew(name string, description string, length int, padding bool) *LLBNumeric {
	return &LLBNumeric{Codec{name, description, length, padding}}
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
	return append(StrToBcd(LeftPad2Len(strconv.Itoa(len(s)), "0", 2)), StrToBcd(s)...), nil
}

// Parse ...
func (codec *LLBNumeric) Decode(b []byte) (string, error) {
	if len(b) < 2 {
		return "", Errors[InvalidLengthError]
	}
	length := BcdToInt(b[:1])
	if length%2 != 0 {
		length = length/2 + 1
	} else {
		length = length / 2
	}
	if length <= 0 || uint64(len(b)) < length+1 {
		return "", Errors[InvalidLengthError]
	}
	return strconv.FormatUint(BcdToInt(b[1:length+1]), 10), nil
}
