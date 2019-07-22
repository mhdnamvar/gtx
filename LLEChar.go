package main

import "strconv"

// LLEChar ...
type LLEChar struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LLEChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 99 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := ASCIIToEbcdic(LeftPad2Len(strconv.Itoa(len(s)), "0", 2))
	return append(length, ASCIIToEbcdic(s)...), nil
}

// Decode ...
func (codec *LLEChar) Decode(b []byte) (string, error) {
	if len(b) < 3 {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToASCII(string(b))
	length, err := strconv.Atoi(string(b[:2]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+2 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[2 : length+2]), nil
}
