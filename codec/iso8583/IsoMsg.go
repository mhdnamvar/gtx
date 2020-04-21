package iso8583

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
	"log"
)

const (
	MaxField = 128
)

type IsoMsg struct {
	bitmap *Bitmap
	fields map[int]string
}

func NewIsoMsg() *IsoMsg {
	return &IsoMsg{&Bitmap{}, make(map[int]string, MaxField)}
}

func (isoMsg IsoMsg) String() string {
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
		return "", FieldNotFound
	}
	return isoMsg.fields[i], nil
}

func (isoMsg *IsoMsg) Encode(isoSpec IsoSpec) ([]byte, error) {
	mti, err := isoSpec[0].Encode(isoMsg.fields[0])
	if err != nil {
		log.Fatal("Error in encoding DE000, ", err)
		return nil, err
	}

	color.White("DE000(%X)", mti)

	bitmap, err := isoSpec[1].Encode(isoMsg.fields[1])
	if err != nil {
		log.Fatalf("Error in encoding DE001: [%s], %v", isoMsg.fields[1], err)
		return nil, err
	}
	color.White("DE001(%X)", bitmap)

	var b []byte
	fields := isoMsg.bitmap.Array()
	for _, f := range fields {
		if f == 1 { // skip the bitmap
			continue
		}
		encoded, err := isoSpec[f].Encode(isoMsg.fields[f])
		if err != nil {
			log.Fatalf("Error in encoding DE%03d(%s), %v", f, isoMsg.fields[f], err)
			return nil, err
		}
		color.White("DE%03d(%X)", f, encoded)
		b = append(b, encoded...)
	}
	return append(append(mti, bitmap...), b...), nil
}

func (isoMsg *IsoMsg) Decode(isoSpec IsoSpec, b []byte) error {
	offset := 0
	s, mtiLen, err := isoSpec[0].Decode(b)
	if err != nil {
		log.Fatal("Error in decoding DE000")
		return err
	}
	isoMsg.Set(0, s)
	color.Green("DE%03d(%s)", 0, s)
	offset = mtiLen
	s, bitmapLen, err := isoSpec[1].Decode(b[offset:])
	if err != nil {
		log.Fatal("Error in decoding DE001")
		return err
	}

	err = isoMsg.bitmap.Parse(s)
	if err != nil {
		log.Fatalf("Error in parsing bitmap: [%s], %v", s, err)
		return err
	}
	color.Green("DE%03d(%s), bitmapLen(%d)", 1, s, bitmapLen)
	offset = mtiLen + bitmapLen
	for _, f := range isoMsg.bitmap.Array() {
		if f > 1 {
			s, dataLen, err := isoSpec[f].Decode(b[offset:])
			if err != nil {
				//log.Fatalf("DE%03d, Error: %v", f, err)
				return err
			}
			color.Green("DE%03d(%s)", f, s)
			isoMsg.Set(f, s)
			offset = offset + dataLen
		}
	}
	return nil
}

func (isoMsg *IsoMsg) Parse(isoSpec IsoSpec, s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		log.Fatalf("Error in parsing, %v", err)
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
