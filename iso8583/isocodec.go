package iso8583

import (
	"encoding/hex"
	"log"
	"math/big"
	"strconv"
	"strings"

	"../utils"
)

type IsoCodec struct {
	Name        string
	Description string
	LenCodec    *IsoCodec
	Encoding    IsoEncoding
	Size        int
	ContentType IsoContentType
	Padding     IsoPadding
}

type IsoSpec []*IsoCodec
type IsoEncoding int
type IsoPadding int
type IsoContentType int

const (
	IsoAscii  IsoEncoding = 0
	IsoEbcdic IsoEncoding = 1
	IsoBinary IsoEncoding = 2

	IsoNoPad    IsoPadding = 0
	IsoLeftPad  IsoPadding = 1
	IsoRightPad IsoPadding = 2

	IsoText    IsoContentType = 0
	IsoNumeric IsoContentType = 1
)

var (
	IsoFixed *IsoCodec = &IsoCodec{}
	IsoLLA   *IsoCodec = &IsoCodec{"LLA", "ASCII variable length, max 99", IsoFixed, IsoAscii, 2, IsoNumeric, IsoLeftPad}
	IsoLLE   *IsoCodec = &IsoCodec{"LLE", "EBCDIC variable length, max 99", IsoFixed, IsoEbcdic, 2, IsoNumeric, IsoLeftPad}
	IsoLLB   *IsoCodec = &IsoCodec{"LLB", "Binary variable length, max 99", IsoFixed, IsoBinary, 1, IsoNumeric, IsoLeftPad}
	IsoLLLA  *IsoCodec = &IsoCodec{"LLLA", "ASCII variable length, max 999", IsoFixed, IsoAscii, 3, IsoNumeric, IsoLeftPad}
	IsoLLLE  *IsoCodec = &IsoCodec{"LLLE", "EBCDIC variable length, max 999", IsoFixed, IsoEbcdic, 3, IsoNumeric, IsoLeftPad}
	IsoLLLB  *IsoCodec = &IsoCodec{"LLLB", "Binary variable length, max 999", IsoFixed, IsoBinary, 2, IsoNumeric, IsoLeftPad}
)

