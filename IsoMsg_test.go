package main

import (
    "testing"
)

func Test_IsoMsg(t *testing.T) {
    isoMsg := IsoMsgNew(ASCII1987)
    f1, err := isoMsg.Get(1)
    assertEqual(t, nil, err)
    assertEqual(t, 1, f1.pos)
    assertEqual(t, "0000000000000000", f1.value)
}

func Test_IsoMsg_FieldNotFound(t *testing.T) {
    isoMsg := IsoMsgNew(ASCII1987)
    f1, err := isoMsg.Get(0)
    assertEqual(t, nil, f1)
    assertEqual(t, IsoFieldNotFoundError, err)    

    f1, err = isoMsg.Get(128)
    assertEqual(t, nil, f1)
    assertEqual(t, IsoFieldNotFoundError, err)    
}

func Test_IsoMsg_InvalidIndex(t *testing.T) {
    isoMsg := IsoMsgNew(ASCII1987)
    f1, err := isoMsg.Get(129)    
    assertEqual(t, OutOfBoundIndexError, err)
    assertEqual(t, nil, f1)
}

func Test_IsoMsg_MTI(t *testing.T) {
    value := "0800"
    isoMsg := IsoMsgNew(ASCII1987)
    isoMsg.Set(0, value)
    mti, err := isoMsg.MTI()
    assertEqual(t, nil, err)
    assertEqual(t, 0, mti.pos)
    assertEqual(t, value, mti.value)
    assertEqual(t, "0000000000000000", isoMsg.bitmap.Encode())
}

func Test_IsoMsg_Bitmap(t *testing.T) {
    isoMsg := IsoMsgNew(ASCII1987)
    isoMsg.Set(0, "0800")
    isoMsg.Set(2, "1234567890")
    isoMsg.Set(12, "1234")
    bitmap, err := isoMsg.Get(1)
    assertEqual(t, nil, err)
    assertEqual(t, "4010000000000000", bitmap.value)

    isoMsg.Set(128, "2D2A98F12D2A98F1")
    bitmap, err = isoMsg.Get(1)
    assertEqual(t, nil, err)
    assertEqual(t, "C0100000000000000000000000000001", bitmap.value)
}