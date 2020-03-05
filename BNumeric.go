package main

import (
	"math/big"
	"strconv"
)

// BNumeric ...
type BNumeric struct {
	Codec
}

// BNumericNew ...
func BNumericNew(name string, description string, length int) *BNumeric {
	return &BNumeric{Codec{name, description, length, true}}
}

// Encode ...
func (codec *BNumeric) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length {
		return nil, Errors[InvalidLengthError]
	}
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, Errors[NumberFormatError]
	}
	return StrToBcd(LeftPad2Len(s, "0", codec.Length)), nil
}

// Decode ...
func (codec *BNumeric) Decode(b []byte) (string, error) {
	if len(b) < codec.Length/2 {
		return "", Errors[InvalidLengthError]
	}
	return strconv.FormatUint(BcdToInt(b), 10), nil
}
