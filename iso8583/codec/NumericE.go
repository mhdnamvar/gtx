package codec

import (
	"../../iso8583"
	"../../utils"
	"math/big"
)

type NumericE struct {
	Id          string
	Label       string
	Encoding    Encoding
	PaddingType Padding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewNumericE(id string, label string, padding Padding, paddingStr string, size int) *NumericE {
	return &NumericE{
		Id:          id,
		Label:       label,
		Encoding:    EncodingE,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultNumericE(size int) *NumericE {
	return &NumericE{
		Encoding:    EncodingE,
		PaddingType: NoPadding,
		PaddingStr:  "0",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *NumericE) Encode(s string) ([]byte, error) {
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	err = codec.Check(s)
	if err != nil {
		return nil, err
	}

	return utils.AsciiToEbcdic(s), nil
}

func (codec *NumericE) Pad(s string) (string, error) {
	if codec.PaddingType == LeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	} else if codec.PaddingType == RightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	}
	return s, nil
}

func (codec *NumericE) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return string(utils.EbcdicToAsciiBytes(data)), len(data), nil
}

func (codec *NumericE) Check(s string) error {
	if codec.PaddingType == NoPadding && (len(s) < codec.MinLen || len(s) > codec.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.MaxLen {
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
