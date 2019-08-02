package main

import (
	"bytes"
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
func IsoMsgNew(mti string, p Protocol) *IsoMsg {
	bitmap := Bitmap{}
	fields := [MaxField + 1]*IsoField{}
	isoMsg := &IsoMsg{p, bitmap, fields}
	isoMsg.Set(0, mti)
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
	fmt.Printf("%03d: %X\n", pos, isoMsg.fields[pos].value)
	return isoMsg.fields[pos], nil
}

// Value ...
func (isoMsg *IsoMsg) Value(pos int) ([]byte, error) {
	f, err := isoMsg.Get(pos)
	if err != nil {
		return nil, err
	}	
	return f.value, nil
}

// Text ...
func (isoMsg *IsoMsg) Text(pos int) (string, error) {
	f, err := isoMsg.Get(pos)
	if err != nil {
		return "", err
	}	
	return f.text, nil
}

// MTI ...
func (isoMsg *IsoMsg) MTI() (*IsoField, error) {
	return isoMsg.Get(0)
}

// Set ...
func (isoMsg *IsoMsg) Set(pos int, s string) {
	if pos == 1 {
		return
	}
	f, err := IsoFieldNew(pos, s, isoMsg.protocol[pos])
	if err != nil {
		log.Fatal("Creating ", isoMsg.protocol[pos], " failed:", err)
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

// String ...
func (isoMsg *IsoMsg) String() string {
	var buffer bytes.Buffer
	for _, f := range isoMsg.fields {
		if f != nil {
			buffer.WriteString(fmt.Sprintf("%-6s%s\n", f.codec, f))
		}
	}
	return buffer.String()
}

// Encode ...
func (isoMsg *IsoMsg) Encode() ([]byte, error) {
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

// Decode ...
func (isoMsg *IsoMsg) Decode(b []byte) error {	
	return nil
}

func (isoMsg *IsoMsg) refresh() {
	b, err := IsoFieldNew(1, isoMsg.bitmap.String(), isoMsg.protocol[1])
	if err != nil {
		log.Fatal(isoMsg.protocol[1], " ", err)
	}
	isoMsg.fields[1] = b
}