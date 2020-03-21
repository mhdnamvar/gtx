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
	Size        int
}

func NewIsoStringA(id string, label string, padding IsoPadding,
	paddingStr string, size int) *IsoStringA {
	return &IsoStringA{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingA,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		Size:        size,
	}
}

func DefaultIsoStringA(size int) *IsoStringA {
	return &IsoStringA{
		Id:          "",
		Label:       "",
		Encoding:    IsoEncodingA,
		PaddingType: IsoNoPadding,
		PaddingStr:  " ",
		Size:        size,
	}
}

func (codec *IsoStringA) Encode(s string) ([]byte, error) {
	// Do padding if required
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	// Check length
	if codec.PaddingType == IsoNoPadding && len(s) != codec.Size {
		return nil, iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.Size {
		return nil, iso8583.Errors[iso8583.InvalidLengthError]
	}

	return []byte(s), nil
}

func (codec *IsoStringA) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.Size), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.Size), nil
	}
	return s, nil
}

func (codec *IsoStringA) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Size {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.Size]
	return string(data), len(data), nil
}
