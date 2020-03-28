package codec

import (
	"../../iso8583"
	"../../utils"
	"fmt"
	"strconv"
)

type LENumericE struct {
	Len  *NumericE
	Data *NumericE
}

func NewLENumericE(id string, label string, padding Padding, paddingStr string, size int) *LENumericE {
	if size > LVarE.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	return &LENumericE{
		DefaultNumericE(LVarE.Size),
		&NumericE{
			Id:          id,
			Label:       label,
			Encoding:    EncodingE,
			PaddingType: padding,
			PaddingStr:  paddingStr,
			MinLen:      0,
			MaxLen:      size,
		},
	}
}

func DefaultLENumericE(size int) *LENumericE {
	if size > LVarE.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	data := DefaultNumericE(size)
	data.MinLen = 0
	return &LENumericE{
		DefaultNumericE(LVarE.Size),
		data,
	}
}

func (codec *LENumericE) Encode(s string) ([]byte, error) {
	err := codec.Check(s)
	if err != nil {
		return nil, err
	}

	b, err := codec.Data.Encode(s)
	if err != nil {
		return nil, err
	}

	l, err := codec.Len.Encode(fmt.Sprintf("%d", len(b)))
	if err != nil {
		return nil, err
	}

	return append(l, b...), nil
}

func (codec *LENumericE) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Len.MaxLen+codec.Data.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}

	bytes := b[:codec.Len.MaxLen]
	s, n, err := codec.Len.Decode(bytes)
	if err != nil {
		return "", 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return "", 0, iso8583.Errors[iso8583.InvalidDataError]
	}

	if n+i > len(b) {
		return "", 0, iso8583.NotEnoughData
	}

	data := b[n : n+i]
	return string(utils.EbcdicToAsciiBytes(data)), len(data), nil
}

func (codec *LENumericE) Check(s string) error {
	if codec.Data.PaddingType == NoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}
