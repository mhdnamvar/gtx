package codec

import (
	"../../iso8583"
	"fmt"
	"strconv"
)

type LAStringA struct {
	Len  *NumericA
	Data *StringA
}

func NewIsoLAStringA(id string, label string, padding Padding, paddingStr string, size int) *LAStringA {
	if size > LVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	return &LAStringA{
		DefaultNumericA(LVarA.Size),
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

func DefaultLAStringA(size int) *LAStringA {
	if size > LVarA.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	isoStringA := DefaultStringA(size)
	isoStringA.MinLen = 0
	return &LAStringA{
		DefaultNumericA(LVarA.Size),
		isoStringA,
	}
}

func (codec *LAStringA) Encode(s string) ([]byte, error) {
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

func (codec *LAStringA) Decode(b []byte) (string, int, error) {
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

func (codec *LAStringA) Check(s string) error {
	if codec.Data.PaddingType == NoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen || len(s) > LVarA.MaxValue {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}
