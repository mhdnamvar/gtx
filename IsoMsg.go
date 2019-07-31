package main

import (
	"bytes"
	"log"
	"fmt"
)

const MaxField = 128
// IsoMsg ...
type IsoMsg struct {
	protocol Protocol
	bitmap   Bitmap
	fields   [MaxField+1]*IsoField
}

// IsoMsgNew ...
func IsoMsgNew(p Protocol) *IsoMsg {
	bitmap := Bitmap{}
	fields := [MaxField+1]*IsoField{}
	isoMsg := &IsoMsg{p, bitmap, fields}
	isoMsg.Set(1, bitmap.Encode())
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

// MTI
func (isoMsg *IsoMsg) MTI() (*IsoField, error) {
	return isoMsg.Get(0)
}

// Set ...
func (isoMsg *IsoMsg) Set(pos int, s string) {
	isoField, err := IsoFieldNew(pos, s)
	if err != nil {
		log.Fatal(err.Error())
	}
	isoMsg.fields[pos] = isoField
	if pos > 1 {
		isoMsg.bitmap.Set(pos)
		isoMsg.fields[1].value = isoMsg.bitmap.Encode()
	}	
}

// Clear ...
func (isoMsg *IsoMsg) Clear(pos int) {
	isoMsg.fields[pos] = nil
	isoMsg.bitmap.Clear(pos)	
	isoMsg.fields[1].value = isoMsg.bitmap.Encode()
}

func (isoMsg *IsoMsg) String() string {
	var buffer bytes.Buffer
    for i, f := range isoMsg.fields {
		if f != nil {
			buffer.WriteString(fmt.Sprintf("%-6s%s\n", isoMsg.protocol[i].GetName(), f.value))
		}		
    }
    return buffer.String()
}

// Encode ...
func (isoMsg *IsoMsg) Encode(s string) ([]byte, error) {
	return nil, nil
}

// Decode ...
func (isoMsg *IsoMsg) Decode(b []byte) (string, error) {
	return "", nil
}
