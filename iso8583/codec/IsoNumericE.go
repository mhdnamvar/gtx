package codec

import (
	"../../iso8583"
	"../../utils"
	"math/big"
)

type IsoNumericE struct {
	Id          string
	Label       string
	Encoding    IsoEncoding
	PaddingType IsoPadding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewIsoNumericE(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoNumericE {
	return &IsoNumericE{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingE,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultIsoNumericE(size int) *IsoNumericE {
	return &IsoNumericE{
		Encoding:    IsoEncodingE,
		PaddingType: IsoNoPadding,
		PaddingStr:  "0",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *IsoNumericE) Encode(s string) ([]byte, error) {
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

func (codec *IsoNumericE) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	}
	return s, nil
}

func (codec *IsoNumericE) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return string(utils.EbcdicToAsciiBytes(data)), len(data), nil
}

func (codec *IsoNumericE) Check(s string) error {
	if codec.PaddingType == IsoNoPadding && (len(s) < codec.MinLen || len(s) > codec.MaxLen) {
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
