package main

import "log"

// IsoMsg ...
type IsoMsg struct {
	protocol Protocol
	bitmap   Bitmap
	fields   [129]*IsoField
}

// IsoMsgNew ...
func IsoMsgNew(p Protocol) *IsoMsg {
	bitmap := Bitmap{}
	fields := [129]*IsoField{}
	fields[1], _ = IsoFieldNew(1, bitmap.Encode())
	return &IsoMsg{p, bitmap, fields}
}

// Set ...
func (isoMsg *IsoMsg) Set(pos int, s string) {
	isoField, err := IsoFieldNew(pos, s)
	if err != nil {
		log.Fatal(err.Error())
	}
	isoMsg.fields[pos] = isoField
	isoMsg.bitmap.Set(pos)
	isoMsg.fields[1].value = isoMsg.bitmap.Encode()
}

// UnSet ...
func (isoMsg *IsoMsg) UnSet(pos int) {
	isoMsg.fields[pos] = nil
	isoMsg.bitmap.Clear(pos)
	isoMsg.fields[1].value = isoMsg.bitmap.Encode()
}

// Encode ...
func (isoMsg *IsoMsg) Encode(s string) ([]byte, error) {
	return nil, nil
}

// Decode ...
func (isoMsg *IsoMsg) Decode(b []byte) (string, error) {
	return "", nil
}
