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
	Size        int
}

func NewIsoNumericE(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoNumericE {
	return &IsoNumericE{
		Id:          id,
		Label:       label,
		Encoding:    IsoEncodingE,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		Size:        size,
	}
}

func DefaultIsoNumericE(size int) *IsoNumericE {
	return &IsoNumericE{
		Id:          "",
		Label:       "",
		Encoding:    IsoEncodingE,
		PaddingType: IsoNoPadding,
		PaddingStr:  "0",
		Size:        size,
	}
}

func (codec *IsoNumericE) Encode(s string) ([]byte, error) {
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

	// Check numeric
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, iso8583.Errors[iso8583.NumberFormatError]
	}

	return utils.AsciiToEbcdic(s), nil
}

func (codec *IsoNumericE) Pad(s string) (string, error) {
	if codec.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.Size), nil
	} else if codec.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.Size), nil
	}
	return s, nil
}

func (codec *IsoNumericE) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Size {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.Size]
	return string(utils.EbcdicToAsciiBytes(data)), len(data), nil
}
