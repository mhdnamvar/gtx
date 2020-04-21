package iso8583

import (
	"../../utils"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type IsoType struct {
	Len   *IsoData
	Value *IsoData
}

func (isoType *IsoType) Encode(s string) ([]byte, error) {
	i := len(s)
	if isoType.Value.ContentType == IsoHexString {
		i /= 2
	}
	encLen, err := isoType.Len.Encode(strconv.Itoa(i))
	if err != nil {
		color.Red("%v", err)
		return nil, err
	}

	encValue, err := isoType.Value.Encode(s)
	if err != nil {
		color.Red("%v", err)
		return nil, err
	}

	return append(encLen, encValue...), nil
}

func (isoType *IsoType) Decode(b []byte) (string, int, error) {

	color.White("data(%X)", b)

	lenSize, dataSize, err := isoType.DecodeLen(b)
	if err != nil {
		return "", 0, InvalidLength
	}

	// PadSize
	padSize := 0
	if isoType.Value.Padding != IsoNoPad {
		padSize = isoType.Value.Max - dataSize
		if padSize < 0 {
			color.Red("Error: dataSize(%d) > Max(%d), ", dataSize, isoType.Value.Max)
			return "", 0, InvalidLength
		}
	}

	if isoType.Value.Encoding == IsoBinary &&
		dataSize % 2 != 0 &&
		isoType.Value.ContentType == IsoNumeric {
		padSize++
	}

	// Size of data
	size := dataSize
	if isoType.Value.Padding != IsoNoPad {
		size = isoType.Value.Max
	}
	odd := size % 2 != 0
	if isoType.Value.Encoding == IsoBinary &&
		isoType.Value.ContentType != IsoHexString &&
		isoType.Value.ContentType != IsoBitmap{
		if odd {
			size /= 2
			size += 1
		} else {
			size /= 2
		}
	}

	if isoType.Value.Encoding != IsoBinary &&
		isoType.Value.ContentType == IsoHexString &&
		isoType.Len == nil{
		 size *= 2
	}

	color.Cyan("len(%d), lenSize(%d), size(%d), binary(%v)",
		len(b), lenSize, size, isoType.Value.Encoding == IsoBinary)
	if len(b) < lenSize + size {
		color.Red("Error: len(b)=%d < lenSize=%d + size=%d", len(b), lenSize, size)
		return "", 0, InvalidLength
	}

	decValue, _, err := isoType.Value.Decode(b[lenSize : lenSize+size])
	if err != nil {
		return "", 0, err
	}

	// Padding
	if isoType.Value.Padding == IsoLeftPad {
		decValue = decValue[padSize:]
	} else if isoType.Value.Padding == IsoRightPad || isoType.Value.Padding == IsoRightPadF{
		decValue = decValue[:len(decValue) - padSize]
	} else if isoType.Value.Padding == IsoNoPad {
		if isoType.Value.Encoding == IsoBinary && dataSize % 2 != 0 && isoType.Value.ContentType == IsoNumeric {
			decValue = decValue[:len(decValue) - padSize]
		}
	}

	// return values
	if isoType.Value.ContentType == IsoBitmap && isoType.Value.Encoding == IsoBinary{
		var bitmap Bitmap
		err = bitmap.Parse(decValue[:8])
		if err != nil {
			return "", 0, err
		}
		sBitmap := 1
		color.Yellow("Second bitmap = %v", bitmap.Get(sBitmap))
		if !bitmap.Get(sBitmap) {
			size /= 2
		}
	}

	if isoType.Value.Encoding == IsoBinary{
		if isoType.Value.ContentType == IsoHexString && isoType.Len == nil{
			return decValue, size/2, nil
		}
	} else if isoType.Value.ContentType == IsoHexString && isoType.Len == nil{
		return decValue, size*2, nil
	}
	return decValue, lenSize + size, nil
}

func (isoType *IsoType) BeforeEncoding(s string) error {
	//Check size of data to be encoded
	if isoType.Len == nil {// DE has fixed length
		if isoType.Value.Padding == IsoNoPad && isoType.Value.ContentType != IsoBitmap{
			if len(s) != isoType.Value.Max {
				return InvalidLength
			}
		} else {
			if len(s) > isoType.Value.Max {
				return InvalidLength
			}
		}
	} else { // DE has variable length
		if len(s) > isoType.Value.Max {
			return InvalidLength
		}
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
	panic("Not implemented!")
}

func (isoType *IsoType) DecodeLen(b []byte) (int, int, error) {
	if isoType.Len == nil {
		color.Cyan("Fixed length")
		return 0, isoType.Value.Max, nil
	}
	color.Cyan("Variable length")
	if isoType.Len.Encoding == IsoAscii {
		i, err := utils.Btoi(b[:isoType.Len.Max])
		return isoType.Len.Max, i, err
	} else if isoType.Len.Encoding == IsoEbcdic {
		i, err := utils.Btoi(utils.EbcdicToAsciiBytes(b[:isoType.Len.Max]))
		return isoType.Len.Max, i, err
	} else if isoType.Len.Encoding == IsoBinary {
		l := isoType.Len.Max/2
		if isoType.Len.Max %2 != 0 {
			l++
		}
		i := utils.BcdToInt(b[:l])
		return l, int(i), nil
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
