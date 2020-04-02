package codec

type IsoCodec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, int, error)
	Check(string) error
	LenSize() int
}

type VarLen struct {
	Size     int
	MaxValue int
}

type IsoPadding int
type IsoContentType int
type IsoEncoding int

const (
	NoPadding    IsoPadding = 0
	LeftPadding  IsoPadding = 1
	RightPadding IsoPadding = 2

	EncodingA IsoEncoding = 0
	EncodingE IsoEncoding = 1
	EncodingB IsoEncoding = 2
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
		MaxValue: 999,
	}
	LVarE VarLen = VarLen{
		Size:     1,
		MaxValue: 9,
	}
	LLVarE VarLen = VarLen{
		Size:     2,
		MaxValue: 99,
	}
	LLLVarE VarLen = VarLen{
		Size:     3,
		MaxValue: 999,
	}
	LVarB VarLen = VarLen{
		Size:     1,
		MaxValue: 9,
	}
	LLVarB VarLen = VarLen{
		Size:     1,
		MaxValue: 99,
	}
	LLLVarB VarLen = VarLen{
		Size:     2,
		MaxValue: 999,
	}
)
