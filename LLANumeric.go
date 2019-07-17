package main

import (
	"math/big"
	"strconv"
)

// LLANumeric ...
type LLANumeric struct {
	Name        string
	Description string
	Length      int
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
	length := []byte(LeftPad2Len(strconv.Itoa(len(s)), "0", 2))
	return append(length, []byte(s)...), nil
}

// Decode ...
func (codec *LLANumeric) Decode(b []byte) (string, error) {
	// if len(b) > codec.Length || len(b) > 99 {
	// 	return "", Errors[InvalidLengthError]
	// }
	// b[:1]

	// i, err := strconv.Atoi(string(b))
	// if err != nil {
	// 	return "", Errors[NumberFormatError]
	// }
	// return strconv.Itoa(i), nil
	return "", nil
}
