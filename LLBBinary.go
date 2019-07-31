package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// LLBBinary ...
type LLBBinary struct {
	Codec
}

// LLBBinaryNew ...
func LLBBinaryNew(name string, description string, length int, padding bool) *LLBBinary {
	return &LLBBinary{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LLBBinary) Encode(s string) ([]byte, error) {
	if len(s)%2 != 0 || len(s)/2 > codec.Length || len(s)/2 > 99 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = LeftPad2Len(s, "0", codec.Length*2)
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, Errors[InvalidDataError]
	}
	return append(StrToBcd(LeftPad2Len(strconv.Itoa(len(s)/2), "0", 2)), b...), nil
}

// Decode ...
func (codec *LLBBinary) Decode(b []byte) (string, error) {
	if len(b) < 2 {
		return "", Errors[InvalidLengthError]
	}
	length := BcdToInt(b[:1])
	if length%2 != 0 {
		length = length + 1
	}
	if length <= 0 || uint64(len(b)) < length+1 {
		return "", Errors[InvalidLengthError]
	}
	return fmt.Sprintf("%X", b[1:length+1]), nil
}
