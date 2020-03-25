package codec

import (
	"../../iso8583"
	"../../utils"
	"encoding/hex"
	"math/big"
	"strings"
)

type NumericB struct {
	Id          string
	Label       string
	Encoding    Encoding
	PaddingType Padding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewNumericB(id string, label string, padding Padding, paddingStr string, size int) *NumericB {
	return &NumericB{
		Id:          id,
		Label:       label,
		Encoding:    EncodingB,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultNumericB(size int) *NumericB {
	return &NumericB{
		Encoding:    EncodingB,
		PaddingType: NoPadding,
		PaddingStr:  "00",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *NumericB) Encode(s string) ([]byte, error) {
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

func (codec *NumericB) Pad(s string) (string, error) {
	if codec.PaddingType == LeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	} else if codec.PaddingType == RightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	}
	return s, nil
}

func (codec *NumericB) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
}

func (codec *NumericB) Check(s string) error {
	if codec.PaddingType == NoPadding && (len(s) < codec.MinLen*2 || len(s) > codec.MaxLen*2) {
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
