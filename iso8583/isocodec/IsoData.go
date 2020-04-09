package isocodec

import (
	"../../utils"
	"encoding/hex"
	"fmt"
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
	if len(s) > isoData.Max {
		return InvalidLength
	}

	if isoData.Min == isoData.Max && isoData.Padding == IsoNoPad && len(s) != isoData.Max {
		return InvalidLength
	}

	if isoData.ContentType == IsoNumeric {
		n := new(big.Int)
		n, ok := n.SetString(s, 10)
		if !ok {
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
