package main

import (
	"fmt"
	"testing"
)

func Test_IsoMsg_New(t *testing.T) {
	isoMsg := IsoMsgNew("0100", ASCII1987)
	f1, err := isoMsg.Get(1)
	assertEqual(t, nil, f1)
	assertEqual(t, IsoFieldNotFoundError, err)
}

func Test_IsoMsg_FieldNotFound(t *testing.T) {
	isoMsg := IsoMsgNew("0100", ASCII1987)
	f1, err := isoMsg.Get(0)
	assertEqual(t, nil, err)
	assertEqual(t, "0100", f1.text)

	f1, err = isoMsg.Get(128)
	assertEqual(t, nil, f1)
	assertEqual(t, IsoFieldNotFoundError, err)
}

func Test_IsoMsg_InvalidIndex(t *testing.T) {
	isoMsg := IsoMsgNew("0100", ASCII1987)
	f1, err := isoMsg.Get(129)
	assertEqual(t, OutOfBoundIndexError, err)
	assertEqual(t, nil, f1)
}

func Test_IsoMsg_MTI(t *testing.T) {
	isoMsg := IsoMsgNew("0100", ASCII1987)
	mti, err := isoMsg.MTI()
	assertEqual(t, nil, err)
	assertEqual(t, 0, mti.pos)
	assertEqual(t, "0100", mti.text)
}

func Test_IsoMsg_AsciiBitmap(t *testing.T) {
	isoMsg := IsoMsgNew("0500", ASCII1987)
	isoMsg.Set(2, "6734000000000000067")
	
	mti, err := isoMsg.Get(0)
	assertEqual(t, nil, err)
	assertEqual(t, "0500", mti.text)

	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, "4000000000000000", bitmap.text)
	assertEqual(t, []byte("00000000000000004000000000000000"), bitmap.value)

	pan, err := isoMsg.Get(2)
	assertEqual(t, nil, err)
	assertEqual(t, "6734000000000000067", pan.text)
	assertEqual(t, []byte("196734000000000000067"), pan.value)
}

func Test_IsoMsg_AsciiBitmapPrimary(t *testing.T) {
	isoMsg := IsoMsgNew("0100", ASCII1987)
	isoMsg.Set(2, "1234567890")
	isoMsg.Set(12, "1234")
	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, []byte("00000000000000004010000000000000"), bitmap.value)
}

func Test_IsoMsg_AsciiBitmapSecondary(t *testing.T) {
	isoMsg := IsoMsgNew("0100", ASCII1987)
	isoMsg.Set(2, "1234567890")
	isoMsg.Set(12, "1234")
	isoMsg.Set(128, "2D2A98F12D2A98F1")

	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, []byte("C0100000000000000000000000000001"), bitmap.value)
}

func Test_IsoMsg_Print(t *testing.T) {
    isoMsg := IsoMsgNew("0100", ASCII1987)
    isoMsg.Set(0, "0100")
	isoMsg.Set(2, "6734000000000000067")
	isoMsg.Set(3, "000000")
	isoMsg.Set(4, "000000000101")
	isoMsg.Set(7, "0806080102")
	isoMsg.Set(11, "823579")
	isoMsg.Set(12, "100100")
	isoMsg.Set(13, "0806")
	isoMsg.Set(19, "528")
	isoMsg.Set(23, "001")
	isoMsg.Set(32, "673005005")
	isoMsg.Set(33, "67390100200")
	isoMsg.Set(37, "080685700000")
	isoMsg.Set(41, "CPOS99  ")
	isoMsg.Set(43, "0650MAES115252800300003540")
	isoMsg.Set(48, "MAES0115941234560407      ")
	isoMsg.Set(49, "978")
	isoMsg.Set(55, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800")
	isoMsg.Set(60, "")
	isoMsg.Set(61, "020058800000000015  001")
	isoMsg.Set(63, "02600000000008005283526LB    ")
	isoMsg.Set(71, "5666")
	isoMsg.Set(128, "2D2A98F12D2A98F1")

	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, []byte("F238220188A1821A0200000000000001"), bitmap.value)

	fmt.Printf("%s\n\n", isoMsg)
}

