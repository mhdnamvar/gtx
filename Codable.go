package main

// Codec ...
type Codec struct {
	Name        string
	Description string
	Length      int
	Padding     bool
}

// Codable ...
type Codable interface {
	GetName() string
	GetDescription() string
	GetLength() int
	GetPadding() bool
	Encode(s string) ([]byte, error)
	Decode(b []byte) (string, error)
}

// Protocol ...
type Protocol []Codable

// GetName ...
func (codec *Codec) GetName() string {
	return codec.Name
}

// GetDescription ...
func (codec *Codec) GetDescription() string {
	return codec.Description
}

// GetLength ...
func (codec *Codec) GetLength() int {
	return codec.Length
}

// GetPadding ...
func (codec *Codec) GetPadding() bool {
	return codec.Padding
}

// String ...
func (codec *Codec) String() string {
	return codec.Name
}