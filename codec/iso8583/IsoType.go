package iso8583

import (
	"../../utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type IsoType struct {
	Len   *IsoData
	Value *IsoData
}

func (isoType *IsoType) Encode(s string) ([]byte, error) {
	//var encLen []byte
	//dataLen := len(s)
	//if isoType.Len != nil {
	//	//log.Println(dataLen)
	//	if isoType.Value.Padding != IsoNoPad {
	//		if dataLen > isoType.Value.Max {
	//			log.Printf("Error: Variable length and dataLen(%d) > isoType.Value.Max(%d)",
	//				dataLen, isoType.Value.Max)
	//			return nil, InvalidLength
	//		}
	//		dataLen = isoType.Value.Max
	//	}
	//	//log.Println(dataLen)
	//	//if isoType.Value.ContentType == IsoHexString || isoType.Value.ContentType == IsoBitmap {
	//	//	dataLen /= 2
	//	//}
	//	log.Println(dataLen)
	//	l, err := isoType.Len.Encode(strconv.Itoa(dataLen))
	//	if err != nil {
	//		log.Println(err)
	//		return nil, err
	//	}
	//	encLen = l
	//}
	//
	//if dataLen > isoType.Value.Max {
	//	log.Printf("Error: Fixed length and dataLen(%d) > isoType.Value.Max(%d)",
	//		dataLen, isoType.Value.Max)
	//	return nil, InvalidLength
	//}

	i := len(s)
	if isoType.Value.ContentType == IsoHexString {
		i /= 2
	}
	encLen, err := isoType.Len.Encode(strconv.Itoa(i))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	encValue, err := isoType.Value.Encode(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return append(encLen, encValue...), nil
}

func (isoType *IsoType) Decode(b []byte) (string, int, error) {

	log.Printf("data=(%X)", b)

	lenSize, dataSize, err := isoType.DecodeLen(b)
	if err != nil {
		return "", 0, InvalidLength
	}

	log.Printf("len(b)=%d, lenSize=%d, dataSize=%d", len(b), lenSize, dataSize)

	padSize := 0
	if isoType.Value.Padding != IsoNoPad {
		padSize = isoType.Value.Max - dataSize
		if padSize < 0 {
			log.Printf("Error: dataSize(%d) > Max(%d), ", dataSize, isoType.Value.Max)
			return "", 0, InvalidLength
		}
	}
	//if isoType.Value.Padding == IsoRightPadF {
	//	padSize += 1
	//}

	size := dataSize
	if isoType.Value.Padding != IsoNoPad {
		size = isoType.Value.Max
	}

	odd := size % 2 != 0
	if isoType.Value.Encoding == IsoBinary && isoType.Value.ContentType != IsoHexString{
		if odd {
			size /= 2
			size += 1
		} else {
			size /= 2
		}
	}

	log.Printf("len(b)=%d, lenSize=%d, dataSize=%d, binary=%v",
		len(b), lenSize, size, isoType.Value.Encoding == IsoBinary)


	if len(b) < lenSize + size {
		log.Printf("Error: len(b)=%d < lenSize=%d + size=%d", len(b), lenSize, size)
		return "", 0, InvalidLength
	}

	decValue, _, err := isoType.Value.Decode(b[lenSize : lenSize+size])
	if err != nil {
		return "", 0, err
	}

	log.Printf("padSize=%d", padSize)
	if isoType.Value.Padding == IsoLeftPad {
		decValue = decValue[padSize:]
	} else if isoType.Value.Padding == IsoRightPad {
		decValue = decValue[:len(decValue) - padSize]
	}

	if isoType.Value.Encoding == IsoBinary {
		if odd && isoType.Value.ContentType != IsoHexString{
			if isoType.Value.Padding == IsoRightPadF{
				log.Println("------------1")
				return decValue[:len(decValue)-1], (lenSize + dataSize)/2+1, nil
			} else if isoType.Value.Padding == IsoLeftPad{
				log.Println("------------2")
				if isoType.Value.ContentType == IsoNumeric {
					return decValue[:len(decValue)-1], (lenSize + dataSize)/2, nil
				}
				return decValue, (lenSize + dataSize)/2, nil
			} else {
				log.Println("------------3")
				return decValue[:len(decValue)-1], (lenSize + dataSize)/2, nil
			}
		} else {
			log.Println("------------4")
			return decValue, (lenSize + dataSize)/2, nil
		}
	} else {
		log.Println("------------5")
		return decValue, lenSize + dataSize, nil
	}

	//afterDecoding, err := isoType.AfterDecoding(decValue)
	//if err != nil {
	//	return "", 0, err
	//}
	//return afterDecoding, lenSize + dataSize, nil

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
	//size := isoType.Value.Max
	//if isoType.Value.Encoding != IsoBinary {
	//	if isoType.Value.ContentType == IsoBitmap || isoType.Value.ContentType == IsoHexString {
	//		size *= 2
	//	}
	//}
	//return size
	panic("Not implemented!")
}

func (isoType *IsoType) DecodeLen(b []byte) (int, int, error) {
	if isoType.Len == nil {
		//size := isoType.Value.Max
		//if len(b) < size && isoType.Value.ContentType != IsoBitmap {
		//	return 0, size, NotEnoughData
		//}
		//return 0, size, nil
		log.Print("Fixed length")
		return 0, isoType.Value.Max, nil
	}

	if len(b) < isoType.Len.Max {
		log.Print("---------len(b) < isoType.Len.Max---------")
		return 0, isoType.Len.Max, NotEnoughData
	}

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
		//if isoType.Value.ContentType == IsoNumeric || isoType.Value.Padding == IsoRightPadF {
		//	if i%2 != 0 {
		//		i += 1
		//	}
		//}
		//if isoType.Value.Encoding == IsoBinary && isoType.Value.ContentType != IsoHexString {
		//	i /= 2
		//}
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
