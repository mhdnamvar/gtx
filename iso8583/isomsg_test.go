package iso8583

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsoMsgGet(t *testing.T) {
	isoMsg := IsoMsgNew()
	isoMsg.Set(0, "320")
	isoMsg.Set(128, "2D2A98F12D2A98F1")
	isoMsg.Set(129, "1234")

	s, err := isoMsg.Get(0)
	assert.Equal(t, "320", s)
	assert.Equal(t, nil, err)

	s, err = isoMsg.Get(128)
	assert.Equal(t, "2D2A98F12D2A98F1", s)
	assert.Equal(t, nil, err)

	s, err = isoMsg.Get(129)
	assert.Equal(t, "", s)
	assert.Equal(t, IsoFieldNotFoundError, err)

	s, err = isoMsg.Get(-1)
	assert.Equal(t, "", s)
	assert.Equal(t, IsoFieldNotFoundError, err)
}

func TestIsoMsgEncodeAscii(t *testing.T) {
	isoMsg := sampleIsoMsg()
	bytes, err := isoMsg.Encode(Ascii87)
	assert.Equal(t, err, nil)
	log.Printf("%X", bytes)
}

func TestIsoMsgParse(t *testing.T) {
	isoMsg := IsoMsgNew()
	s := "30323130" +
		"4632333832323031383841313832314130323030303030303030303030303031" +
		"3139" + "36373334303030303030303030303030303637" +
		"303030303030" +
		"303030303030303030313031" +
		"30383036303830313032" +
		"383233353739" +
		"313030313030" +
		"30383036" +
		"353238" +
		"303031" +
		"3039" + "363733303035303035" + // 32
		"3131" + "3637333930313030323030" + // 33
		"303830363835373030303030" +
		"43504F5339392020" + // 41
		"303635304d4145533131353235323830303330303030333534302020202020202020202020202020" + // 43
		"303236" + "4d41455330313135393431323334353630343037202020202020" + // 48
		"393738" + // 49
		"303938" + "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800" +
		"303030" + "" + // 60
		"303233" + "3032303035383830303030303030303031352020303031" + // 61
		"303239" + "30323630303030303030303030383030353238333532364c4220202020" + // 63
		"35363636" + // 71
		"32443241393846313244324139384631" // 128

	err := isoMsg.Parse(Ascii87, s)
	assert.Equal(t, err, nil)

	mti, err := isoMsg.Get(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0210", mti)

	bitmap, err := isoMsg.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "F238220188A1821A0200000000000001", bitmap)

	pan, err := isoMsg.Get(2)
	assert.Equal(t, nil, err)
	assert.Equal(t, "6734000000000000067", pan)

	f3, err := isoMsg.Get(3)
	assert.Equal(t, nil, err)
	assert.Equal(t, "000000", f3)

	f4, err := isoMsg.Get(4)
	assert.Equal(t, nil, err)
	assert.Equal(t, "000000000101", f4)

	f7, err := isoMsg.Get(7)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0806080102", f7)

	f11, err := isoMsg.Get(11)
	assert.Equal(t, nil, err)
	assert.Equal(t, "823579", f11)

	f12, err := isoMsg.Get(12)
	assert.Equal(t, nil, err)
	assert.Equal(t, "100100", f12)

	f13, err := isoMsg.Get(13)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0806", f13)

	f19, err := isoMsg.Get(19)
	assert.Equal(t, nil, err)
	assert.Equal(t, "528", f19)

	f23, err := isoMsg.Get(23)
	assert.Equal(t, nil, err)
	assert.Equal(t, "001", f23)

	f32, err := isoMsg.Get(32)
	assert.Equal(t, nil, err)
	assert.Equal(t, "673005005", f32)

	f33, err := isoMsg.Get(33)
	assert.Equal(t, nil, err)
	assert.Equal(t, "67390100200", f33)

	f37, err := isoMsg.Get(37)
	assert.Equal(t, nil, err)
	assert.Equal(t, "080685700000", f37)

	f41, err := isoMsg.Get(41)
	assert.Equal(t, nil, err)
	assert.Equal(t, "CPOS99  ", f41)

	f43, err := isoMsg.Get(43)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0650MAES115252800300003540              ", f43)

	f48, err := isoMsg.Get(48)
	assert.Equal(t, nil, err)
	assert.Equal(t, "MAES0115941234560407      ", f48)

	f49, err := isoMsg.Get(49)
	assert.Equal(t, nil, err)
	assert.Equal(t, "978", f49)

	f55, err := isoMsg.Get(55)
	assert.Equal(t, nil, err)
	assert.Equal(t, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F1"+
		"0201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF2"+
		"7A982025C009505F070AC9800", f55)

	f60, err := isoMsg.Get(60)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", f60)

	f61, err := isoMsg.Get(61)
	assert.Equal(t, nil, err)
	assert.Equal(t, "020058800000000015  001", f61)

	f63, err := isoMsg.Get(63)
	assert.Equal(t, nil, err)
	assert.Equal(t, "02600000000008005283526LB    ", f63)

	f71, err := isoMsg.Get(71)
	assert.Equal(t, nil, err)
	assert.Equal(t, "5666", f71)

	f128, err := isoMsg.Get(128)
	assert.Equal(t, nil, err)
	assert.Equal(t, "2D2A98F12D2A98F1", f128)

	log.Print(isoMsg)
}

