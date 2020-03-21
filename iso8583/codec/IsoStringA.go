package codec

import (
	"../../iso8583"
	"../../utils"
)

type IsoStringA struct {
	Id          string
	Label       string
	Encoding    IsoEncoding
	PaddingType IsoPadding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewIsoStringA(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoStringA {
	return &IsoStringA{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingA,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultIsoStringA(size int) *IsoStringA {
	return &IsoStringA{
		Encoding:    IsoEncodingA,
		PaddingType: IsoNoPadding,
		PaddingStr:  " ",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *IsoStringA) Encode(s string) ([]byte, error) {
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	err = codec.Check(s)
	if err != nil {
		return nil, err
	}

	return []byte(s), nil
}

func (codec *IsoStringA) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	}
	return s, nil
}

func (codec *IsoStringA) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return string(data), len(data), nil
}

func (codec *IsoStringA) Check(s string) error {
	if codec.PaddingType == IsoNoPadding && (len(s) < codec.MinLen || len(s) > codec.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.MaxLen {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	return nil
}
