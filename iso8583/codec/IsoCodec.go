package codec

type IsoCodec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, int, error)
	Pad(string) (string, error)
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
	IsoEncodingB IsoEncoding = 2
)
