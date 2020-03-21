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
	MinLen      int
	MaxLen      int
}

func NewIsoStringB(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoStringB {
	return &IsoStringB{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingB,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultIsoStringB(size int) *IsoStringB {
	return &IsoStringB{
		Encoding:    IsoEncodingB,
		PaddingType: IsoNoPadding,
		PaddingStr:  "20",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *IsoStringB) Encode(s string) ([]byte, error) {
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	err = codec.Check(s)
	if err != nil {
		return nil, err
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, iso8583.Errors[iso8583.InvalidDataError]
	}

	return b, nil
}

func (codec *IsoStringB) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	}
	return s, nil
}

func (codec *IsoStringB) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
}

func (codec *IsoStringB) Check(s string) error {
	if codec.PaddingType == IsoNoPadding && (len(s) < codec.MinLen*2 || len(s) > codec.MaxLen*2) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.MaxLen*2 {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	return nil
}
