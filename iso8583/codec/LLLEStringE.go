package codec

import (
	"../../iso8583"
	"../../utils"
	"fmt"
	"strconv"
)

type LLLEStringE struct {
	Len  *NumericE
	Data *StringE
}

func NewLLLEStringE(id string, label string, padding IsoPadding, paddingStr string, size int) *LLLEStringE {
	if size > LLLVarE.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	return &LLLEStringE{
		DefaultNumericE(LLLVarE.Size),
		&StringE{
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

func DefaultLLLEStringE(size int) *LLLEStringE {
	if size > LLLVarE.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	data := DefaultStringE(size)
	data.MinLen = 0
	return &LLLEStringE{
		DefaultNumericE(LLLVarE.Size),
		data,
	}
}

func (codec *LLLEStringE) Encode(s string) ([]byte, error) {
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

	l, err := codec.Len.Encode(fmt.Sprintf("%d", len(b)))
	if err != nil {
		fmt.Println("==> 2")
		return nil, err
	}

	fmt.Printf("%x %x", l, b)
	return append(l, b...), nil
}

func (codec *LLLEStringE) Decode(b []byte) (string, int, error) {
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

func (codec *LLLEStringE) Check(s string) error {
	if codec.Data.PaddingType == NoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen || len(s) > LLLVarE.MaxValue {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}

func (codec *LLLEStringE) LenSize() int {
	return LLLVarE.Size
}
