package main

import (
	"strconv"
)

type IsoText struct {
	Encoding IsoEncoding
	Name     string
	Desc     string
	Length   *IsoLength
	Padding  IsoPadding
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

func (isoText *IsoText) DataLen(s string) ([]byte, error) {
	length := []byte{}
	switch isoText.Length.Type {
	case FIXED:
		if isoText.Padding == NONE {
			if len(s) != isoText.Length.Value {
				return nil, Errors[InvalidLengthError]
			}
		} else {
			if len(s) > isoText.Length.Value {
				return nil, Errors[InvalidLengthError]
			}
		}
	case LLVAR:
		if len(s) == 0 || len(s) > isoText.Length.Value || len(s) > 99 {
			return nil, Errors[InvalidLengthError]
		}
		length = []byte(LeftPad2Len(strconv.Itoa(len(s)), "0", 2))
	case LLLVAR:
		if len(s) == 0 || len(s) > isoText.Length.Value || len(s) > 999 {
			return nil, Errors[InvalidLengthError]
		}
		length = []byte(LeftPad2Len(strconv.Itoa(len(s)), "0", 3))
	default:
		return nil, InvalidLengthTypeError
	}

	return length, nil
}

func (isoText *IsoText) Encode(s string) ([]byte, error) {
	if isoText.Padding == LEFT {
		s = LeftPad2Len(s, " ", isoText.Length.Value)
	} else if isoText.Padding == RIGHT {
		s = RightPad2Len(s, " ", isoText.Length.Value)
	}

	dataLen, err := isoText.DataLen(s)
	if err != nil {
		return nil, err
	}

	switch isoText.Encoding {
	case ASCII:
		return append(dataLen, []byte(s)...), nil
	case BINARY:
		return append(dataLen, StrToBcd((s))...), nil
	case EBCDIC:
		return append(dataLen, ASCIIToEbcdic(s)...), nil
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
