package isocodec

import (
	"../../utils"
)

type IsoType struct {
	Len   *IsoData
	Value *IsoData
}

func (isoType *IsoType) Encode(s string) ([]byte, error) {
	encLen, err := isoType.Len.Encode(s[:isoType.Len.Size()])
	if err != nil {
		return nil, err
	}
	encValue, err := isoType.Value.Encode(s[isoType.Len.Size():])
	if err != nil {
		return nil, err
	}
	return append(encLen, encValue...), nil
}

func (isoType *IsoType) Decode(b []byte) (string, int, error) {
	decLen, sizeOfLen, err := isoType.DecodeLen(b)
	if err != nil {
		return "", 0, err
	}

	if len(b) < sizeOfLen+decLen {
		return "", 0, NotEnoughData
	}

	decValue, _, err := isoType.Value.Decode(b[sizeOfLen:decLen])
	if err != nil {
		return "", 0, err
	}
	return decValue, sizeOfLen + decLen, nil
}

func (isoType *IsoType) BeforeEncoding(s string) error {
	err := isoType.Len.BeforeEncoding(s)
	if err != nil {
		return err
	}

	err = isoType.Value.BeforeEncoding(s)
	if err != nil {
		return err
	}

	return nil
}

func (isoType *IsoType) BeforeDecoding(b []byte) error {
	panic("Implement me")
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

func (isoType *IsoType) PadString() string {
	return ""
}

func (isoType *IsoType) Size() int {
	return 0
}

func (isoType *IsoType) DecodeLen(b []byte) (int, int, error) {
	if isoType.Len == nil {
		return 0, 0, nil
	}

	if len(b) < isoType.Len.Max {
		return 0, isoType.Len.Max, NotEnoughData
	}
	if isoType.Len.Encoding == IsoAscii {
		i, err := utils.Btoi(b[:isoType.Len.Max])
		return i, isoType.Len.Max, err
	} else if isoType.Len.Encoding == IsoEbcdic {
		i, err := utils.Btoi(utils.EbcdicToAsciiBytes(b[:isoType.Len.Max]))
		return i, isoType.Len.Max, err
	} else if isoType.Len.Encoding == IsoBinary {
		i := utils.BcdToInt(b[:isoType.Len.Max])
		return int(i), isoType.Len.Max, nil
	} else {
		return 0, isoType.Len.Max, NotSupported
	}
}
