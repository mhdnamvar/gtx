package main

// AChar ...
type AChar struct {
	Codec
}

// ACharNew ...
func ACharNew(name string, description string, length int) *AChar {
	return &AChar{Codec{name, description, length, true}}
}

// Encode ...
func (codec *AChar) Encode(s string) ([]byte, error) {
	if len(s) > codec.Length {
		return nil, Errors[InvalidLengthError]
	}
	return []byte(RightPad2Len(s, " ", codec.Length)), nil
}

// Decode ...
func (codec *AChar) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	return string(b[:codec.Length]), nil
}
