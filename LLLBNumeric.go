package main

import (
	"math/big"
	"strconv"
)

// LLLBNumeric ...
type LLLBNumeric struct {
	Codec
}

// LLLBNumericNew ...
func LLLBNumericNew(name string, description string, length int, padding bool) *LLLBNumeric {
	return &LLLBNumeric{Codec{name, description, length, padding}}
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
	return append(StrToBcd(LeftPad2Len(strconv.Itoa(len(s)), "0", 3)), StrToBcd(s)...), nil
}

// Decode ...
func (codec *LLLBNumeric) Decode(b []byte) (string, error) {
	if len(b) < 3 {
		return "", Errors[InvalidLengthError]
	}
	length := BcdToInt(b[:2])
	if length%2 != 0 {
		length = length/2 + 1
	} else {
		length = length / 2
	}
	if length <= 0 || uint64(len(b)) < length+2 {
		return "", Errors[InvalidLengthError]
	}
	return strconv.FormatUint(BcdToInt(b[2:length+2]), 10), nil
}
