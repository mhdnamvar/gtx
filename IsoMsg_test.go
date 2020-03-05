package main

import (
	"fmt"
	"testing"
)

func Test_IsoMsg_New(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0100")
	f1, err := isoMsg.Get(1)
	assertEqual(t, nil, f1)
	assertEqual(t, IsoFieldNotFoundError, err)
}

func Test_IsoMsg_FieldNotFound(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0100")
	f1, err := isoMsg.Get(0)
	assertEqual(t, nil, err)
	assertEqual(t, "0100", f1.text)

	f1, err = isoMsg.Get(128)
	assertEqual(t, nil, f1)
	assertEqual(t, IsoFieldNotFoundError, err)
}

func Test_IsoMsg_InvalidIndex(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0100")
	f1, err := isoMsg.Get(129)
	assertEqual(t, OutOfBoundIndexError, err)
	assertEqual(t, nil, f1)
}

func Test_IsoMsg_MTI(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0100")
	mti, err := isoMsg.Get(0)
	assertEqual(t, nil, err)
	assertEqual(t, 0, mti.pos)
	assertEqual(t, "0100", mti.text)
}

func Test_IsoMsg_AsciiBitmap(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0500")
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
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0100")
	isoMsg.Set(2, "1234567890")
	isoMsg.Set(12, "1234")
	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, []byte("00000000000000004010000000000000"), bitmap.value)
}

func Test_IsoMsg_AsciiBitmapSecondary(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0100")
	isoMsg.Set(2, "1234567890")
	isoMsg.Set(12, "1234")
	isoMsg.Set(128, "2D2A98F12D2A98F1")

	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, []byte("C0100000000000000000000000000001"), bitmap.value)
}

func Test_IsoMsg_Print(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
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
	isoMsg := IsoMsgNew(ASCII1987)
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

	_, err := isoMsg.Bytes()
	assertEqual(t, nil, err)
	// fmt.Printf("%X", encoded)
}

func Test_IsoMsg_Bytes(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
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

	b, err := isoMsg.Bytes()
	assertEqual(t, nil, err)
	fmt.Printf("%X\n\n", b)
	fmt.Printf("%v", isoMsg)

}

func Test_IsoMsg_AsciiFields(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	isoMsg.Set(0, "0500")
	isoMsg.Set(2, "6734000000000000067")
	isoMsg.Set(12, "100100")
	isoMsg.Set(128, "2D2A98F12D2A98F1")

	mti, err := isoMsg.Get(0)
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

func Test_IsoMsg_Parse(t *testing.T) {
	isoMsg := IsoMsgNew(ASCII1987)
	err := isoMsg.ParseString("303231304632333832323031383841313832304130323030303030303030303030303031313936373334303030303030303030303030303637303030303030303030303030303030313031303830363038303130323832333537393130303130303038303635323830303130393637333030353030353131363733393031303032303030383036383537303030303043504F5339392020303635304D41455331313532353238303033303030303335343020202020202020202020202020203032364D414553303131353934313233343536303430372020202020203937383039385F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800303233303230303538383030303030303030303135202030303130323930323630303030303030303030383030353238333532364C42202020203536363632443241393846313244324139384631")
	assertEqual(t, nil, err)
	fmt.Printf("%v\n", isoMsg)

	mti, err := isoMsg.Get(0)
	assertEqual(t, nil, err)
	assertEqual(t, "0210", mti.text)

	bitmap, err := isoMsg.Get(1)
	assertEqual(t, nil, err)
	assertEqual(t, "F238220188A1820A0200000000000001", bitmap.text)

	pan, err := isoMsg.Get(2)
	assertEqual(t, nil, err)
	assertEqual(t, "6734000000000000067", pan.text)

	f3, err := isoMsg.Get(3)
	assertEqual(t, nil, err)
	assertEqual(t, "000000", f3.text)

	f4, err := isoMsg.Get(4)
	assertEqual(t, nil, err)
	assertEqual(t, "000000000101", f4.text)

	f7, err := isoMsg.Get(7)
	assertEqual(t, nil, err)
	assertEqual(t, "0806080102", f7.text)

	f11, err := isoMsg.Get(11)
	assertEqual(t, nil, err)
	assertEqual(t, "823579", f11.text)

	f12, err := isoMsg.Get(12)
	assertEqual(t, nil, err)
	assertEqual(t, "100100", f12.text)

	f13, err := isoMsg.Get(13)
	assertEqual(t, nil, err)
	assertEqual(t, "0806", f13.text)

	f19, err := isoMsg.Get(19)
	assertEqual(t, nil, err)
	assertEqual(t, "528", f19.text)

	f23, err := isoMsg.Get(23)
	assertEqual(t, nil, err)
	assertEqual(t, "001", f23.text)

	f32, err := isoMsg.Get(32)
	assertEqual(t, nil, err)
	assertEqual(t, "673005005", f32.text)

	f33, err := isoMsg.Get(33)
	assertEqual(t, nil, err)
	assertEqual(t, "67390100200", f33.text)

	f37, err := isoMsg.Get(37)
	assertEqual(t, nil, err)
	assertEqual(t, "080685700000", f37.text)

	f41, err := isoMsg.Get(41)
	assertEqual(t, nil, err)
	assertEqual(t, "CPOS99  ", f41.text)

	f43, err := isoMsg.Get(43)
	assertEqual(t, nil, err)
	assertEqual(t, "0650MAES115252800300003540              ", f43.text)

	f48, err := isoMsg.Get(48)
	assertEqual(t, nil, err)
	assertEqual(t, "MAES0115941234560407      ", f48.text)

	f49, err := isoMsg.Get(49)
	assertEqual(t, nil, err)
	assertEqual(t, "978", f49.text)

	mac, err := isoMsg.Get(128)
	assertEqual(t, nil, err)
	assertEqual(t, "2D2A98F12D2A98F1", mac.text)

}
