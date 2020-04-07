package isocodec

type IsoCodec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, int, error)
	BeforeEncoding(string) error
	BeforeDecoding([]byte) error
	Pad(s string) (string, error)
	PadString() string
	Size() int
}

type IsoSpec [][]*IsoCodec
type IsoEncoding int
type IsoPadding int
type IsoContentType int
type IsoPadString string

const (
	IsoAscii  IsoEncoding = 0
	IsoEbcdic IsoEncoding = 1
	IsoBinary IsoEncoding = 2

	IsoNoPad    IsoPadding = 0
	IsoLeftPad  IsoPadding = 1
	IsoRightPad IsoPadding = 2

	IsoString    IsoContentType = 0
	IsoNumeric   IsoContentType = 1
	IsoHexString IsoContentType = 3
	IsoAmount    IsoContentType = 4
	IsoBitmap    IsoContentType = 5
	IsoTrack2    IsoContentType = 6
	IsoTrack3    IsoContentType = 7
)

var (
	StringA = &IsoType{
		nil,
		&IsoData{IsoAscii, 0, 4, IsoString, IsoNoPad},
	}

	LLAStringA = &IsoType{
		&IsoData{IsoAscii, 0, 0, IsoString, IsoNoPad},
		&IsoData{IsoAscii, 0, 0, IsoString, IsoNoPad},
	}
)
