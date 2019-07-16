package main

import (
	"math/big"
	"strconv"
)

// ANumeric ...
type ANumeric struct {
	Fields Field
}

// Encode ...
func (codec *ANumeric) Encode(s string) ([]byte, error) {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, Errors[NumberFormatError]
	}
	return []byte(LeftPad2Len(s, "0", codec.Fields.Length)), nil
}

// Decode ...
func (codec *ANumeric) Decode(b []byte) (string, error) {
	if len(b) > codec.Fields.Length {
		return "", Errors[InvalidLengthError]
	}
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return "", Errors[NumberFormatError]
	}
	return strconv.Itoa(i), nil
}
