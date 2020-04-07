package isocodec

import (
	"../../utils"
	"encoding/hex"
)

type IsoData struct {
	Encoding    IsoEncoding
	Min         int
	Max         int
	ContentType IsoContentType
	Padding     IsoPadding
}

type IsoType struct {
	Len   *IsoData
	Value *IsoData
}

type IsoCodec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, int, error)
	Check(string) error
	Pad(s string) (string, error)
	PadString() string
	Size() int
}

type IsoSpec [][]*IsoCodec
type IsoEncoding int
type IsoPadding int
type IsoContentType int
type IsoPadString string

const (
	IsoAscii  IsoEncoding = 0
	IsoEbcdic IsoEncoding = 1
	IsoBinary IsoEncoding = 2

	IsoNoPad    IsoPadding = 0
	IsoLeftPad  IsoPadding = 1
	IsoRightPad IsoPadding = 2

	IsoString    IsoContentType = 0
	IsoNumeric   IsoContentType = 1
	IsoHexString IsoContentType = 3
	IsoAmount    IsoContentType = 4
	IsoBitmap    IsoContentType = 5
	IsoTrack2    IsoContentType = 6
	IsoTrack3    IsoContentType = 7
)

var (
	StringA = &IsoType{
		nil,
		&IsoData{IsoAscii, 0, 4, IsoString, IsoNoPad},
	}

	LLAStringA = &IsoType{
		&IsoData{IsoAscii, 0, 0, IsoString, IsoNoPad},
		&IsoData{IsoAscii, 0, 0, IsoString, IsoNoPad},
	}
)

func (isoType *IsoType) Encode(s string) ([]byte, error) {
	encLen, err := isoType.Len.Encode(s[:isoType.Len.Size()])
	if err != nil {
		return nil, err
	}
	encValue, err := isoType.Len.Encode(s[isoType.Len.Size():])
	if err != nil {
		return nil, err
	}
	return append(encLen, encValue...), nil
}

func (isoType *IsoType) Decode([]byte) (string, int, error) {
	panic("implement me")
}

func (isoType *IsoType) Check(string) error {
	panic("implement me")
}

func (isoType *IsoType) Pad(s string) (string, error) {
	size := isoType.Len.Size()
	padLen, err := isoType.Len.Pad(s[:size])
	if err != nil {
		return "", err
	}

	padValue, err := isoType.Value.Pad(s[size:])
	if err != nil {
		return "", err
	}

	return padLen + padValue, nil
}

func (isoData *IsoData) Encode(s string) ([]byte, error) {
	p, err := isoData.Pad(s)
	if err != nil {
		return nil, err
	}

	if isoData.Encoding == IsoAscii {
		return []byte(p), nil
	} else if isoData.Encoding == IsoBinary {
		if isoData.ContentType == IsoNumeric {
			return utils.StrToBcd(p), nil
		}
		bytes, err := hex.DecodeString(p)
		if err != nil {
			return nil, InvalidData
		}
		return bytes, nil
	} else if isoData.Encoding == IsoEbcdic {
		return utils.AsciiToEbcdic(p), nil
	} else {
		return nil, NotSupported
	}
}

func (isoData *IsoData) Decode([]byte) (string, int, error) {
	panic("implement me")
}

func (isoData *IsoData) Check(string) error {
	panic("implement me")
}

func (isoData *IsoData) Pad(s string) (string, error) {
	if isoData.Padding == IsoLeftPad {
		return utils.LeftPad2Len(s, isoData.PadString(), isoData.Size()), nil
	} else if isoData.Padding == IsoRightPad {
		return utils.RightPad2Len(s, isoData.PadString(), isoData.Size()), nil
	} else {
		return s, nil
	}
}

func (isoData *IsoData) Size() int {
	var size int
	if isoData.Encoding == IsoBinary {
		size = isoData.Max * 2
	} else {
		size = isoData.Max
	}
	return size
}

func (isoData *IsoData) PadString() string {
	var p string
	if isoData.Encoding == IsoBinary {
		if isoData.ContentType == IsoNumeric {
			p = "00"
		} else if isoData.ContentType == IsoAmount {
			p = "FF"
		} else {
			p = "20"
		}
	} else {
		if isoData.ContentType == IsoNumeric {
			p = "0"
		} else if isoData.ContentType == IsoAmount {
			p = "F"
		} else {
			p = " "
		}
	}
	return p
}
