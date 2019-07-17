package main

// Codable ...
type Codable interface {
	Encode(s string) ([]byte, error)
	Decode(b []byte) (string, error)
}
