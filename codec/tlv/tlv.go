package tlv

import (
	"../../utils"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type TlvParser interface {
	Name() string
	Lookup(tag string, data string) Tlv
	Decode(data string) (decoded map[string]Tlv)
	Encode(tag string, data string) string
	Dump(w io.Writer, data string)
	New() DefaultTlvParser
}

type DefaultTlvParser struct{}

type Tlv struct {
	Tag        string
	LengthCode string
	Length     string
	Value      string
}

func (tlv *Tlv) Size() int {
	return len(tlv.Tag) + len(tlv.LengthCode) + len(tlv.Value)
}

func (defaultTlvParser *DefaultTlvParser) Encode(tag string, data string) string {
	length := len(data) / 2
	lenHex := fmt.Sprintf("%X", length)
	lenBytes := len(lenHex) / 2

	if length < 127 {
		return fmt.Sprintf("%s%s%s", tag, utils.LeftPad2Len(lenHex, "0", 2), data)
	} else if length < 255 {
		return fmt.Sprintf("%s%s%s%s", tag, "81", utils.LeftPad2Len(lenHex, "0", 2), data)
	} else if lenBytes < 2 {
		return fmt.Sprintf("%s%s%s%s", tag, "82", utils.LeftPad2Len(lenHex, "0", 4), data)
	} else if lenBytes < 3 {
		return fmt.Sprintf("%s%s%s%s", tag, "83", utils.LeftPad2Len(lenHex, "0", 6), data)
	} else if lenBytes < 4 {
		return fmt.Sprintf("%s%s%s%s", tag, "84", utils.LeftPad2Len(lenHex, "0", 8), data)
	} else {
		fmt.Printf("ERROR: Incorrect length")
		return ""
	}

}

func (tlv *Tlv) Print() {
	fmt.Printf("%-4s %-s %-s\n", tlv.Tag, tlv.Length, tlv.Value)
}

func NewTlvParser() DefaultTlvParser {
	return DefaultTlvParser{}
}

func (defaultTlvParser *DefaultTlvParser) Name() string {
	return "DefaultTlvParser"
}

func (defaultTlvParser *DefaultTlvParser) New() DefaultTlvParser {
	return DefaultTlvParser{}
}

func (defaultTlvParser *DefaultTlvParser) Lookup(tag string, data string) Tlv {
	decoded := defaultTlvParser.Decode(data)
	return decoded[strings.ToUpper(tag)]
}

func getTlv(emvData string) Tlv {
	data := strings.Replace(strings.ToUpper(emvData), " ", "", -1)
	tag := data[0:2]
	if utils.Hex2Bin(data[0:2])[3:8] == "11111" {
		tag = data[0:4]
		for i := 2; utils.Hex2Dec(data[i:i+2]) >= 128; i = i + 2 {
			tag = data[0 : i+4]
		}
	}

	lengthHex := data[len(tag) : len(tag)+2]
	valueIndex := len(tag) + 2
	lengthCode := lengthHex
	if lengthHex[0:1] == "8" {
		lengthBytes, _ := strconv.Atoi(lengthHex[1:])
		lengthHex = data[len(tag)+2 : len(tag)+2+lengthBytes*2]
		valueIndex += lengthBytes * 2
		lengthCode += lengthHex
	}
	length := utils.Hex2Dec(lengthHex) * 2

	if valueIndex+length > len(data) {
		fmt.Printf("*** ERROR: Invalid EMV data length(%d)\n", len(data))
		os.Exit(1)
	}
	value := ""
	if length > 0 {
		value = data[valueIndex : valueIndex+length]
	}
	return Tlv{tag, lengthCode, lengthHex, value}
}

func (defaultTlvParser *DefaultTlvParser) Decode(data string) (decoded map[string]Tlv) {
	decoded = make(map[string]Tlv)
	for data != "" {
		tlv := getTlv(data)
		decoded[tlv.Tag] = Tlv{tlv.Tag, tlv.LengthCode, tlv.Length, tlv.Value}
		if utils.Hex2Bin(tlv.Tag)[2:3] == "1" {
			innerTlv := getTlv(tlv.Value)
			decoded[innerTlv.Tag] = Tlv{innerTlv.Tag, tlv.LengthCode, innerTlv.Length, innerTlv.Value}
			limit := len(tlv.Value)
			index := innerTlv.Size()
			for index < limit {
				innerTlv = getTlv(tlv.Value[index:])
				decoded[innerTlv.Tag] = Tlv{innerTlv.Tag, tlv.LengthCode, innerTlv.Length, innerTlv.Value}
				index += innerTlv.Size()
			}
		}
		data = data[tlv.Size():]
	}
	return decoded
}

func (defaultTlvParser *DefaultTlvParser) Dump(w io.Writer, data string) (n int, err error) {
	decoded := defaultTlvParser.Decode(data)
	var buffer bytes.Buffer
	for _, tlv := range decoded {
		buffer.WriteString(fmt.Sprintf("%-4s %-s %-s\n", tlv.Tag, tlv.Length, tlv.Value))
	}
	return fmt.Fprintln(w, buffer.String())
}
