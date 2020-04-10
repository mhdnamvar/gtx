package isocodec

import (
	"../../utils"
	"fmt"
	"strconv"
	"strings"
)

type IsoType struct {
	Len   *IsoData
	Value *IsoData
}

func (isoType *IsoType) Encode(s string) ([]byte, error) {
	var encLen []byte
	if isoType.Len != nil {
		dataLen := len(s)
		if isoType.Value.ContentType == IsoHexString ||
			isoType.Value.ContentType == IsoBitmap {
			dataLen /= 2
		}
		l, err := isoType.Len.Encode(strconv.Itoa(dataLen))
		if err != nil {
			return nil, err
		}
		encLen = l
	}

	encValue, err := isoType.Value.Encode(s)
	if err != nil {
		return nil, err
	}

	return append(encLen, encValue...), nil
}

func (isoType *IsoType) Decode(b []byte) (string, int, error) {
	lenSize, decLen, err := isoType.DecodeLen(b)
	if err != nil {
		return "", 0, err
	}

	if len(b) < lenSize+decLen {
		return "", 0, NotEnoughData
	}

	decValue, _, err := isoType.Value.Decode(b[lenSize : lenSize+decLen])
	if err != nil {
		return "", 0, err
	}

	return decValue, lenSize + decLen, nil

	//afterDecoding, err := isoType.AfterDecoding(decValue)
	//if err != nil {
	//	return "", 0, err
	//}
	//return afterDecoding, lenSize + decLen, nil

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
	err := isoType.Len.BeforeDecoding(b)
	if err != nil {
		return err
	}

	err = isoType.Value.BeforeDecoding(b)
	if err != nil {
		return err
	}

	return nil
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
	size := isoType.Value.Max
	if isoType.Value.Encoding != IsoBinary {
		if isoType.Value.ContentType == IsoBitmap || isoType.Value.ContentType == IsoHexString {
			size *= 2
		}
	}
	return size
}

func (isoType *IsoType) DecodeLen(b []byte) (int, int, error) {
	if isoType.Len == nil {
		size := isoType.Size()
		if len(b) < size &&	isoType.Value.ContentType != IsoBitmap{
			return 0, size, NotEnoughData
		}
		return 0, size, nil
	}

	if len(b) < isoType.Len.Max {
		return 0, isoType.Len.Max, NotEnoughData
	}
	if isoType.Len.Encoding == IsoAscii {
		i, err := utils.Btoi(b[:isoType.Len.Max])
		return isoType.Len.Max, i,  err
	} else if isoType.Len.Encoding == IsoEbcdic {
		i, err := utils.Btoi(utils.EbcdicToAsciiBytes(b[:isoType.Len.Max]))
		return isoType.Len.Max, i,  err
	} else if isoType.Len.Encoding == IsoBinary {
		i := utils.BcdToInt(b[:isoType.Len.Max])
		if isoType.Value.ContentType == IsoNumeric || isoType.Value.Padding == IsoRightPadF{
			if i%2 != 0 {
				i+=1
			}
		}
		if isoType.Value.Encoding == IsoBinary && isoType.Value.ContentType != IsoHexString{
			i /= 2
		}
		return isoType.Len.Max, int(i), nil
	} else {
		return isoType.Len.Max, 0, NotSupported
	}
}

func (isoType *IsoType) AfterEncoding(b []byte) ([]byte, error) {
	return isoType.Value.AfterEncoding(b)
}

func (isoType *IsoType) AfterDecoding(decValue string) (string, error) {
	if isoType.Value.Padding == IsoRightPadF {
		if strings.EqualFold(decValue[len(decValue)-1:], "F") {
			decValue = decValue[:len(decValue)-1]
		}
	}
	return decValue, nil
}

func (isoType *IsoType) String() string {
	if isoType.Len != nil {
		return fmt.Sprintf(
`&IsoType{
		Len: %v, 
		Value: %v,
	}`, isoType.Len, isoType.Value)
	}
	return fmt.Sprintf(
`&IsoType{
		Value: %v,
	}`, isoType.Value)
}
