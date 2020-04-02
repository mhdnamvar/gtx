package codec

import "fmt"

type IsoError struct {
	code    int
	message string
}

var (
	OutOfBoundIndex    = &IsoError{1001, "Out of bound index error, pos should be between 0-128"}
	FieldNotFound      = &IsoError{1002, "Iso field not found"}
	NotSupportedBitmap = &IsoError{1003, "Bitmap bigger than 2 bytes not supported"}
	InvalidLength      = &IsoError{1004, "Fields length type is not valid, it should be Fixed, LLVar or LLLVar"}
	InvalidData        = &IsoError{1005, "Invalid data"}
	NotEnoughData      = &IsoError{1006, "Not enough data"}
	NotSupported       = &IsoError{1007, "Not supported"}
)

func (e *IsoError) Error() string {
	return e.message
}

func (e *IsoError) String() string {
	return fmt.Sprintf("Iso Error %d: %s", e.code, e.message)
}
