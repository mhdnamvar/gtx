package codec

import (
	"../../iso8583"
	"../../utils"
	"math/big"
)

type IsoNumericA struct {
	Len  *IsoCodec
	Data *IsoCodec
}

func (codec *IsoNumericA) New(id string, label string, size int) *IsoNumericA {
	return &IsoNumericA{
		nil,
		&IsoCodec{
			Id:          id,
			Label:       label,
			Encoding:    IsoEncodingA,
			PaddingType: IsoLeftPadding,
			PaddingStr:  "0",
			MinLen:      0,
			MaxLen:      size,
		},
	}
}

func (codec *IsoNumericA) Encode(s string) ([]byte, error) {
	// Do padding if required
	s, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	// Check length
	if codec.Data.PaddingType == IsoNoPadding && len(s) != codec.Data.MaxLen {
		return nil, iso8583.Errors[iso8583.InvalidLengthError]
	}
	if len(s) > codec.Data.MaxLen {
		return nil, iso8583.Errors[iso8583.InvalidLengthError]
	}

	// Check numeric
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, iso8583.Errors[iso8583.NumberFormatError]
	}

	return []byte(s), nil
}

func (codec *IsoNumericA) Pad(s string) (string, error) {
	if codec.Data.PaddingType == IsoLeftPadding {
		return utils.LeftPad2Len(s, codec.Data.PaddingStr, codec.Data.MaxLen), nil
	} else if codec.Data.PaddingType == IsoRightPadding {
		return utils.RightPad2Len(s, codec.Data.PaddingStr, codec.Data.MaxLen), nil
	}
	return s, nil
}

func (codec *IsoNumericA) Decode(b []byte) (string, int, error) {
	if len(b) < codec.Data.MaxLen {
		return "", 0, iso8583.NotEnoughData
	}
	data := b[:codec.Data.MaxLen]
	return string(data), len(data), nil
}
