package codec

import (
	"../../iso8583"
	"../../utils"
	"fmt"
)

type IsoLAStringA struct {
	Len  *IsoNumericA
	Data *IsoStringA
}

func NewIsoLAStringA(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoLAStringA {
	return &IsoLAStringA{
		DefaultIsoNumericA(1),
		&IsoStringA{
			Id:          id,
			Label:       label,
			Encoding:    IsoEncodingA,
			PaddingType: padding,
			PaddingStr:  paddingStr,
			MinLen:      0,
			MaxLen:      size,
		},
	}
}

func DefaultIsoLAStringA(size int) *IsoLAStringA {
	isoStringA := DefaultIsoStringA(size)
	isoStringA.MinLen = 0
	return &IsoLAStringA{
		DefaultIsoNumericA(1),
		isoStringA,
	}
}

func (codec *IsoLAStringA) Encode(s string) ([]byte, error) {
	err := codec.Check(s)
	if err != nil {
		fmt.Println("==> 1")
		return nil, err
	}

	l, err := codec.Len.Encode(fmt.Sprintf("%d", len(s)))
	if err != nil {
		fmt.Println("==> 2")
		return nil, err
	}

	b, err := codec.Data.Encode(s)
	if err != nil {
		fmt.Println("==> 3")
		return nil, err
	}
	fmt.Printf("%x %x", l, b)
	return append(l, b...), nil
}

func (codec *IsoLAStringA) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Len.MaxLen+codec.Data.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}

	bytes := b[:codec.Len.MaxLen]
	i, err := utils.Btoi(bytes)
	if err != nil {
		return "", 0, err
	}

	data := b[:i]
	return string(data), len(data), nil
}

func (codec *IsoLAStringA) Check(s string) error {
	if codec.Data.PaddingType == IsoNoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen || len(s) > 9 {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}
