package iso

/*
import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

// MaxField ...
const MaxField = 128

// IsoMsg ...
type IsoMsg struct {
	protocol Protocol
	bitmap   Bitmap
	fields   [MaxField + 1]*IsoField
}

// IsoMsgNew ...
func IsoMsgNew(p Protocol) *IsoMsg {
	bitmap := Bitmap{}
	fields := [MaxField + 1]*IsoField{}
	isoMsg := &IsoMsg{p, bitmap, fields}
	return isoMsg
}

// Get ...
func (isoMsg *IsoMsg) Get(pos int) (*IsoField, error) {
	if pos < 0 || pos > MaxField {
		return nil, OutOfBoundIndexError
	}
	if isoMsg.fields[pos] == nil {
		return nil, IsoFieldNotFoundError
	}
	return isoMsg.fields[pos], nil
}

// Set ...
func (isoMsg *IsoMsg) Set(pos int, s string) {
	if pos == 1 {
		_, err := isoMsg.bitmap.Decode(s)
		if err != nil {
			log.Fatal("Incorrect bitmap ", err)
		}
		return
	}

	f, err := IsoFieldNew(pos, s, isoMsg.protocol[pos])
	if err != nil {
		log.Fatal("Creating ", isoMsg.protocol[pos], " failed. ", err)
	}
	isoMsg.fields[pos] = f

	if pos > 1 {
		isoMsg.bitmap.Set(pos)
		isoMsg.refresh()
	}
}

// Clear ...
func (isoMsg *IsoMsg) Clear(pos int) {
	isoMsg.fields[pos] = nil
	isoMsg.bitmap.Clear(pos)
	isoMsg.refresh()
}

// Encode ...
func (isoMsg *IsoMsg) String() string {
	var buffer bytes.Buffer
	for _, f := range isoMsg.fields {
		if f != nil {
			buffer.WriteString(fmt.Sprintf("%-6s%s\n", f.codec, f))
		}
	}
	return buffer.String()
}

// Bytes ...
func (isoMsg *IsoMsg) Bytes() ([]byte, error) {
	var encoded []byte
	mti, err := isoMsg.Get(0)
	if err != nil {
		return nil, err
	}
	encoded = append(encoded, mti.value...)
	for _, i := range isoMsg.bitmap.Array() {
		f, err := isoMsg.Get(i)
		if err != nil {
			return nil, err
		}
		encoded = append(encoded, f.value...)
	}
	return encoded, nil
}

// ParseString ...
func (isoMsg *IsoMsg) ParseString(s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	return isoMsg.Parse(b)
}

// Parse ...
func (isoMsg *IsoMsg) Parse(b []byte) error {
	codecs := isoMsg.protocol
	mti, err := codecs[0].Decode(b)
	if err != nil {
		return err
	}
	isoMsg.Set(0, mti)

	b = b[typeLen(codecs[0])+len(mti):]
	bitmap, err := codecs[1].Decode(b)
	if err != nil {
		return err
	}
	isoMsg.Set(1, bitmap)

	b = b[typeLen(codecs[1])+len(bitmap):]
	for _, i := range isoMsg.bitmap.Array() {
		if i == 1 {
			continue
		}
		value, err := codecs[i].Decode(b)
		if err != nil {
			log.Fatal(i, value, err)
			return err
		}
		isoMsg.Set(i, value)
		b = b[typeLen(codecs[i])+len(value):]
	}

	return nil
}

func typeLen(v interface{}) int {
	switch v.(type) {
	case *LLLANumeric, *LLLBNumeric, *LLLENumeric, *LLLAChar, *LLLBChar, *LLLEChar:
		return 3
	case *LLANumeric, *LLENumeric, *LLBNumeric, *LLAChar, *LLEChar, *LLBChar:
		return 2
	case *LAChar, *LEChar:
		return 1
	default:
		return 0
	}
}

func (isoMsg *IsoMsg) refresh() {
	b, err := IsoFieldNew(1, isoMsg.bitmap.String(), isoMsg.protocol[1])
	if err != nil {
		log.Fatal(isoMsg.protocol[1], " ", err)
	}
	isoMsg.fields[1] = b
}
*/
