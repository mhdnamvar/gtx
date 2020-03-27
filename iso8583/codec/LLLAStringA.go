package codec

import (
	"../../iso8583"
	"fmt"
	"strconv"
)

type LLLAStringA struct {
	Len  *NumericA
	Data *StringA
}

func NewIsoLLLAStringA(id string, label string, padding Padding, paddingStr string, size int) *LLLAStringA {
	if size > LLLVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	return &LLLAStringA{
		DefaultNumericA(LLLVarA.Size),
		&StringA{
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

func DefaultLLLAStringA(size int) *LLLAStringA {
	if size > LLLVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	isoStringA := DefaultStringA(size)
	isoStringA.MinLen = 0
	return &LLLAStringA{
		DefaultNumericA(LLLVarA.Size),
		isoStringA,
	}
}

func (codec *LLLAStringA) Encode(s string) ([]byte, error) {
	err := codec.Check(s)
	if err != nil {
		fmt.Println("==> 1")
		return nil, err
	}

	b, err := codec.Data.Encode(s)
	if err != nil {
		fmt.Println("==> 3")
		return nil, err
	}

	l, err := codec.Len.Encode(fmt.Sprintf("%03d", len(b)))
	if err != nil {
		fmt.Println("==> 2")
		return nil, err
	}

	fmt.Printf("%x %x", l, b)
	return append(l, b...), nil
}

func (codec *LLLAStringA) Decode(b []byte) (string, int, error) {
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

func (codec *LLLAStringA) Check(s string) error {
	if codec.Data.PaddingType == NoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen || len(s) > LLLVarA.MaxValue {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}
