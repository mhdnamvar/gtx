package codec

import (
	"fmt"
	"strconv"

	"../../iso8583"
)

type LLAStringA struct {
	Len  *NumericA
	Data *StringA
}

func NewLLAStringA(id string, label string, padding IsoPadding, paddingStr string, size int) *LLAStringA {
	if size > LLVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	return &LLAStringA{
		DefaultNumericA(LLVarA.Size),
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

func DefaultLLAStringA(size int) *LLAStringA {
	if size > LLVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	isoStringA := DefaultStringA(size)
	isoStringA.MinLen = 0
	return &LLAStringA{
		DefaultNumericA(LLVarA.Size),
		isoStringA,
	}
}

func (codec *LLAStringA) Encode(s string) ([]byte, error) {
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

	l, err := codec.Len.Encode(fmt.Sprintf("%02d", len(b)))
	if err != nil {
		fmt.Println("==> 2")
		return nil, err
	}

	fmt.Printf("%x %x", l, b)
	return append(l, b...), nil
}

func (codec *LLAStringA) Decode(b []byte) (string, int, error) {
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

func (codec *LLAStringA) Check(s string) error {
	if codec.Data.PaddingType == NoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen || len(s) > LLVarA.MaxValue {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}

func (codec *LLAStringA) LenSize() int {
	return LLVarA.Size
}
