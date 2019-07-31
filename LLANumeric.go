package main

import (
	"math/big"
	"strconv"
)

// LLANumeric ...
type LLANumeric struct {
	Codec
}

// LLANumericNew ...
func LLANumericNew(name string, description string, length int, padding bool) *LLANumeric {
	return &LLANumeric{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LLANumeric) Encode(s string) ([]byte, error) {
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
	length := []byte(LeftPad2Len(strconv.Itoa(len(s)), "0", 2))
	return append(length, []byte(s)...), nil
}

// Decode ...
func (codec *LLANumeric) Decode(b []byte) (string, error) {
	if len(b) < 3 {
		return "", Errors[InvalidLengthError]
	}
	length, err := strconv.Atoi(string(b[:2]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+2 {
		return "", Errors[InvalidLengthError]
	}
	n := new(big.Int)
	n, ok := n.SetString(string(b[2:length+2]), 10)
	if !ok {
		return "", Errors[NumberFormatError]
	}
	s := n.String()
	if codec.Padding {
		s = LeftPad2Len(s, "0", codec.Length)
	}
	return s, nil
}
