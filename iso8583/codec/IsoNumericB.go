package codec

import (
	"../../iso8583"
	"../../utils"
	"encoding/hex"
	"math/big"
	"strings"
)

type IsoNumericB struct {
	Id          string
	Label       string
	Encoding    IsoEncoding
	PaddingType IsoPadding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewIsoNumericB(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoNumericB {
	return &IsoNumericB{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingB,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultIsoNumericB(size int) *IsoNumericB {
	return &IsoNumericB{
		Encoding:    IsoEncodingB,
		PaddingType: IsoNoPadding,
		PaddingStr:  "00",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *IsoNumericB) Encode(s string) ([]byte, error) {
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	err = codec.Check(s)
	if err != nil {
		return nil, err
	}

	return utils.StrToBcd(s), nil
}

func (codec *IsoNumericB) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	}
	return s, nil
}

func (codec *IsoNumericB) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
}

func (codec *IsoNumericB) Check(s string) error {
	if codec.PaddingType == IsoNoPadding && (len(s) < codec.MinLen*2 || len(s) > codec.MaxLen*2) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.MaxLen*2 {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	// Check numeric
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return iso8583.Errors[iso8583.NumberFormatError]
	}
	return nil
}
