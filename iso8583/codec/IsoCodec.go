package codec

type Codec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, int, error)
	Check(string) error
}

type VarLen struct {
	Size     int
	MaxValue int
}

type Padding int
type ContentType int
type Encoding int

const (
	NoPadding    Padding = 0
	LeftPadding  Padding = 1
	RightPadding Padding = 2

	EncodingA Encoding = 0
	EncodingE Encoding = 1
	EncodingB Encoding = 2
)

var (
	LVarA VarLen = VarLen{
		Size:     1,
		MaxValue: 9,
	}
	LLVarA VarLen = VarLen{
		Size:     2,
		MaxValue: 99,
	}
	LLLVarA VarLen = VarLen{
		Size:     3,
		MaxValue: 99,
	}
)