func Test_IsoMsg_Encode(t *testing.T) {
    isoMsg := IsoMsgNew("0100", ASCII1987)
    isoMsg.Set(0, "0100")
	isoMsg.Set(2, "6734000000000000067")
	isoMsg.Set(3, "000000")
	isoMsg.Set(4, "000000000101")
	isoMsg.Set(7, "0806080102")
	isoMsg.Set(11, "823579")
	isoMsg.Set(12, "100100")
	isoMsg.Set(13, "0806")
	isoMsg.Set(19, "528")
	isoMsg.Set(23, "001")
	isoMsg.Set(32, "673005005")
	isoMsg.Set(33, "67390100200")
	isoMsg.Set(37, "080685700000")
	isoMsg.Set(41, "CPOS99  ")
	isoMsg.Set(43, "0650MAES115252800300003540")
	isoMsg.Set(48, "MAES0115941234560407      ")
	isoMsg.Set(49, "978")
	isoMsg.Set(55, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800")
	isoMsg.Set(60, "")
	isoMsg.Set(61, "020058800000000015  001")
	isoMsg.Set(63, "02600000000008005283526LB    ")
	isoMsg.Set(71, "5666")
	isoMsg.Set(128, "2D2A98F12D2A98F1")

	_, err := isoMsg.Encode()
	assertEqual(t, nil, err)
	// fmt.Printf("%X", encoded)
}

func Test_IsoMsg_Decode(t *testing.T) {
	isoMsg := IsoMsgNew("0100", ASCII1987)
	isoMsg.Set(2, "6734000000000000067")
	isoMsg.Set(3, "000000")
	isoMsg.Set(4, "000000000101")
	isoMsg.Set(7, "0806080102")
	isoMsg.Set(11, "823579")
	isoMsg.Set(12, "100100")
	isoMsg.Set(13, "0806")
	isoMsg.Set(19, "528")
	isoMsg.Set(23, "001")
	isoMsg.Set(32, "673005005")
	isoMsg.Set(33, "67390100200")
	isoMsg.Set(37, "080685700000")
	isoMsg.Set(41, "CPOS99  ")
	isoMsg.Set(43, "0650MAES115252800300003540")
	isoMsg.Set(48, "MAES0115941234560407      ")
	isoMsg.Set(49, "978")
	isoMsg.Set(55, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800")
	isoMsg.Set(60, "")
	isoMsg.Set(61, "020058800000000015  001")
	isoMsg.Set(63, "02600000000008005283526LB    ")
	isoMsg.Set(71, "5666")
	isoMsg.Set(128, "2D2A98F12D2A98F1")

	b, err := isoMsg.Encode()
    assertEqual(t, nil, err)
    fmt.Printf("%X\n\n", b)
    
    err = isoMsg.Decode(b)
    assertEqual(t, nil, err)
}


func Test_IsoMsg_AsciiFields(t *testing.T) {
	isoMsg := IsoMsgNew("0500", ASCII1987)
	isoMsg.Set(2, "6734000000000000067")
	isoMsg.Set(12, "100100")
	isoMsg.Set(128, "2D2A98F12D2A98F1")

	mti, err := isoMsg.MTI()
	assertEqual(t, nil, err)
	assertEqual(t, "0500", mti.text)
	assertEqual(t, []byte{0x30, 0x35, 0x30, 0x30}, mti.value)
	
	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, "C0100000000000000000000000000001", bitmap.text)
	assertEqual(t, []byte("C0100000000000000000000000000001"), bitmap.value)

	pan, err := isoMsg.Get(2)
	assertEqual(t, nil, err)
	assertEqual(t, "6734000000000000067", pan.text)
	assertEqual(t, []byte("196734000000000000067"), pan.value)

	time, err := isoMsg.Get(12)
	assertEqual(t, nil, err)
	assertEqual(t, "100100", time.text)
	assertEqual(t, []byte("100100"), time.value)

	mac, err := isoMsg.Get(128)
	assertEqual(t, nil, err)
	assertEqual(t, "2D2A98F12D2A98F1", mac.text)
	assertEqual(t, []byte("2D2A98F12D2A98F1"), mac.value)
}