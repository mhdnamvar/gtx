package main

// Codable ...
type Codable interface {
	Encode(s string) ([]byte, error)
}

// Field ...
type Field struct {
	Name        string
	Description string
	Length      int
}
