package codec

import (
	"../../iso8583"
	"fmt"
	"strconv"
)

type IsoLANumericA struct {
	Len  *IsoNumericA
	Data *IsoNumericA
}

func NewIsoLANumericA(id string, label string, padding IsoPadding, paddingStr string, size int) *IsoLANumericA {
	if size > LVar {
		panic(iso8583.InvalidLengthError)
	}
	return &IsoLANumericA{
		DefaultIsoNumericA(1),
		&IsoNumericA{
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

func DefaultIsoLANumericA(size int) *IsoLANumericA {
	if size > LVar {
		panic(iso8583.InvalidLengthError)
	}
	isoNumericA := DefaultIsoNumericA(size)
	isoNumericA.MinLen = 0
	return &IsoLANumericA{
		DefaultIsoNumericA(1),
		isoNumericA,
	}
}

func (codec *IsoLANumericA) Encode(s string) ([]byte, error) {
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

func (codec *IsoLANumericA) Decode(b []byte) (string, int, error) {
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

func (codec *IsoLANumericA) Check(s string) error {
	if codec.Data.PaddingType == IsoNoPadding &&
		(len(s) < codec.Data.MaxLen || len(s) > codec.Data.MaxLen) {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	if len(s) > codec.Data.MaxLen || len(s) > LVar {
		return iso8583.Errors[iso8583.InvalidLengthError]
	}

	return nil
}
