package main

// IsoField ...
type IsoField struct {
	pos   int
	value string
}

// IsoFieldNew ...
func IsoFieldNew(pos int, value string) (*IsoField, error) {
	if pos < 0 || pos > 129 {
		return nil, OutOfBoundIndexError
	}
	return &IsoField{pos, value}, nil
}
