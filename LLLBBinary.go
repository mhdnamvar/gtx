package main

import (
	"encoding/hex"
	"strconv"
	"fmt"
)
// LLLBBinary ...
type LLLBBinary struct {
	Name        string
	Description string
	Length      int
	Padding		bool
}

// Encode ...
func (codec *LLLBBinary) Encode(s string) ([]byte, error) {
	if len(s)%2 != 0 || len(s)/2 > codec.Length || len(s)/2 > 999 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = LeftPad2Len(s, "0", codec.Length*2)
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, Errors[InvalidDataError]
	}	
	return append(StrToBcd(LeftPad2Len(strconv.Itoa(len(s)/2), "0", 3)), b...), nil
}

// Decode ...
func (codec *LLLBBinary) Decode(b []byte) (string, error) {
	if len(b) < 2 {
		return "", Errors[InvalidLengthError]
	}
	length := BcdToInt(b[:2])
	if length%2 != 0 {
		length = length+1
	}
	if length <= 0 || uint64(len(b)) < length+2 {
		return "", Errors[InvalidLengthError]
	}
	return fmt.Sprintf("%X", b[2:length+2]), nil
}
