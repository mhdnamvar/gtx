package codec

type Codec interface {
	New() *interface{}
	Encode(string) ([]byte, error)
	Decode([]byte) (string, int, error)
	Pad(string) (string, error)
}

type IsoCodec struct {
	Id          string
	Label       string
	Encoding    IsoEncoding
	PaddingType IsoPadding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

type IsoPadding int
type IsoContentType int
type IsoEncoding int

const (
	IsoNoPadding    IsoPadding = 0
	IsoLeftPadding  IsoPadding = 1
	IsoRightPadding IsoPadding = 2

	IsoEncodingA IsoEncoding = 0
	IsoEncodingE IsoEncoding = 1
	IsoEncodingB IsoEncoding = 3

	//IsoAlpha               IsoContentType = 0
	//IsoNumericA             IsoContentType = 1
	//IsoAmount              IsoContentType = 2
	//IsoSpecial             IsoContentType = 3
	//IsoAlphaNumeric        IsoContentType = 4
	//IsoAlphaSpecial        IsoContentType = 5
	//IsoNumericSpecial      IsoContentType = 6
	//IsoAlphaNumericSpecial IsoContentType = 7
	//IsoBinary              IsoContentType = 8
	//IsoTrack2              IsoContentType = 9
	//IsoTrack3              IsoContentType = 10
)
