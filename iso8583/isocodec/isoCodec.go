package isocodec

type IsoCodec interface {
	BeforeEncoding(string) error
	Encode(string) ([]byte, error)
	AfterEncoding([]byte) ([]byte, error)
	BeforeDecoding([]byte) error
	Decode([]byte) (string, int, error)
	AfterDecoding(string) (string, error)
	Pad(s string) (string, error)
	PadString() string
	Size() int
}

type IsoSpec []*IsoType
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
	IsoLeftPadF IsoPadding = 3
	IsoRightPadF IsoPadding = 4

	IsoString    IsoContentType = 0
	IsoNumeric   IsoContentType = 1
	IsoHexString IsoContentType = 2
	IsoAmount    IsoContentType = 3
	IsoBitmap    IsoContentType = 4
	IsoTrack2    IsoContentType = 5
	IsoTrack3    IsoContentType = 6
)

func (isoEncoding IsoEncoding) String() string {
	switch isoEncoding {
	case 0:
		return "IsoAscii"
	case 1:
		return "IsoEbcdic"
	case 2:
		return "IsoBinary"
	}
	return "IsoEncoding not defined"
}

func (isoPadding IsoPadding) String() string {
	switch isoPadding {
	case 0:
		return "IsoNoPad"
	case 1:
		return "IsoLeftPad"
	case 2:
		return "IsoRightPad"
	case 3:
		return "IsoLeftPadF"
	case 4:
		return "IsoRightPadF"
	}
	return "IsoPadding not defined"
}

func (isoContentType IsoContentType) String() string {
	switch isoContentType {
	case 0:
		return "IsoString"
	case 1:
		return "IsoNumeric"
	case 2:
		return "IsoHexString"
	case 3:
		return "IsoAmount"
	case 4:
		return "IsoBitmap"
	case 5:
		return "IsoTrack2"
	case 6:
		return "IsoTrack3"
	}
	//log.Printf("-------------%v\n", isoContentType)
	return "IsoContentType not defined"
}