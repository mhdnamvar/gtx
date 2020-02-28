package main

import (
	"strconv"
	"fmt"
)

// LLLABinary ...
type LLLABinary struct {
	Codec
}

// LLLABinaryNew ...
func LLLABinaryNew(name string, description string, length int, padding bool) *LLLABinary {
	return &LLLABinary{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LLLABinary) Encode(s string) ([]byte, error) {	
	if len(s)%2 != 0 || len(s) > codec.Length {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := []byte(LeftPad2Len(strconv.Itoa(len(s)), "0", 3))
	return append(length, []byte(s)...), nil
}

// Decode ...
func (codec *LLLABinary) Decode(b []byte) (string, error) {
	fmt.Printf("----len(b)---- %d\n", len(b))
	if len(b) < 3 {
		return "", Errors[InvalidLengthError]
	}
	length, err := strconv.Atoi(string(b[:3]))
	fmt.Printf("---length--b[:3]--- %d %X\n", length, b[:3])
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+3 {		
		return "", Errors[InvalidLengthError]
	}	
	return string(b[3 : length+3]), nil
}
