package main

const (
	// InvalidLengthError ...
	InvalidLengthError = 1001
	// NumberFormatError ...
	NumberFormatError = 1002
)

// IsoError ...
type IsoError struct {
	code    int
	message string
}

// NewIsoError ...
func NewIsoError(code int, message string) *IsoError {
	return &IsoError{code: code, message: message}
}

// Error ...
func (e *IsoError) Error() string {
	return e.message
}

// Errors ...
var Errors map[int]*IsoError

func init() {
	Errors = make(map[int]*IsoError)
	Errors[InvalidLengthError] = NewIsoError(InvalidLengthError, "Invalid length")
	Errors[NumberFormatError] = NewIsoError(NumberFormatError, "Invalid value, should be numeric")
}
