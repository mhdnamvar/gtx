package isocodec

import (
	"../../utils"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
)

type IsoData struct {
	Encoding    IsoEncoding
	Min         int
	Max         int
	ContentType IsoContentType
	Padding     IsoPadding
}

func (isoData *IsoData) Encode(s string) ([]byte, error) {
	err := isoData.BeforeEncoding(s)
	if err != nil {
		return nil, err
	}

	p, err := isoData.Pad(s)
	if err != nil {
		return nil, err
	}

	if isoData.Encoding == IsoAscii {
		return []byte(p), nil
	} else if isoData.Encoding == IsoBinary {
		if len(p)%2 != 0 {
			p += "0"
		}
		bytes, err := hex.DecodeString(p)
		if err != nil {
			log.Printf("[%s], err[%v]", p, err)
			return nil, InvalidData
		}
		return bytes, nil
	} else if isoData.Encoding == IsoEbcdic {
		return utils.AsciiToEbcdic(p), nil
	} else {
		return nil, NotSupported
	}
}

func (isoData *IsoData) Decode(b []byte) (string, int, error) {
	if isoData.Encoding == IsoAscii {
		return string(b), len(b), nil
	} else if isoData.Encoding == IsoEbcdic {
		return string(utils.EbcdicToAsciiBytes(b)), len(b), nil
	} else if isoData.Encoding == IsoBinary {
		return strings.ToUpper(hex.EncodeToString(b)), len(b), nil
	} else {
		return "", len(b), NotSupported
	}
}

func (isoData *IsoData) BeforeEncoding(s string) error {
	if len(s) > isoData.Size() {
		if Debug {
			log.Printf("BeforeEncoding: len(s)=%d, isoData.Size()=%d", len(s), isoData.Size())
		}
		return InvalidLength
	}

	if isoData.ContentType != IsoBitmap {
		if isoData.Padding == IsoNoPad && len(s) != isoData.Size() {
			if isoData.Min == isoData.Max {
				if Debug {
					log.Printf("BeforeEncoding: len(%s)=%d should be the same as isoData.Size()=%d",
						s, len(s), isoData.Size())
				}
				return InvalidLength
			}
		}
	}

	// check numeric content
	if isoData.ContentType == IsoNumeric {
		n := new(big.Int)
		n, ok := n.SetString(s, 10)
		if !ok {
			if Debug {
				log.Printf("BeforeEncoding: invalid numeric content")
			}
			return InvalidData
		}
	}

	return nil
}

func (isoData *IsoData) BeforeDecoding(b []byte) error {
	panic("Implement me")
}

func (isoData *IsoData) Pad(s string) (string, error) {
	if isoData.Padding == IsoLeftPad {
		return utils.LeftPad2Len(s, isoData.PadString(), isoData.Size()), nil
	} else if isoData.Padding == IsoRightPad {
		return utils.RightPad2Len(s, isoData.PadString(), isoData.Size()), nil
	} else if isoData.Padding == IsoRightPadF {
		if len(s)%2 != 0 {
			return utils.RightPad(s, "F", 1), nil
		}
		return s, nil
	} else {
		return s, nil
	}
}

func (isoData *IsoData) Size() int {
	size := isoData.Max
	if isoData.Encoding == IsoBinary {
		size *= 2
		if isoData.ContentType == IsoNumeric {
			if isoData.Padding == IsoRightPadF || isoData.Padding == IsoRightPad {
				if size%2 != 0 {
					size += 1
				}
			}
		}
	} else {
		if isoData.ContentType == IsoHexString || isoData.ContentType == IsoBitmap {
			size *= 2
		}
	}
	return size
}

func (isoData *IsoData) PadString() string {
	var p string
	if isoData.Encoding == IsoBinary {
		if isoData.ContentType == IsoNumeric {
			p = "00"
		} else {
			p = "20"
		}
	} else {
		if isoData.ContentType == IsoNumeric {
			p = "0"
		} else {
			p = " "
		}
	}
	return p
}

func (isoData *IsoData) AfterEncoding(b []byte) ([]byte, error) {
	if isoData.Padding == IsoRightPadF {
		if isoData.Encoding == IsoBinary {
			if len(b)%2 != 0 {
				return append(b, []byte{0xF}...), nil
			}
		} else if isoData.Encoding == IsoAscii {
			return []byte(utils.RightPad2Len(string(b), " ", isoData.Max)), nil
		}
	}
	return b, nil
}

func (isoData *IsoData) AfterDecoding(s string) error {
	panic("Implement me")
}

func (isoData *IsoData) String() string {
	return fmt.Sprintf(
		`&IsoData{
			Encoding: %s, 
			Min: %d,
			Max: %d,
			ContentType: %s,
			Padding: %s,
		}`,
		isoData.Encoding,
		isoData.Min,
		isoData.Max,
		isoData.ContentType,
		isoData.Padding,
	)
}
