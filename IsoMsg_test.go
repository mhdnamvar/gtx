package main

import (
    "testing"
    "fmt"
)

func Test_IsoMsg_New(t *testing.T) {
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

func Test_IsoMsg_Print(t *testing.T) {
    isoMsg := IsoMsgNew(ASCII1987)
    isoMsg.Set(2, "6734000000000000067");
    isoMsg.Set(3, "000000");
    isoMsg.Set(4, "000000000101");
    isoMsg.Set(7, "0806080102");
    isoMsg.Set(11, "823579");
    isoMsg.Set(12, "100100");
    isoMsg.Set(13, "0806");
    isoMsg.Set(19, "528");
    isoMsg.Set(23, "001");
    isoMsg.Set(32, "673005005");
    isoMsg.Set(33, "67390100200");
    isoMsg.Set(37, "080685700000");
    isoMsg.Set(41, "CPOS99  ");
    isoMsg.Set(43, "0650MAES115252800300003540");
    isoMsg.Set(48, "MAES0115941234560407      ");
    isoMsg.Set(49, "978");
    isoMsg.Set(55, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800")
    isoMsg.Set(60, "");
    isoMsg.Set(61, "020058800000000015  001");
    isoMsg.Set(63, "02600000000008005283526LB    ");
    isoMsg.Set(71, "5666");
    isoMsg.Set(128, "2D2A98F12D2A98F1");
    
    bitmap, err := isoMsg.Get(1)
    assertEqual(t, nil, err)
    assertEqual(t, "F238220188A1821A0200000000000001", bitmap.value)

    fmt.Printf("%s", isoMsg)
}

func Test_IsoMsg_Encode(t *testing.T) {
    isoMsg := IsoMsgNew(ASCII1987)
    isoMsg.Set(2, "6734000000000000067");
    isoMsg.Set(3, "000000");
    isoMsg.Set(4, "000000000101");
    isoMsg.Set(7, "0806080102");
    isoMsg.Set(11, "823579");
    isoMsg.Set(12, "100100");
    isoMsg.Set(13, "0806");
    isoMsg.Set(19, "528");
    isoMsg.Set(23, "001");
    isoMsg.Set(32, "673005005");
    isoMsg.Set(33, "67390100200");
    isoMsg.Set(37, "080685700000");
    isoMsg.Set(41, "CPOS99  ");
    isoMsg.Set(43, "0650MAES115252800300003540");
    isoMsg.Set(48, "MAES0115941234560407      ");
    isoMsg.Set(49, "978");
    isoMsg.Set(55, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800")
    isoMsg.Set(60, "");
    isoMsg.Set(61, "020058800000000015  001");
    isoMsg.Set(63, "02600000000008005283526LB    ");
    isoMsg.Set(71, "5666");
    isoMsg.Set(128, "2D2A98F12D2A98F1");
    
    _, err := isoMsg.Encode()
    assertEqual(t, nil, err)   
    // fmt.Printf("%X", encoded)
}