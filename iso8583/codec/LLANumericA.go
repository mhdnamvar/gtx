package codec

import (
	"../../iso8583"
	"fmt"
	"strconv"
)

type LLANumericA struct {
	Len  *NumericA
	Data *NumericA
}

func NewLLANumericA(id string, label string, padding Padding, paddingStr string, size int) *LLANumericA {
	if size > LLVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	return &LLANumericA{
		DefaultNumericA(LLVarA.Size),
		&NumericA{
			Id:          id,
			Label:       label,
			Encoding:    EncodingA,
			PaddingType: padding,
			PaddingStr:  paddingStr,
			MinLen:      0,
			MaxLen:      size,
		},
	}
}

func DefaultLLANumericA(size int) *LLANumericA {
	if size > LLVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	isoNumericA := DefaultNumericA(size)
	isoNumericA.MinLen = 0
	return &LLANumericA{
		DefaultNumericA(LLVarA.Size),
		isoNumericA,
	}
}

func (codec *LLANumericA) Encode(s string) ([]byte, error) {
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

func (codec *LLANumericA) Decode(b []byte) (string, int, error) {
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
	return string(data), len(data), nil
}

func (codec *LLANumericA) Check(s string) error {
	if codec.Data.PaddingType == NoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen || len(s) > LLVarA.MaxValue {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}
