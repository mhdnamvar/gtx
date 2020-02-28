package main

import "strconv"

// LLEChar ...
type LLEChar struct {
	Codec
}

// LLECharNew ...
func LLECharNew(name string, description string, length int, padding bool) *LLEChar {
	return &LLEChar{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LLEChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 99 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := AsciiToEbcdic(LeftPad2Len(strconv.Itoa(len(s)), "0", 2))
	return append(length, AsciiToEbcdic(s)...), nil
}

// Decode ...
func (codec *LLEChar) Decode(b []byte) (string, error) {
	if len(b) < 3 {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToAscii(string(b))
	length, err := strconv.Atoi(string(b[:2]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+2 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[2 : length+2]), nil
}
