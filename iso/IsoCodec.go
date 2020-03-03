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

	value, err := codec.doEncode(str)
	if err != nil {
		return nil, err
	}

	return append(length, value...), nil
}

func (codec *IsoCodec) CheckLen(s string) error {
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

func (codec *IsoCodec) EncodeLen(s string, str string) ([]byte, error) {
	err := codec.CheckLen(s)
	if err != nil {
		return nil, err
	}
	if codec.LenCodec.Size != FixSize { // LLVAR and LLLVAR
		lenPad, err := codec.padLen(str)
		if err != nil {
			return nil, err
		}
		bytes, err := codec.LenCodec.doEncode(lenPad)
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

func (codec *IsoCodec) doEncode(s string) ([]byte, error) {
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
	n, err := codec.DecodeLen(b)
	if err != nil {
		return "", err
	}
	decode, err := codec.doDecode(b[n:])
	fmt.Printf("value [%s]\n", decode)
	return decode, err
}

func (codec *IsoCodec) doDecode(b []byte) (string, error) {
	fmt.Printf("decode [%X]\n", b)
	if codec.LenCodec.Size == FixSize && len(b) != codec.Size {
		return "", Errors[InvalidLengthError]
	} else if len(b) > codec.Size {
		return "", Errors[InvalidLengthError]
	}
	fmt.Printf("len is ok\n")
	if codec.Encoding == ASCII {
		return string(b), nil
	} else if codec.Encoding == EBCDIC {
		return string(EbcdicToAsciiBytes(b)), nil
	} else if codec.Encoding == BINARY {
		return hex.EncodeToString(b), nil
	} else {
		return "", NotSupportedEncodingError
	}
}

func (codec *IsoCodec) DecodeLen(b []byte) (int, error) {
	if codec.LenCodec.Size == FixSize {
		return 0, nil
	}
	var s string
	if codec.LenCodec.Encoding == ASCII {
		if codec.LenCodec.Size == LLVarSize {
			if len(b) != LLVarSize {
				return 0, Errors[InvalidLengthError]
			}
			s = string(b)
		} else if codec.LenCodec.Size == LLLVarSize {
			if len(b) != LLLVarSize {
				return 0, Errors[InvalidLengthError]
			}
			s = string(b)
		}
		return 0, Errors[InvalidLengthError]
	} else if codec.LenCodec.Encoding == EBCDIC {
		if codec.LenCodec.Size == LLVarSize {
			if len(b) != LLVarSize {
				return 0, Errors[InvalidLengthError]
			}
			s = string(EbcdicToAsciiBytes(b))
		} else if codec.LenCodec.Size == LLLVarSize {
			if len(b) != LLLVarSize {
				return 0, Errors[InvalidLengthError]
			}
			s = string(EbcdicToAsciiBytes(b))
		}
		return 0, Errors[InvalidLengthError]
	} else if codec.LenCodec.Encoding == BINARY {
		if codec.LenCodec.Size == LLVarBinarySize {
			if len(b) != LLVarBinarySize {
				return 0, Errors[InvalidLengthError]
			}
			s = string(BcdToInt(b))
		} else if codec.LenCodec.Size == LLLVarBinarySize {
			if len(b) != LLLVarBinarySize {
				return 0, Errors[InvalidLengthError]
			}
			s = string(BcdToInt(b))
		}
	} else {
		return 0, NotSupportedEncodingError
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n, nil
}
