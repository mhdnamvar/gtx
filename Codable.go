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
	Encode(s string) ([]byte, error)
	Decode(b []byte) (string, error)
}

// Protocol ...
type Protocol [129]Codable
