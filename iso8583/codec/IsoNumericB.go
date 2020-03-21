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
	Size        int
}

func NewIsoNumericB(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoNumericB {
	return &IsoNumericB{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingB,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		Size:        size,
	}
}

func DefaultIsoNumericB(size int) *IsoNumericB {
	return &IsoNumericB{
		Id:          "",
		Label:       "",
		Encoding:    IsoEncodingB,
		PaddingType: IsoNoPadding,
		PaddingStr:  "00",
		Size:        size,
	}
}

func (codec *IsoNumericB) Encode(s string) ([]byte, error) {
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

	// Check numeric
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, iso8583.Errors[iso8583.NumberFormatError]
	}

	return utils.StrToBcd(s), nil
}

func (codec *IsoNumericB) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.Size*2), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.Size*2), nil
	}
	return s, nil
}

func (codec *IsoNumericB) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Size {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.Size]
	return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
}
