package iso8583

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

const MaxField = 128

type IsoMsg struct {
	bitmap *IsoBitmap
	fields map[int]string
}

func IsoMsgNew() *IsoMsg {
	return &IsoMsg{&IsoBitmap{}, make(map[int]string, MaxField)}
}

func (isoMsg *IsoMsg) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("DE000=\"%s\"\n", isoMsg.fields[0]))
	fields := isoMsg.bitmap.Array()
	for _, f := range fields {
		buffer.WriteString(fmt.Sprintf("DE%03d=\"%s\"\n", f, isoMsg.fields[f]))
	}
	return buffer.String()
}

func (isoMsg *IsoMsg) Set(i int, s string) {
	if i < 0 || i > MaxField {
		return
	}
	if i != 1 {
		isoMsg.bitmap.Set(i)
		isoMsg.fields[1] = isoMsg.bitmap.String()
	}
	isoMsg.fields[i] = s
}

func (isoMsg *IsoMsg) Get(i int) (string, error) {
	if i < 0 || i > MaxField {
		return "", IsoFieldNotFoundError
	}
	return isoMsg.fields[i], nil
}

func (isoMsg *IsoMsg) Encode(p IsoProtocol) ([]byte, error) {
	mti, err := p[0].Encode(isoMsg.fields[0])
	if err != nil {
		return nil, err
	}

	bitmap, err := p[1].Encode(isoMsg.fields[1])
	if err != nil {
		return nil, err
	}

	var b []byte
	fields := isoMsg.bitmap.Array()
	for _, f := range fields {
		encoded, err := p[f].Encode(isoMsg.fields[f])
		if err != nil {
			log.Println("DE:", f)
			return nil, err
		}
		b = append(b, encoded...)
	}
	return append(append(mti, bitmap...), b...), nil
}

func (isoMsg *IsoMsg) Decode(p IsoProtocol, b []byte) error {
	offset := 0
	s, i, err := p[0].Decode(b)
	if err != nil {
		return err
	}
	//log.Printf("DE000=\"%s\"", s)
	isoMsg.Set(0, s)

	offset = i
	s, j, err := p[1].Decode(b[offset:])
	if err != nil {
		return err
	}
	//log.Printf("DE001=\"%s\"", s)

	err = isoMsg.bitmap.Parse(s)
	if err != nil {
		log.Println(err)
		return err
	}

	offset = p[0].LenCodec.Size + i + p[1].LenCodec.Size + j
	for _, f := range isoMsg.bitmap.Array() {
		if f > 1 {
			//log.Printf("offset=%d, b=%X", offset, b[offset:])
			s, n, err := p[f].Decode(b[offset:])
			if err != nil {
				return err
			}
			//log.Printf("DE%03d=\"%s\"", f, s)
			isoMsg.Set(f, s)
			offset = offset + p[f].LenCodec.Size + n
		}
	}
	return nil
}

func (isoMsg *IsoMsg) Parse(p IsoProtocol, s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	return isoMsg.Decode(p, b)
}

/*
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
