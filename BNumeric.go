package main

import (
	"math/big"
)

// BNumeric ...
type BNumeric struct {
	Fields Field
}

// Encode ...
func (codec *BNumeric) Encode(s string) ([]byte, error) {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, Errors[NumberFormatError]
	}
	return StrToBcd(LeftPad2Len(s, "0", codec.Fields.Length)), nil
}
