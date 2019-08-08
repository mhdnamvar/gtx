package main

// ABinary ...
type ABinary struct {
	Codec
}

// ABinaryNew ...
func ABinaryNew(name string, description string, length int) *ABinary {
	return &ABinary{Codec{name, description, length, true}}
}

// Encode ...
func (codec *ABinary) Encode(s string) ([]byte, error) {
	if len(s)%2 != 0 {
		return nil, Errors[InvalidLengthError]
	}
	s = LeftPad2Len(s, "0", codec.Length*2)
	return []byte(s), nil
}

// Decode ...
func (codec *ABinary) Decode(b []byte) (string, error) {
	if len(b) < codec.Length {
		return "", Errors[InvalidLengthError]
	}
	return string(b[:codec.Length]), nil
}
