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
func IsoMsgNew(p Protocol) *IsoMsg {
	bitmap := Bitmap{}
	fields := [MaxField + 1]*IsoField{}
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
	isoField, err := IsoFieldNew(pos, s, isoMsg.protocol[pos])
	if err != nil {
		log.Fatal(isoMsg.protocol[pos], " ", err.Error())
	}
	isoMsg.fields[pos] = isoField
	if pos > 1 {
		isoMsg.bitmap.Set(pos)
		isoMsg.fields[1].text = isoMsg.bitmap.Encode()
	}
}

// Clear ...
func (isoMsg *IsoMsg) Clear(pos int) {
	isoMsg.fields[pos] = nil
	isoMsg.bitmap.Clear(pos)
	isoMsg.fields[1].text = isoMsg.bitmap.Encode()
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
	// mti, err := isoMsg.protocol[0].Decode(b)
	// if err != nil {
	// 	return err
	// }
	// isoMsg.Set(0, mti)
	
	// bitmap, err := isoMsg.protocol[1].Decode(b[len(mti):])
	// if err != nil {
	// 	return err
	// }
	// isoMsg.Set(1, bitmap)

	// for _, i := range isoMsg.bitmap.Array() {
	// 	f, err := isoMsg.Get(i)
	// 	if err != nil {
	// 		return nil, err
	// 	}		
	// 	encoded = append(encoded, f.value...)
	// 	fmt.Printf("%03d: %X\n", i, f.value)
	// }

	return nil
}
