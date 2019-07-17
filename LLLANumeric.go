package main

import (
	"math/big"
	"strconv"
)

// LLLANumeric ...
type LLLANumeric struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LLLANumeric) Encode(s string) ([]byte, error) {
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

	return []byte(s), nil
}

// Decode ...
func (codec *LLLANumeric) Decode(b []byte) (string, error) {
	if len(b) > codec.Length {
		return "", Errors[InvalidLengthError]
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return "", Errors[NumberFormatError]
	}
	return strconv.Itoa(i), nil
}