func pad(codec *IsoCodec, s string) (string, error) {
	var size int
	var p string
	if codec.Encoding == IsoBinary {
		size = codec.Size * 2
		p = "20"
		if codec.ContentType == IsoNumeric {
			p = "00"
		}
	} else {
		size = codec.Size
		p = " "
		if codec.ContentType == IsoNumeric {
			p = "0"
		}
	}

	if codec.Padding == IsoLeftPad {
		return utils.LeftPad2Len(s, p, size), nil
	} else if codec.Padding == IsoRightPad {
		if codec.ContentType == IsoNumeric {
			return s, NotSupported
		}
		return utils.RightPad2Len(s, p, size), nil
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
	if codec.LenCodec.Size != IsoFixed.Size {
		l := len(s)
		if codec.Encoding == IsoBinary {
			l = len(s) / 2
			if codec.LenCodec.Size == IsoLLB.Size {
				if l > codec.Size || l > 99 {
					return Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == IsoLLLB.Size {
				if l > codec.Size || l > 999 {
					return Errors[InvalidLengthError]
				}
			} else {
				return Errors[InvalidLengthError]
			}
		} else {
			if codec.LenCodec.Size == IsoLLA.Size || codec.LenCodec.Size == IsoLLE.Size {
				if l > codec.Size || l > 99 {
					return Errors[InvalidLengthError]
				}
			} else if codec.LenCodec.Size == IsoLLLA.Size || codec.LenCodec.Size == IsoLLLE.Size {
				if l > codec.Size || l > 999 {
					return Errors[InvalidLengthError]
				}
			} else {
				return Errors[InvalidLengthError]
			}
		}
		return nil
	} else {
		l := len(s)
		if codec.Encoding == IsoBinary {
			l = len(s) / 2
		}
		if codec.Padding == IsoNoPad {
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

	if codec.LenCodec.Size != IsoFixed.Size {
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
	if codec.Encoding == IsoBinary {
		l = len(s) / 2
	}
	return pad(codec.LenCodec, strconv.Itoa(l))
}

func doEncode(codec *IsoCodec, s string) ([]byte, error) {
	if codec.ContentType == IsoNumeric {
		n := new(big.Int)
		n, ok := n.SetString(s, 10)
		if !ok {
			return nil, Errors[NumberFormatError]
		}
	}
	if codec.Encoding == IsoAscii {
		return []byte(s), nil
	} else if codec.Encoding == IsoBinary {
		if codec.ContentType == IsoNumeric {
			return utils.StrToBcd(s), nil
		}
		bytes, err := hex.DecodeString(s)
		if err != nil {
			return nil, Errors[InvalidDataError]
		}
		return bytes, nil
	} else if codec.Encoding == IsoEbcdic {
		return utils.AsciiToEbcdic(s), nil
	} else {
		return nil, NotSupportedEncodingError
	}
}

func (codec *IsoCodec) Decode(b []byte) (string, int, error) {
	l, n, err := decodeLen(codec, b)
	if err != nil {
		return "", 0, err
	}

	if len(b) < len(l)+n {
		return "", 0, NotEnoughData
	}

	data := b[len(l) : len(l)+n]
	if codec.LenCodec.Size == IsoFixed.Size && len(data) != codec.Size {
		return "", 0, Errors[InvalidLengthError]
	} else if len(data) > codec.Size {
		return "", 0, Errors[InvalidLengthError]
	}

	if codec.Encoding == IsoAscii {
		return string(data), len(data), nil
	} else if codec.Encoding == IsoEbcdic {
		return string(utils.EbcdicToAsciiBytes(data)), len(data), nil
	} else if codec.Encoding == IsoBinary {
		return strings.ToUpper(hex.EncodeToString(data)), len(data), nil
	} else {
		return "", 0, NotSupportedEncodingError
	}
}

func decodeLen(codec *IsoCodec, b []byte) ([]byte, int, error) {
	if codec.LenCodec.Size == IsoFixed.Size {
		return []byte{}, codec.Size, nil
	}

	if codec.LenCodec.Encoding == IsoAscii {
		if codec.LenCodec.Size == IsoLLA.Size {
			if len(b) < IsoLLA.Size {
				log.Fatalf("Invalid IsoAscii LLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:IsoLLA.Size]
			i, err := utils.Btoi(bytes)
			return bytes, i, err
		} else if codec.LenCodec.Size == IsoLLLA.Size {
			if len(b) < IsoLLLA.Size {
				log.Fatalf("Invalid IsoAscii LLLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:IsoLLLA.Size]
			i, err := utils.Btoi(bytes)
			return bytes, i, err
		} else {
			return nil, 0, Errors[InvalidLengthError]
		}
	} else if codec.LenCodec.Encoding == IsoEbcdic {
		if codec.LenCodec.Size == IsoLLE.Size {
			if len(b) < IsoLLE.Size {
				log.Fatalf("Invalid IsoEbcdic LLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := utils.EbcdicToAsciiBytes(b[:IsoLLE.Size])
			i, err := utils.Btoi(bytes)
			return bytes, i, err
		} else if codec.LenCodec.Size == IsoLLLE.Size {
			if len(b) < IsoLLLE.Size {
				log.Fatalf("Invalid IsoEbcdic LLLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := utils.EbcdicToAsciiBytes(b[:IsoLLLE.Size])
			i, err := utils.Btoi(bytes)
			return bytes, i, err
		} else {
			return nil, 0, Errors[InvalidLengthError]
		}
	} else if codec.LenCodec.Encoding == IsoBinary {
		if codec.LenCodec.Size == IsoLLB.Size {
			if len(b) < IsoLLB.Size {
				log.Fatalf("Invalid IsoBinary LLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:IsoLLB.Size]
			n := utils.BcdToInt(bytes)
			return bytes, int(n), nil
		} else if codec.LenCodec.Size == IsoLLLB.Size {
			if len(b) < IsoLLLB.Size {
				log.Fatalf("Invalid IsoBinary LLLVar: %X", b)
				return nil, 0, Errors[InvalidLengthError]
			}
			bytes := b[:IsoLLLB.Size]
			n := utils.BcdToInt(bytes)
			return bytes, int(n), nil
		}
	} else {
		return nil, 0, NotSupportedEncodingError
	}
	return nil, 0, Errors[InvalidLengthError]
}
