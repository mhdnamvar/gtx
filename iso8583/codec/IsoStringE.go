package codec

import (
	"../../iso8583"
	"../../utils"
)

type IsoStringE struct {
	Id          string
	Label       string
	Encoding    IsoEncoding
	PaddingType IsoPadding
	PaddingStr  string
	Size        int
}

func NewIsoStringE(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoStringE {
	return &IsoStringE{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingE,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		Size:        size,
	}
}

func DefaultIsoStringE(size int) *IsoStringE {
	return &IsoStringE{
		Id:          "",
		Label:       "",
		Encoding:    IsoEncodingE,
		PaddingType: IsoNoPadding,
		PaddingStr:  " ",
		Size:        size,
	}
}

func (codec *IsoStringE) Encode(s string) ([]byte, error) {
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
	return utils.AsciiToEbcdic(s), nil
}

func (codec *IsoStringE) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.Size), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.Size), nil
	}
	return s, nil
}

func (codec *IsoStringE) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Size {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.Size]
	return string(utils.EbcdicToAsciiBytes(data)), len(data), nil
}
