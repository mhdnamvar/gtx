package main

import "strconv"

// LAChar ...
type LAChar struct {
	Codec
}

// LACharNew ...
func LACharNew(name string, description string, length int, padding bool) *LAChar {
	return &LAChar{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LAChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 9 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := []byte(strconv.Itoa(len(s)))
	return append(length, []byte(s)...), nil
}

// Decode ...
func (codec *LAChar) Decode(b []byte) (string, error) {
	if len(b) < 2 {
		return "", Errors[InvalidLengthError]
	}
	length, err := strconv.Atoi(string(b[:1]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+1 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[1 : length+1]), nil
}
