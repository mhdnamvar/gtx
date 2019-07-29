package main

// AChar ...
type AChar Codec

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
	return string(b), nil
}
