package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

type IsoCodec struct {
	LenCodec  *IsoCodec
	Encoding  IsoEncoding
	Name      string
	Desc      string
	Size      int
	Padding   IsoPadding
	IsNumeric bool
}

type IsoEncoding int
type IsoPadding int

const (
	FIXED_SIZE  int = 0
	LLVAR_SIZE  int = 2
	LLLVAR_SIZE int = 3

	ASCII  IsoEncoding = 0
	BINARY IsoEncoding = 1
	EBCDIC IsoEncoding = 2

	NONE  IsoPadding = 0
	LEFT  IsoPadding = 1
	RIGHT IsoPadding = 2
)

func IsoTextCodec(lenCodec *IsoCodec, encoding IsoEncoding, name string, desc string, size int, padding IsoPadding) *IsoCodec {
	return &IsoCodec{lenCodec, encoding, name, desc, size, padding, false}
}

func IsoNumericCodec(lenCodec *IsoCodec, encoding IsoEncoding, name string, desc string, size int, padding IsoPadding) *IsoCodec {
	return &IsoCodec{lenCodec, encoding, name, desc, size, padding, true}
}

func FIXED() *IsoCodec {
	return &IsoCodec{}
}

func LLVAR(encoding IsoEncoding) *IsoCodec {
	len := LLVAR_SIZE
	if encoding == BINARY {
		len = 1
	}
	return &IsoCodec{FIXED(), encoding, "LLVAR", "LLVAR length encoder", len, LEFT, true}
}

func LLLVAR(encoding IsoEncoding) *IsoCodec {
	len := LLLVAR_SIZE
	if encoding == BINARY {
		len = 2
	}
	return &IsoCodec{FIXED(), encoding, "LLLVAR", "LLLVAR length encoder", len, LEFT, true}
}

func (codec *IsoCodec) Pad(s string) (string, error) {
	l := codec.Size
	padStr := " "
	if codec.IsNumeric {
		padStr = "0"
	}

	if codec.Encoding == BINARY {
		l = codec.Size * 2
		padStr = "20"
		if codec.IsNumeric {
			padStr = "00"
		}
	}

	switch codec.Padding {
	case LEFT:
		return LeftPad2Len(s, padStr, l), nil
	case RIGHT:
		if codec.IsNumeric {
			return s, NotSupported
		}
		return RightPad2Len(s, padStr, l), nil
	default:
		return s, nil
	}
}

func (codec *IsoCodec) Encode(s string) ([]byte, error) {
	fmt.Printf("Input: [%s]\n", s)

	// pad data
	data, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Padding: [%s]\n", data)

	// pad & encode length
	var length []byte
	if codec.LenCodec.Size != FIXED_SIZE {
		// validate data length
		l := len(s)

		if codec.Encoding == BINARY {
			l = len(s) / 2
			if codec.LenCodec.Size == 1 {
				if l == 0 || l > codec.Size || l > 99 {
					fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
					return nil, Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == 2 {
				if l == 0 || l > codec.Size || l > 999 {
					fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
					return nil, Errors[InvalidLengthError]
				}
			}
		} else {
			if codec.LenCodec.Size == LLVAR_SIZE {
				if l == 0 || l > codec.Size || l > 99 {
					fmt.Printf("Error LLVAR: %s\n", Errors[InvalidLengthError])
					return nil, Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == LLLVAR_SIZE {
				if l == 0 || l > codec.Size || l > 999 {
					fmt.Printf("Error LLLVAR: %s\n", Errors[InvalidLengthError])
					return nil, Errors[InvalidLengthError]
				}
			}
		}

		// pad length
		l = len(data)
		if codec.Encoding == BINARY {
			l = len(data) / 2
		}
		lenPad, err := codec.LenCodec.Pad(strconv.Itoa(l))
		if err != nil {
			return nil, err
		}
		fmt.Printf("lenPad: [%s]\n", lenPad)

		// encode length
		dataLen, err := codec.LenCodec.execute(lenPad)
		if err != nil {
			return nil, err
		}
		length = dataLen
	} else {
		// validate FIXED_SIZE length
		l := len(s)
		if codec.Encoding == BINARY {
			l = len(s) / 2
		}
		if codec.Padding == NONE {
			if l != codec.Size {
				fmt.Printf("Error in FIXED length: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if l > codec.Size {
				fmt.Printf("Error in FIXED: %s\n", Errors[InvalidLengthError])
				return nil, Errors[InvalidLengthError]
			}
		}
	}
	fmt.Printf("Length: [%X]\n", length)

	// encode value
	value, err := codec.execute(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Value: [%X]\n", value)

	encoded := append(length, value...)
	fmt.Printf("Encoded: [%X]\n", encoded)
	return encoded, nil
}

func (codec *IsoCodec) execute(s string) ([]byte, error) {
	if codec.IsNumeric {
		n := new(big.Int)
		n, ok := n.SetString(s, 10)
		if !ok {
			fmt.Printf("Error: Input is not  numeric: [%s]\n", s)
			return nil, Errors[NumberFormatError]
		}
	}

	switch codec.Encoding {
	case ASCII:
		return []byte(s), nil
	case BINARY:
		if codec.IsNumeric {
			return StrToBcd(s), nil
		}
		bytes, err := hex.DecodeString(s)
		if err != nil {
			fmt.Printf("Error in BINARY encoding value: %s\n", Errors[InvalidDataError])
			return nil, Errors[InvalidDataError]
		}
		return bytes, nil
	case EBCDIC:
		return AsciiToEbcdic(s), nil
	default:
		return nil, NotSupportedEncodingError
	}

}

/*
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
*/
