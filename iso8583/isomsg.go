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

func (isoMsg *IsoMsg) Encode(isoSpec IsoSpec) ([]byte, error) {
	mti, err := isoSpec[0].Encode(isoMsg.fields[0])
	if err != nil {
		return nil, err
	}

	bitmap, err := isoSpec[1].Encode(isoMsg.fields[1])
	if err != nil {
		return nil, err
	}

	var b []byte
	fields := isoMsg.bitmap.Array()
	for _, f := range fields {
		encoded, err := isoSpec[f].Encode(isoMsg.fields[f])
		if err != nil {
			log.Println("DE:", f)
			return nil, err
		}
		b = append(b, encoded...)
	}
	return append(append(mti, bitmap...), b...), nil
}

func (isoMsg *IsoMsg) Decode(isoSpec IsoSpec, b []byte) error {
	offset := 0
	s, i, err := isoSpec[0].Decode(b)
	if err != nil {
		return err
	}
	//log.Printf("DE000=\"%s\"", s)
	isoMsg.Set(0, s)

	offset = i
	s, j, err := isoSpec[1].Decode(b[offset:])
	if err != nil {
		return err
	}
	//log.Printf("DE001=\"%s\"", s)

	err = isoMsg.bitmap.Parse(s)
	if err != nil {
		log.Println(err)
		return err
	}

	offset = isoSpec[0].LenCodec.Size + i + isoSpec[1].LenCodec.Size + j
	for _, f := range isoMsg.bitmap.Array() {
		if f > 1 {
			//log.Printf("offset=%d, b=%X", offset, b[offset:])
			s, n, err := isoSpec[f].Decode(b[offset:])
			if err != nil {
				return err
			}
			//log.Printf("DE%03d=\"%s\"", f, s)
			isoMsg.Set(f, s)
			offset = offset + isoSpec[f].LenCodec.Size + n
		}
	}
	return nil
}

func (isoMsg *IsoMsg) Parse(isoSpec IsoSpec, s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	return isoMsg.Decode(isoSpec, b)
}

func (isoMsg *IsoMsg) Dump(isoSpec IsoSpec) (string, error) {
	enc, err := isoMsg.Encode(isoSpec)
	if err != nil {
		return "", err
	}
	return hex.Dump(enc), nil
}
