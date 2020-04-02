package codec

import (
	"../../iso8583"
	"../../utils"
	"fmt"
	"math/big"
)

type NumericA struct {
	Id          string
	Label       string
	Encoding    IsoEncoding
	PaddingType IsoPadding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewNumericA(id string, label string, padding IsoPadding, paddingStr string, size int) *NumericA {
	return &NumericA{
		Id:          id,
		Label:       label,
		Encoding:    EncodingA,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultNumericA(size int) *NumericA {
	return &NumericA{
		Encoding:    EncodingA,
		PaddingType: NoPadding,
		PaddingStr:  "0",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *NumericA) Encode(s string) ([]byte, error) {
	s, err := codec.Pad(s)
	if err != nil {
		fmt.Println("==> 4")
		return nil, err
	}

	err = codec.Check(s)
	if err != nil {
		fmt.Println("==> 4")
		return nil, err
	}

	return []byte(s), nil
}

func (codec *NumericA) Pad(s string) (string, error) {
	if codec.PaddingType == LeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	} else if codec.PaddingType == RightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	}
	return s, nil
}

func (codec *NumericA) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return string(data), len(data), nil
}

func (codec *NumericA) Check(s string) error {
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

func (codec *NumericA) LenSize() int {
	return 0
}
