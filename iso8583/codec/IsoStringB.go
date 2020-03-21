package codec

import (
	"../../iso8583"
	"../../utils"
	"encoding/hex"
	"strings"
)

type IsoStringB struct {
	Id          string
	Label       string
	Encoding    IsoEncoding
	PaddingType IsoPadding
	PaddingStr  string
	Size        int
}

func NewIsoStringB(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoStringB {
	return &IsoStringB{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingB,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		Size:        size,
	}
}

func DefaultIsoStringB(size int) *IsoStringB {
	return &IsoStringB{
		Id:          "",
		Label:       "",
		Encoding:    IsoEncodingB,
		PaddingType: IsoNoPadding,
		PaddingStr:  "20",
		Size:        size,
	}
}

func (codec *IsoStringB) Encode(s string) ([]byte, error) {
	// Do padding if required
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	// Check length
	if codec.PaddingType == IsoNoPadding && len(s) != codec.Size*2 {
		return nil, iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.Size*2 {
		return nil, iso8583.Errors[iso8583.InvalidLengthError]
	}

	bytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, iso8583.Errors[iso8583.InvalidDataError]
	}
	return bytes, nil
}

func (codec *IsoStringB) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.Size*2), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.Size*2), nil
	}
	return s, nil
}

func (codec *IsoStringB) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Size {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.Size]
	return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
}
