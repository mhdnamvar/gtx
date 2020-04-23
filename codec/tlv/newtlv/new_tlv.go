package newtlv

import (
	. "../../../utils"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
)

type TLV struct {
	Tag    string
	Length string
	Value  string
}

func (tlv *TLV) Encode() ([]byte, error) {

	var buf bytes.Buffer

	tag, err := hex.DecodeString(tlv.Tag)
	if err != nil {
		return nil, err
	}

	i := hex.DecodedLen(len(tlv.Value))
	length := IntToBcd(uint64(i))
	tlv.Length = hex.EncodeToString(length)

	value, err := hex.DecodeString(tlv.Value)
	if err != nil {
		return nil, err
	}

	buf.Write(tag)
	buf.Write(length)
	buf.Write(value)

	return buf.Bytes(), nil
}

func (tlv *TLV) Parse(s string) error {
	//color.Blue("%X", s[:2])
	i := Hex2Dec(s[:2])
	fmt.Println(s[:2], i)
	n := i - 80
	if n > 0 {
		color.Blue("Len is more than a byte, %d", n)
		tlv.Length = s[2:n]
		//l := binary.BigEndian.Uint64(data[1:n])
	} else {
		tlv.Length = s[:2]
	}
	return nil
}
