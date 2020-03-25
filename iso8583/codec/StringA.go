package codec

import (
	"../../iso8583"
	"../../utils"
)

type StringA struct {
	Id          string
	Label       string
	Encoding    Encoding
	PaddingType Padding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewStringA(id string, label string, padding Padding, paddingStr string, size int) *StringA {
	return &StringA{
		Id:          id,
		Label:       label,
		Encoding:    EncodingA,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultStringA(size int) *StringA {
	return &StringA{
		Encoding:    EncodingA,
		PaddingType: NoPadding,
		PaddingStr:  " ",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *StringA) Encode(s string) ([]byte, error) {
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	err = codec.Check(s)
	if err != nil {
		return nil, err
	}

	return []byte(s), nil
}

func (codec *StringA) Pad(s string) (string, error) {
	if codec.PaddingType == LeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	} else if codec.PaddingType == RightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen), nil
	}
	return s, nil
}

func (codec *StringA) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return string(data), len(data), nil
}

func (codec *StringA) Check(s string) error {
	if codec.PaddingType == NoPadding && (len(s) < codec.MinLen || len(s) > codec.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.MaxLen {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	return nil
}
