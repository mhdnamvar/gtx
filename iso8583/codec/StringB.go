package codec

import (
	"../../iso8583"
	"../../utils"
	"encoding/hex"
	"strings"
)

type StringB struct {
	Id          string
	Label       string
	Encoding    Encoding
	PaddingType Padding
	PaddingStr  string
	MinLen      int
	MaxLen      int
}

func NewStringB(id string, label string, padding Padding, paddingStr string, size int) *StringB {
	return &StringB{
		Id:          id,
		Label:       label,
		Encoding:    EncodingB,
		PaddingType: padding,
		PaddingStr:  paddingStr,
		MinLen:      size,
		MaxLen:      size,
	}
}

func DefaultStringB(size int) *StringB {
	return &StringB{
		Encoding:    EncodingB,
		PaddingType: NoPadding,
		PaddingStr:  "20",
		MinLen:      size,
		MaxLen:      size,
	}
}

func (codec *StringB) Encode(s string) ([]byte, error) {
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

func (codec *StringB) Pad(s string) (string, error) {
	if codec.PaddingType == LeftPadding {
		return utils.LeftPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	} else if codec.PaddingType == RightPadding {
		return utils.RightPad2Len(s, codec.PaddingStr, codec.MaxLen*2), nil
	}
	return s, nil
}

func (codec *StringB) Decode(b []byte) (string, int, error) {
	if len(b) < codec.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.MaxLen]
	return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
}

func (codec *StringB) Check(s string) error {
	if codec.PaddingType == NoPadding && (len(s) < codec.MinLen*2 || len(s) > codec.MaxLen*2) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.MaxLen*2 {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}
	return nil
}
