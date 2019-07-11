package main

import (
	"math/big"
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
