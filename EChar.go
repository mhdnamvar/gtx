package main

// EChar ...
type EChar struct {
	Codec
}

// ECharNew ...
func ECharNew(name string, description string, length int) *EChar {
	return &EChar{Codec{name, description, length, true}}
}

// Encode ...
func (codec *EChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length {
		return nil, Errors[InvalidLengthError]
	}
	return AsciiToEbcdic(RightPad2Len(s, " ", codec.Length)), nil
}

// Decode ...
func (codec *EChar) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	a := EbcdicToAscii(string(b))
	return string(a), nil
}
