package main

import "strconv"

// LEChar ...
type LEChar struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Encode ...
func (codec *LEChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 9 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := ASCIIToEbcdic(strconv.Itoa(len(s)))
	return append(length, ASCIIToEbcdic(s)...), nil
}

// Decode ...
func (codec *LEChar) Decode(b []byte) (string, error) {
	if len(b) < 2 {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToASCII(string(b))
	length, err := strconv.Atoi(string(b[:1]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+1 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[1 : length+1]), nil
}
