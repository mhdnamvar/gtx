package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

type IsoCodec struct {
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

func IsoLenCodec(encoding IsoEncoding, lenType IsoLenType, value int) *IsoLength {
	return &IsoLength{encoding, lenType, value}
}

func IsoTextCodec(encoding IsoEncoding, name string, desc string, length *IsoLength, padding IsoPadding) *IsoCodec {
	return &IsoCodec{encoding, name, desc, length, padding, false}
}

func IsoNumericCodec(encoding IsoEncoding, name string, desc string, length *IsoLength, padding IsoPadding) *IsoCodec {
	return &IsoCodec{encoding, name, desc, length, padding, true}
}

func (codec *IsoCodec) CodecLen(s string) ([]byte, error) {
	switch codec.Length.Encoding {
	case ASCII:
		return codec.AsciiLen(s)
	case EBCDIC:
		return codec.EbcdicLen(s)
	case BINARY:
		return codec.BinaryLen(s)
	default:
		return nil, InvalidLengthTypeError
	}
}

func (codec *IsoCodec) AsciiLen(s string) ([]byte, error) {
	l := len(s)
	switch codec.Length.Type {
	case FIXED:
		if codec.Padding == NONE {
			if l != codec.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if l > codec.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		}
		return []byte{}, nil
	case LLVAR:
		if l == 0 || l > codec.Length.Value || l > 99 {
			fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return []byte(LeftPad2Len(strconv.Itoa(l), "0", 2)), nil
	case LLLVAR:
		if l == 0 || l > codec.Length.Value || l > 999 {
			fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return []byte(LeftPad2Len(strconv.Itoa(l), "0", 3)), nil
	default:
		fmt.Printf("Error: %s\n", InvalidLengthTypeError)
		return nil, InvalidLengthTypeError
	}
}

func (codec *IsoCodec) EbcdicLen(s string) ([]byte, error) {
	l := len(s)
	switch codec.Length.Type {
	case FIXED:
		if codec.Padding == NONE {
			if l != codec.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if l > codec.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		}
		return []byte{}, nil
	case LLVAR:
		if l == 0 || l > codec.Length.Value || l > 99 {
			fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return AsciiToEbcdic(LeftPad2Len(strconv.Itoa(l), "0", 2)), nil
	case LLLVAR:
		if l == 0 || l > codec.Length.Value || l > 999 {
			fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return AsciiToEbcdic(LeftPad2Len(strconv.Itoa(l), "0", 3)), nil
	default:
		fmt.Printf("Error: %s\n", InvalidLengthTypeError)
		return nil, InvalidLengthTypeError
	}
}

func (codec *IsoCodec) BinaryLen(s string) ([]byte, error) {
	l := len(s) / 2
	switch codec.Length.Type {
	case FIXED:
		if codec.Padding == NONE {
			if l != codec.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if l > codec.Length.Value {
				fmt.Printf("Error FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		}
		return nil, nil
	case LLVAR:
		if l == 0 || l > codec.Length.Value || l > 99 {
			fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return StrToBcd(LeftPad2Len(strconv.Itoa(l), "0", 2)), nil
	case LLLVAR:
		if l == 0 || l > codec.Length.Value || l > 999 {
			fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
			return nil, Errors[InvalidLengthError]
		}
		return StrToBcd(LeftPad2Len(strconv.Itoa(l), "0", 3)), nil
	default:
		return nil, Errors[InvalidLengthError]
	}
}

func (codec *IsoCodec) Pad(s string) (string, error) {
	l := codec.Length.Value
	padStr := " "
	if codec.IsNumeric {
		padStr = "0"
	}

	if codec.Encoding == BINARY {
		l = codec.Length.Value * 2
		padStr = "20"
		if codec.IsNumeric {
			padStr = "00"
		}
	}

	if codec.Padding == LEFT {
		return LeftPad2Len(s, padStr, l), nil
	} else if codec.Padding == RIGHT {
		if codec.IsNumeric {
			return s, NotSupported
		}
		return RightPad2Len(s, padStr, l), nil
	}
	return s, nil
}

func (codec *IsoCodec) Encode(s string) ([]byte, error) {
	fmt.Printf("Input: [%s]\n", s)
	if codec.IsNumeric {
		n := new(big.Int)
		n, ok := n.SetString(s, 10)
		if !ok {
			fmt.Printf("Error: Input is not  numeric: [%s]\n", s)
			return nil, Errors[NumberFormatError]
		}
	}

	data := s
	if codec.Padding != NONE {
		padding, err := codec.Pad(s)
		if err != nil {
			return nil, err
		}
		data = padding
		fmt.Printf("Padding: [%s]\n", data)
	}

	dataLen, err := codec.CodecLen(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("LenBytes: [%X]\n", dataLen)

	switch codec.Encoding {
	case ASCII:
		bytes := append(dataLen, []byte(data)...)
		fmt.Printf("ASCII encoded: [%X]\n", bytes)
		return bytes, nil
	case BINARY:
		b, err := hex.DecodeString(data)
		if err != nil {
			fmt.Printf("Error BINARY: %s\n", Errors[InvalidDataError])
			return nil, Errors[InvalidDataError]
		}
		bytes := append(dataLen, b...)
		fmt.Printf("Binary encoded: [%X]\n", bytes)
		return bytes, nil
	case EBCDIC:
		bytes := append(dataLen, AsciiToEbcdic(data)...)
		fmt.Printf("EBCDIC encoded: [%X]\n", bytes)
		return bytes, nil
	default:
		return nil, NotSupportedEncodingError
	}
}

func (codec *IsoCodec) Decode(b []byte) (string, error) {
	switch codec.Encoding {
	case ASCII:
		switch codec.Length.Encoding {
		case ASCII:
			switch codec.Length.Type {
			case FIXED:
				return decodeAscii(b, codec.Length.Value)
			case LLVAR:
				return decodeAsciiVar(b, 2)
			case LLLVAR:
				return decodeAsciiVar(b, 3)
			default:
				return "", InvalidLengthTypeError
			}
		case EBCDIC:
			return "", nil
		case BINARY:
			return "", nil
		default:
			return "", nil
		}
	case EBCDIC:
		switch codec.Length.Type {
		case FIXED:
			return decodeEbcdic(b, codec.Length.Value)
		case LLVAR:
			return decodeEbcdicVar(b, 2)
		case LLLVAR:
			return decodeEbcdicVar(b, 3)
		default:
			return "", InvalidLengthTypeError
		}
	case BINARY:
		return "", nil
	default:
		return "", NotSupportedEncodingError
	}
}

func decodeAscii(b []byte, fixedLen int) (string, error) {
	if len(b) < fixedLen {
		return "", Errors[InvalidLengthError]
	}
	s := string(b[:fixedLen])
	fmt.Printf("ASCII decoded: [%s]\n", s)
	return s, nil
}

func decodeAsciiVar(b []byte, maxLen int) (string, error) {
	if len(b) < maxLen+1 {
		return "", Errors[InvalidLengthError]
	}
	length, err := strconv.Atoi(string(b[:3]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+maxLen {
		return "", Errors[InvalidLengthError]
	}
	return string(b[maxLen : length+maxLen]), nil
}

func decodeEbcdic(b []byte, fixedLen int) (string, error) {
	if len(b) < fixedLen {
		return "", Errors[InvalidLengthError]
	}
	a := EbcdicToAscii(string(b))
	return string(a), nil
}

func decodeEbcdicVar(b []byte, maxLen int) (string, error) {
	if len(b) < maxLen+1 {
		return "", Errors[InvalidLengthError]
	}
	b = EbcdicToAscii(string(b))
	length, err := strconv.Atoi(string(b[:maxLen]))
	if err != nil || length <= 0 {
		return "", Errors[InvalidLengthError]
	}
	if len(b) < length+maxLen {
		return "", Errors[InvalidLengthError]
	}
	return string(b[maxLen : length+maxLen]), nil
}