func sampleIsoMsg() *IsoMsg {
	isoMsg := IsoMsgNew()
	isoMsg.Set(0, "210")
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
	isoMsg.Set(55, "5F2A") //0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800
	isoMsg.Set(60, "")
	isoMsg.Set(61, "020058800000000015  001")
	isoMsg.Set(63, "02600000000008005283526LB    ")
	isoMsg.Set(71, "5666")
	isoMsg.Set(128, "2D2A98F12D2A98F1")
	return isoMsg
}

func TestIsoMsgParseAscii87(t *testing.T) {
	s := "303231304632333832323031383841313832314130323030303030303030303030303031313936373334303030303030303030303030303637303030303030303030303030303030313031303830363038303130323832333537393130303130303038303635323830303130393637333030353030353131363733393031303032303030383036383537303030303043504F5339392020303635304D41455331313532353238303033303030303335343020202020202020202020202020203032364D414553303131353934313233343536303430372020202020203937383039385F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800303030303233303230303538383030303030303030303135202030303130323930323630303030303030303030383030353238333532364C42202020203536363632443241393846313244324139384631"
	isoMsg := IsoMsgNew()
	err := isoMsg.Parse(Ascii87, s)
	assert.Equal(t, nil, err)
}

func TestIsoMsgDump(t *testing.T) {
	s, err := sampleIsoMsg().Dump(Ascii87)
	assert.Equal(t, nil, err)
	fmt.Println(s)
}

func TestIsoMsgEncodeBinary(t *testing.T) {
	isoMsg := IsoMsgNew()
	isoMsg.Set(0, "0210")
	isoMsg.Set(2, "6734000000000000067") //TODO: This field has encoding issue - IFB_FLLNUM
	isoMsg.Set(3, "000000")
	isoMsg.Set(28, "1234")      //TODO: This field has encoding issue - IFB_AMOUNT
	isoMsg.Set(32, "673005005") //TODO: This field has encoding issue - IFB_LLNUM
	isoMsg.Set(34, "1234ABCD")
	isoMsg.Set(36, "1234ABCD")
	isoMsg.Set(37, "0806857")
	isoMsg.Set(47, "1234ABCD")
	isoMsg.Set(128, "2D2A98F12D2A98F1")
	bytes, err := isoMsg.Encode(Binary87)
	assert.Equal(t, err, nil)
	log.Printf("%X", bytes)
}
