package main

import "strconv"

// LLLEChar ...
type LLLEChar struct {
	Codec
}

// LLLECharNew ...
func LLLECharNew(name string, description string, length int, padding bool) *LLLEChar {
	return &LLLEChar{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LLLEChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 999 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := AsciiToEbcdic(LeftPad2Len(strconv.Itoa(len(s)), "0", 3))
	return append(length, AsciiToEbcdic(s)...), nil
}

// Parse ...
func (codec *LLLEChar) Decode(b []byte) (string, error) {
	if len(b) < 4 {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToAscii(string(b))
	length, err := strconv.Atoi(string(b[:3]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+3 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[3 : length+3]), nil
}
