package main

import (
	"encoding/hex"
	"fmt"
)
// ABinary ...
type ABinary struct {
	Name        string
	Description string
	Length      int
}

// Encode ...
func (codec *ABinary) Encode(s string) ([]byte, error) {
	if len(s)%2 != 0 || len(s)/2 > codec.Length {
		return nil, Errors[InvalidLengthError]
	}
	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, Errors[InvalidDataError]
	}
	return  data, nil
}

// Decode ...
func (codec *ABinary) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	return fmt.Sprintf("%X", b), nil
}
