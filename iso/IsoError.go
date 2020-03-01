package main

import "fmt"

const (
	// TODO: should be change like OutOfBoundIndexError
	InvalidLengthError = 1001
	// NumberFormatError ...
	NumberFormatError = 1002
	// InvalidDataError ...
	InvalidDataError = 1003
)

var (
	OutOfBoundIndexError      = IsoErrorNew(1004, "Out of bound index error, pos should be between 0-128")
	IsoFieldNotFoundError     = IsoErrorNew(1005, "Iso field not found")
	NotSupportedBitmapError   = IsoErrorNew(1006, "Bitmap bigger than 2 bytes not supported")
	NotSupportedEncodingError = IsoErrorNew(1007, "Message encoding not supported, it should be ASCII, BINARY or EBCDIC")
	InvalidLengthTypeError    = IsoErrorNew(1008, "Fields length type is not valid, it should be FIXED, LLVAR or LLLVAR")
	NotSupported              = IsoErrorNew(1009, "Not supported")
)

// IsoError ...
type IsoError struct {
	code    int
	message string
}

// IsoErrorNew ...
func IsoErrorNew(code int, message string) *IsoError {
	return &IsoError{code: code, message: message}
}

// Error ...
func (e *IsoError) Error() string {
	return e.message
}

// String
func (e *IsoError) String() string {
	return fmt.Sprintf("%d %s", e.code, e.message)
}

// Errors ...
var Errors map[int]*IsoError

// TODO: should be change like OutOfBoundIndexError
func init() {
	Errors = make(map[int]*IsoError)
	Errors[InvalidLengthError] = IsoErrorNew(InvalidLengthError, "Invalid length")
	Errors[NumberFormatError] = IsoErrorNew(NumberFormatError, "Invalid value, should be numeric")
	Errors[InvalidDataError] = IsoErrorNew(InvalidDataError, "Invalid data")
}
