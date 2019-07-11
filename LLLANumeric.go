package main

import "math/big"

// LLLANumeric ...
type LLLANumeric struct {
	Fields  Field
	Padding bool
}

// Encode ...
func (codec *LLLANumeric) Encode(s string) ([]byte, error) {
	if len(s) > codec.Fields.Length || len(s) > 999 {
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

	return []byte(s), nil
}
