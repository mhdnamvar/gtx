package main

import (
	"encoding/hex"
	"math/big"
	"strconv"
)

type IsoCodec struct {
	LenCodec  *IsoCodec
	Encoding  IsoEncoding
	Size      int
	Padding   IsoPadding
	IsNumeric bool
}

type IsoEncoding int
type IsoPadding int

const (
	ASCII  IsoEncoding = 0
	BINARY IsoEncoding = 1
	EBCDIC IsoEncoding = 2

	NoPadding    IsoPadding = 0
	LeftPadding  IsoPadding = 1
	RightPadding IsoPadding = 2
)

func IsoText(lenCodec *IsoCodec, encoding IsoEncoding, size int, padding IsoPadding) *IsoCodec {
	return &IsoCodec{lenCodec, encoding, size, padding, false}
}

func IsoNumeric(lenCodec *IsoCodec, encoding IsoEncoding, size int, padding IsoPadding) *IsoCodec {
	return &IsoCodec{lenCodec, encoding,size, padding, true}
}

func FIXED() *IsoCodec {
	return &IsoCodec{}
}

func LLVAR(encoding IsoEncoding) *IsoCodec {
	size := 2
	if encoding == BINARY {
		size = 1
	}
	return &IsoCodec{FIXED(), encoding, size, LeftPadding, true}
}

func LLLVAR(encoding IsoEncoding) *IsoCodec {
	size := 3
	if encoding == BINARY {
		size = 2
	}
	return &IsoCodec{FIXED(), encoding, size, LeftPadding, true}
}

func (codec *IsoCodec) Pad(s string) (string, error) {
	var size int
	var pad string
	if codec.Encoding == BINARY {
		size = codec.Size * 2
		pad = "20"
		if codec.IsNumeric {
			pad = "00"
		}
	} else {
		size = codec.Size
		pad = " "
		if codec.IsNumeric {
			pad = "0"
		}
	}

	if codec.Padding == LeftPadding {
		return LeftPad2Len(s, pad, size), nil
	} else if codec.Padding == RightPadding {
		if codec.IsNumeric {
			return s, NotSupported
		}
		return RightPad2Len(s, pad, size), nil
	} else {
		return s, nil
	}
}

func (codec *IsoCodec) Encode(s string) ([]byte, error) {
	str, err := codec.Pad(s)
	if err != nil {
		return nil, err
	}

	length, err := codec.EncodeLen(s, str)
	if err != nil {
		return nil, err
	}

	value, err := codec.execute(str)
	if err != nil {
		return nil, err
	}

	return append(length, value...), nil
}

func (codec *IsoCodec) CheckLen(s string) error {
	if codec.LenCodec.Size != 0 {
		l := len(s)
		if codec.Encoding == BINARY {
			l = len(s) / 2
			if codec.LenCodec.Size == 1 {
				if l == 0 || l > codec.Size || l > 99 {
					return Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == 2 {
				if l == 0 || l > codec.Size || l > 999 {
					return Errors[InvalidLengthError]
				}
			}
		} else {
			if codec.LenCodec.Size == 2 {
				if l == 0 || l > codec.Size || l > 99 {
					return Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == 3 {
				if l == 0 || l > codec.Size || l > 999 {
					return Errors[InvalidLengthError]
				}
			}
		}
		return nil
	} else {
		l := len(s)
		if codec.Encoding == BINARY {
			l = len(s) / 2
		}
		if codec.Padding == NoPadding {
			if l != codec.Size {
				return Errors[InvalidLengthError]
			}
		} else {
			if l > codec.Size {
				return Errors[InvalidLengthError]
			}
		}
		return nil
	}
}

func (codec *IsoCodec) EncodeLen(s string, str string) ([]byte, error) {
	err := codec.CheckLen(s)
	if err != nil {
		return nil, err
	}

	if codec.LenCodec.Size != 0 { // LLVAR and LLLVAR
		lenPad, err := codec.padLen(str)
		if err != nil {
			return nil, err
		}
		bytes, err := codec.LenCodec.execute(lenPad)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	} else {
		return []byte{}, nil
	}
}

func (codec *IsoCodec) padLen(s string) (string, error) {
	l := len(s)
	if codec.Encoding == BINARY {
		l = len(s) / 2
	}
	return codec.LenCodec.Pad(strconv.Itoa(l))
}

func (codec *IsoCodec) execute(s string) ([]byte, error) {
	if codec.IsNumeric {
		n := new(big.Int)
		n, ok := n.SetString(s, 10)
		if !ok {
			return nil, Errors[NumberFormatError]
		}
	}
	if codec.Encoding == ASCII {
		return []byte(s), nil
	} else if codec.Encoding == BINARY {
		if codec.IsNumeric {
			return StrToBcd(s), nil
		}
		bytes, err := hex.DecodeString(s)
		if err != nil {
			return nil, Errors[InvalidDataError]
		}
		return bytes, nil
	} else if codec.Encoding == EBCDIC {
		return AsciiToEbcdic(s), nil
	} else {
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
	log.Printf("ASCII decoded: [%s]\n", s)
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
