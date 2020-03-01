package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

type IsoText struct {
	Encoding  IsoEncoding
	Name      string
	Desc      string
	Length    *IsoLength
	Padding   IsoPadding
	IsNumeric bool
}

type IsoLength struct {
	Encoding IsoEncoding
	Type     IsoLenType
	Value    int
}

type IsoEncoding int
type IsoPadding int
type IsoLenType int

const (
	FIXED  IsoLenType = 0
	LLVAR  IsoLenType = 1
	LLLVAR IsoLenType = 2

	ASCII  IsoEncoding = 0
	BINARY IsoEncoding = 1
	EBCDIC IsoEncoding = 2

	NONE  IsoPadding = 0
	LEFT  IsoPadding = 1
	RIGHT IsoPadding = 2
)

func IsoLengthNew(encoding IsoEncoding, lenType IsoLenType, value int) *IsoLength {
	return &IsoLength{encoding, lenType, value}
}

func IsoTextNew(encoding IsoEncoding, name string, desc string, length *IsoLength, padding IsoPadding) *IsoText {
	return &IsoText{encoding, name, desc, length, padding, false}
}

func IsoNumericNew(encoding IsoEncoding, name string, desc string, length *IsoLength, padding IsoPadding) *IsoText {
	return &IsoText{encoding, name, desc, length, padding, true}
}

func (isoText *IsoText) DataLen(s string) ([]byte, error) {
	switch isoText.Length.Encoding {
	case ASCII:
		return isoText.AsciiLen(s)
	case EBCDIC:
		return isoText.EbcdicLen(s)
	case BINARY:
		return isoText.BinaryLen(s)
	default:
		return nil, InvalidLengthTypeError
	}
}

func (isoText *IsoText) AsciiLen(s string) ([]byte, error) {
	l := len(s)
	switch isoText.Length.Type {
	case FIXED:
		if isoText.Padding == NONE {
			if l != isoText.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if l > isoText.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		}
		return []byte{}, nil
	case LLVAR:
		if l == 0 || l > isoText.Length.Value || l > 99 {
			fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return []byte(LeftPad2Len(strconv.Itoa(l), "0", 2)), nil
	case LLLVAR:
		if l == 0 || l > isoText.Length.Value || l > 999 {
			fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return []byte(LeftPad2Len(strconv.Itoa(l), "0", 3)), nil
	default:
		fmt.Printf("Error: %s\n", InvalidLengthTypeError)
		return nil, InvalidLengthTypeError
	}
}

func (isoText *IsoText) EbcdicLen(s string) ([]byte, error) {
	l := len(s)
	switch isoText.Length.Type {
	case FIXED:
		if isoText.Padding == NONE {
			if l != isoText.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if l > isoText.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		}
		return []byte{}, nil
	case LLVAR:
		if l == 0 || l > isoText.Length.Value || l > 99 {
			fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return AsciiToEbcdic(LeftPad2Len(strconv.Itoa(l), "0", 2)), nil
	case LLLVAR:
		if l == 0 || l > isoText.Length.Value || l > 999 {
			fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return AsciiToEbcdic(LeftPad2Len(strconv.Itoa(l), "0", 3)), nil
	default:
		fmt.Printf("Error: %s\n", InvalidLengthTypeError)
		return nil, InvalidLengthTypeError
	}
}

func (isoText *IsoText) BinaryLen(s string) ([]byte, error) {
	l := len(s) / 2
	switch isoText.Length.Type {
	case FIXED:
		if isoText.Padding == NONE {
			if l != isoText.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if l > isoText.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		}
		return nil, nil
	case LLVAR:
		if l == 0 || l > isoText.Length.Value || l > 99 {
			fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return StrToBcd(LeftPad2Len(strconv.Itoa(l), "0", 2)), nil
	case LLLVAR:
		if l == 0 || l > isoText.Length.Value || l > 999 {
			fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return StrToBcd(LeftPad2Len(strconv.Itoa(l), "0", 3)), nil
	default:
		return nil, Errors[InvalidLengthError]
	}
}

func (isoText *IsoText) Pad(s string) (string, error) {
	l := isoText.Length.Value
	padStr := " "
	if isoText.IsNumeric {
		padStr = "0"
	}

	if isoText.Encoding == BINARY {
		l = isoText.Length.Value * 2
		padStr = "20"
		if isoText.IsNumeric {
			padStr = "00"
		}
	}

	if isoText.Padding == LEFT {
		return LeftPad2Len(s, padStr, l), nil
	} else if isoText.Padding == RIGHT {
		if isoText.IsNumeric {
			return s, NotSupported
		}
		return RightPad2Len(s, padStr, l), nil
	}
	return s, nil
}

func (isoText *IsoText) Encode(s string) ([]byte, error) {
	fmt.Printf("Input: [%s]\n", s)
	if isoText.IsNumeric {
		n := new(big.Int)
		n, ok := n.SetString(s, 10)
		if !ok {
			fmt.Printf("Error: Input is not  numeric: [%s]\n", s)
			return nil, Errors[NumberFormatError]
		}
	}

	data := s
	if isoText.Padding != NONE {
		padding, err := isoText.Pad(s)
		if err != nil {
			return nil, err
		}
		data = padding
		fmt.Printf("Padding: [%s]\n", data)
	}

	dataLen, err := isoText.DataLen(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("LenBytes: [%X]\n", dataLen)

	switch isoText.Encoding {
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

func (isoText *IsoText) Decode(b []byte) (string, error) {
	if len(b) < isoText.Length.Value {
		return "", Errors[InvalidLengthError]
	}
	return string(b[:isoText.Length.Value]), nil
}
