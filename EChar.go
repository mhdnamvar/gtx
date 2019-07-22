package main

// EChar ...
type EChar struct {
	Name        string
	Description string
	Length      int
}

// Encode ...
func (codec *EChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length {
		return nil, Errors[InvalidLengthError]
	}
	return ASCIIToEbcdic(RightPad2Len(s, " ", codec.Length)), nil
}

// Decode ...
func (codec *EChar) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	a := EbcdicToASCII(string(b))
	return string(a), nil
}
