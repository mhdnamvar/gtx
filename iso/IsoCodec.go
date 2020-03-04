package main

import (
	"encoding/hex"
	"log"
	"math/big"
	"strconv"
	"strings"
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

	// Encoding type
	ASCII  IsoEncoding = 0
	BINARY IsoEncoding = 1
	EBCDIC IsoEncoding = 2

	// Padding type
	NoPadding    IsoPadding = 0
	LeftPadding  IsoPadding = 1
	RightPadding IsoPadding = 2

	// FIXED length
	FixSize int = 0

	// LLVAR Length
	LLVarSize       int = 2
	LLVarBinarySize int = 1

	// LLLVAR Length
	LLLVarSize       int = 3
	LLLVarBinarySize int = 2
)

func IsoText(lenCodec *IsoCodec, encoding IsoEncoding, size int, padding IsoPadding) *IsoCodec {
	return &IsoCodec{lenCodec, encoding, size, padding, false}
}

func IsoNumeric(lenCodec *IsoCodec, encoding IsoEncoding, size int, padding IsoPadding) *IsoCodec {
	return &IsoCodec{lenCodec, encoding, size, padding, true}
}

func FIXED() *IsoCodec {
	return &IsoCodec{}
}

func LLVAR(encoding IsoEncoding) *IsoCodec {
	size := LLVarSize
	if encoding == BINARY {
		size = LLVarBinarySize
	}
	return &IsoCodec{FIXED(), encoding, size, LeftPadding, true}
}

func LLLVAR(encoding IsoEncoding) *IsoCodec {
	size := LLLVarSize
	if encoding == BINARY {
		size = LLLVarBinarySize
	}
	return &IsoCodec{FIXED(), encoding, size, LeftPadding, true}
}

func pad(codec *IsoCodec, s string) (string, error) {
	var size int
	var p string
	if codec.Encoding == BINARY {
		size = codec.Size * 2
		p = "20"
		if codec.IsNumeric {
			p = "00"
		}
	} else {
		size = codec.Size
		p = " "
		if codec.IsNumeric {
			p = "0"
		}
	}

	if codec.Padding == LeftPadding {
		return LeftPad2Len(s, p, size), nil
	} else if codec.Padding == RightPadding {
		if codec.IsNumeric {
			return s, NotSupported
		}
		return RightPad2Len(s, p, size), nil
	} else {
		return s, nil
	}
}

func (codec *IsoCodec) Encode(s string) ([]byte, error) {
	str, err := pad(codec, s)
	if err != nil {
		return nil, err
	}

	length, err := encodeLen(codec, s, str)
	if err != nil {
		return nil, err
	}

	value, err := doEncode(codec, str)
	if err != nil {
		return nil, err
	}

	return append(length, value...), nil
}

func checkLen(codec *IsoCodec, s string) error {
	if codec.LenCodec.Size != FixSize {
		l := len(s)
		if codec.Encoding == BINARY {
			l = len(s) / 2
			if codec.LenCodec.Size == LLVarBinarySize {
				if l == 0 || l > codec.Size || l > 99 {
					return Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == LLLVarBinarySize {
				if l == 0 || l > codec.Size || l > 999 {
					return Errors[InvalidLengthError]
				}
			} else {
				return Errors[InvalidLengthError]
			}
		} else {
			if codec.LenCodec.Size == LLVarSize {
				if l == 0 || l > codec.Size || l > 99 {
					return Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == LLLVarSize {
				if l == 0 || l > codec.Size || l > 999 {
					return Errors[InvalidLengthError]
				}
			} else {
				return Errors[InvalidLengthError]
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

func encodeLen(codec *IsoCodec, s string, str string) ([]byte, error) {
	err := checkLen(codec, s)
	if err != nil {
		return nil, err
	}

	if codec.LenCodec.Size != FixSize {
		lenPad, err := padLen(codec, str)
		if err != nil {
			return nil, err
		}
		bytes, err := doEncode(codec.LenCodec, lenPad)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	} else {
		return []byte{}, nil
	}
}

func padLen(codec *IsoCodec, s string) (string, error) {
	l := len(s)
	if codec.Encoding == BINARY {
		l = len(s) / 2
	}
	return pad(codec.LenCodec, strconv.Itoa(l))
}

func doEncode(codec *IsoCodec, s string) ([]byte, error) {
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

func (codec *IsoCodec) Decode(b []byte) (string, error) {
	bytes, n, err := decodeLen(codec, b)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if len(b) < len(bytes)+n {
		return "", NotEnoughData
	}

	data := b[len(bytes) : len(bytes)+n]
	if codec.LenCodec.Size == FixSize && len(data) != codec.Size {
		return "", Errors[InvalidLengthError]
	} else if len(data) > codec.Size {
		return "", Errors[InvalidLengthError]
	}

	if codec.Encoding == ASCII {
		return string(data), nil
	} else if codec.Encoding == EBCDIC {
		return string(EbcdicToAsciiBytes(data)), nil
	} else if codec.Encoding == BINARY {
		return strings.ToUpper(hex.EncodeToString(data)), nil
	} else {
		return "", NotSupportedEncodingError
	}
}

func decodeLen(codec *IsoCodec, b []byte) ([]byte, int, error) {
	if codec.LenCodec.Size == FixSize {
		return []byte{}, codec.Size, nil
	}

	if codec.LenCodec.Encoding == ASCII {
		if codec.LenCodec.Size == LLVarSize {
			if len(b) < LLVarSize {
				log.Fatalf("Invalid ASCII LLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:LLVarSize]
			i, err := Btoi(bytes)
			return bytes, i, err
		} else if codec.LenCodec.Size == LLLVarSize {
			if len(b) < LLLVarSize {
				log.Fatalf("Invalid ASCII LLLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:LLLVarSize]
			i, err := Btoi(bytes)
			return bytes, i, err
		} else {
			return nil, 0, Errors[InvalidLengthError]
		}
	} else if codec.LenCodec.Encoding == EBCDIC {
		if codec.LenCodec.Size == LLVarSize {
			if len(b) < LLVarSize {
				log.Fatalf("Invalid EBCDIC LLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := EbcdicToAsciiBytes(b[:LLVarSize])
			i, err := Btoi(bytes)
			return bytes, i, err
		} else if codec.LenCodec.Size == LLLVarSize {
			if len(b) < LLLVarSize {
				log.Fatalf("Invalid EBCDIC LLLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := EbcdicToAsciiBytes(b[:LLLVarSize])
			i, err := Btoi(bytes)
			return bytes, i, err
		} else {
			return nil, 0, Errors[InvalidLengthError]
		}
	} else if codec.LenCodec.Encoding == BINARY {
		if codec.LenCodec.Size == LLVarBinarySize {
			if len(b) < LLVarBinarySize {
				log.Fatalf("Invalid BINARY LLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:LLVarBinarySize]
			n := BcdToInt(bytes)
			return bytes, int(n), nil
		} else if codec.LenCodec.Size == LLLVarBinarySize {
			if len(b) < LLLVarBinarySize {
				log.Fatalf("Invalid BINARY LLLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:LLLVarBinarySize]
			n := BcdToInt(bytes)
			return bytes, int(n), nil
		}
	} else {
		return nil, 0, NotSupportedEncodingError
	}
	return nil, 0, Errors[InvalidLengthError]
}
