package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

type IsoNumeric struct{ IsoText }

func IsoNumericNew(encoding IsoEncoding, name string, desc string, length *IsoLength, padding IsoPadding) *IsoNumeric {
	return &IsoNumeric{IsoText{encoding, name, desc, length, padding}}
}

func (isoNumeric *IsoNumeric) Pad(s string) (string, error) {
	l := isoNumeric.Length.Value
	padStr := "0"
	if isoNumeric.Encoding == BINARY {
		l = isoNumeric.Length.Value * 2
		padStr = "00"
	}
	if isoNumeric.Padding == LEFT {
		return LeftPad2Len(s, padStr, l), nil
	} else if isoNumeric.Padding == RIGHT {
		return s, NotSupported
	}
	return s, nil
}

func (isoNumeric *IsoNumeric) Encode(s string) ([]byte, error) {
	fmt.Printf("Input: [%s]\n", s)
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		fmt.Printf("Error: Input is not  numeric: [%s]\n", s)
		return nil, Errors[NumberFormatError]
	}

	data := s
	if isoNumeric.Padding != NONE {
		padding, err := isoNumeric.Pad(s)
		if err != nil {
			return nil, err
		}
		data = padding
		fmt.Printf("Padding: [%s]\n", data)
	}

	dataLen, err := isoNumeric.DataLen(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("LenBytes: [%X]\n", dataLen)

	switch isoNumeric.Encoding {
	case ASCII:
		fmt.Printf("ASCII encoded: [%X]\n", append(dataLen, []byte(data)...))
		return append(dataLen, []byte(data)...), nil
	case BINARY:
		b, err := hex.DecodeString(data)
		if err != nil {
			fmt.Printf("Error BINARY: %s\n", Errors[InvalidDataError])
			return nil, Errors[InvalidDataError]
		}
		fmt.Printf("Binary encoded: [%X]\n", append(dataLen, b...))
		return append(dataLen, b...), nil
	case EBCDIC:
		fmt.Printf("EBCDIC encoded: [%X]\n", append(dataLen, AsciiToEbcdic(data)...))
		return append(dataLen, AsciiToEbcdic(data)...), nil
	default:
		return nil, NotSupportedEncodingError
	}
}
