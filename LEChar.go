package main

import "strconv"

// LEChar ...
type LEChar struct {
	Codec
}

// LECharNew ...
func LECharNew(name string, description string, length int, padding bool) *LEChar {
	return &LEChar{Codec{name, description, length, padding}}
}

// Encode ...
func (codec *LEChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length || len(s) > 9 {
		return nil, Errors[InvalidLengthError]
	}
	if codec.Padding {
		s = RightPad2Len(s, " ", codec.Length)
	}
	length := AsciiToEbcdic(strconv.Itoa(len(s)))
	return append(length, AsciiToEbcdic(s)...), nil
}

// Parse ...
func (codec *LEChar) Decode(b []byte) (string, error) {
	if len(b) < 2 {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToAscii(string(b))
	length, err := strconv.Atoi(string(b[:1]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+1 {
		return "", Errors[InvalidLengthError]
	}
	return string(b[1 : length+1]), nil
}
