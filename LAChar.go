package main

// LAChar ...
type LAChar struct {
	Name        string
	Description string
	Length      int
}

// Encode ...
func (codec *LAChar) Encode(s string) ([]byte, error) {
	return []byte(RightPad2Len(s, " ", codec.Length)), nil
}

// Decode ...
func (codec *LAChar) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	return string(b), nil
}
