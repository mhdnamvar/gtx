package codec

import (
	"../../iso8583"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type LBStringB struct {
	Len  *NumericB
	Data *StringB
}

func NewLBStringB(id string, label string, padding Padding, paddingStr string, size int) *LBStringB {
	if size > LVarB.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	return &LBStringB{
		DefaultNumericB(LVarB.Size),
		&StringB{
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

func DefaultLBStringB(size int) *LBStringB {
	if size > LVarB.MaxValue {
		panic(iso8583.InvalidLengthError)
	}
	data := DefaultStringB(size)
	data.MinLen = 0
	return &LBStringB{
		DefaultNumericB(LVarB.Size),
		data,
	}
}

func (codec *LBStringB) Encode(s string) ([]byte, error) {
	err := codec.Check(s)
	if err != nil {
		fmt.Println("==> 5")
		return nil, err
	}

	b, err := codec.Data.Encode(s)
	if err != nil {
		fmt.Println("==> 6")
		return nil, err
	}

	l, err := codec.Len.Encode(fmt.Sprintf("%02d", len(b)))
	if err != nil {
		fmt.Println("==> 7")
		return nil, err
	}

	return append(l, b...), nil
}

func (codec *LBStringB) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Len.MaxLen+codec.Data.MaxLen {
		fmt.Println("==> 1")
		return "", 0, iso8583.NotEnoughData
	}

	bytes := b[:codec.Len.MaxLen]
	s, n, err := codec.Len.Decode(bytes)
	if err != nil {
		fmt.Println("==> 2")
		return "", 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("==> 3")
		return "", 0, iso8583.Errors[iso8583.InvalidDataError]
	}

	if n+i > len(b) {
		fmt.Println("==> 4")
		return "", 0, iso8583.NotEnoughData
	}

	data := b[n : n+i]
	return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
}

func (codec *LBStringB) Check(s string) error {
	if codec.Data.PaddingType == NoPadding &&
		(len(s) < codec.Data.MaxLen*2 || len(s) > codec.Data.MaxLen*2) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen*2 {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}
