package codec

import (
	"../../iso8583"
	"../../utils"
)

type StringE struct {
	Id          string
	Label       string
	Encoding    Encoding
	PaddingType Padding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewStringE(id string, label string, padding Padding, paddingStr string, size int) *StringE {
	return &StringE{
		Id:          id,
		Label:       label,
		Encoding:    EncodingE,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultIsoStringE(size int) *StringE {
	return &StringE{
		Encoding:    EncodingE,
		PaddingType: NoPadding,
		PaddingStr:  " ",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *StringE) Encode(s string) ([]byte, error) {
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

func (codec *StringE) Pad(s string) (string, error) {
	if codec.PaddingType == LeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	} else if codec.PaddingType == RightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	}
	return s, nil
}

func (codec *StringE) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return string(utils.EbcdicToAsciiBytes(data)), len(data), nil
}

func (codec *StringE) Check(s string) error {
	if codec.PaddingType == NoPadding && (len(s) < codec.MinLen || len(s) > codec.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.MaxLen {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	return nil
}
