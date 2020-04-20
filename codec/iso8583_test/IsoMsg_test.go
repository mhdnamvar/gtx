package iso8583_test

import (
	"encoding/hex"
	"log"
	"strings"
	"testing"

	. "../iso8583"
	"github.com/stretchr/testify/assert"
)

func sampleIsoMsg() *IsoMsg {
	isoMsg := NewIsoMsg()
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
	isoMsg.Set(34, "1234567890")
	isoMsg.Set(36, "1234567890")
	isoMsg.Set(37, "080685700000")
	isoMsg.Set(41, "CPOS99  ")
	isoMsg.Set(43, "0650MAES115252800300003540")
	isoMsg.Set(48, "MAES0115941234560407      ")
	isoMsg.Set(49, "978")
	isoMsg.Set(55, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800")
	isoMsg.Set(56, "ABCD1234")
	isoMsg.Set(60, "")
	isoMsg.Set(61, "020058800000000015  001")
	isoMsg.Set(63, "02600000000008005283526LB    ")
	isoMsg.Set(71, "5666")
	isoMsg.Set(128, "2D2A98F12D2A98F1")
	return isoMsg
}

func TestIsoMsgGet(t *testing.T) {
	isoMsg := NewIsoMsg()
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
	assert.Equal(t, FieldNotFound, err)

	s, err = isoMsg.Get(-1)
	assert.Equal(t, "", s)
	assert.Equal(t, FieldNotFound, err)
}

func TestIsoMsgEncodeAscii(t *testing.T) {
	isoMsg := sampleIsoMsg()
	expected := "30323130463233383232303144384131383331413032303030303030303030303030303131393637333430303030" +
		"30303030303030303036373030303030303030303030303030303130313038303630383031303238323335373931303031303" +
		"03038303635323830303130393637333030353030353131363733393031303032303031303132333435363738393030313031" +
		"32333435363738393030383036383537303030303043504F5339392020303635304D414553313135323532383030333030303" +
		"03335343020202020202020202020202020203032364D414553303131353934313233343536303430372020202020203937383" +
		"039385F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA00001" +
		"1223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800303" +
		"0384142434431323334303030303233303230303538383030303030303030303135202030303130323930323630303030303030" +
		"303030383030353238333532364C42202020203536363632443241393846313244324139384631"
	bytes, err := isoMsg.Encode(DefaultAscii87)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, strings.ToUpper(hex.EncodeToString(bytes)))
}

func TestIsoMsgParseDefaultAscii87(t *testing.T) {
	isoMsg := NewIsoMsg()
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
		"303236" + "4D41455330313135393431323334353630343037202020202020" + // 48
		"393738" + // 49
		"303938" + "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800" +
		"303030" + "" + // 60
		"303233" + "3032303035383830303030303030303031352020303031" + // 61
		"303239" + "30323630303030303030303030383030353238333532364c4220202020" + // 63
		"35363636" + // 71
		"32443241393846313244324139384631" // 128

	err := isoMsg.Parse(DefaultAscii87, s)
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

}

func TestIsoMsgParseAscii87(t *testing.T) {
	s := "30323130463233383232303144384131383331413032303030303030303030303030303131393637333430" +
		"303030303030303030303030363730303030303030303030303030303031303130383036303830313032383" +
		"233353739313030313030303830363532383030313039363733303035303035313136373339303130303230" +
		"303130313233343536373839303031303132333435363738393030383036383537303030303043504F533939" +
		"2020303635304D41455331313532353238303033303030303335343020202020202020202020202020203032" +
		"364D414553303131353934313233343536303430372020202020203937383039360209789A031409029C0100" +
		"9F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA0000112233440458560000" +
		"00000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800303038" +
		"4142434431323334303030303233303230303538383030303030303030303135202030303130323930323630" +
		"303030303030303030383030353238333532364C42202020203536363632443241393846313244324139384631"
	isoMsg := NewIsoMsg()
	err := isoMsg.Parse(DefaultAscii87, s)
	assert.Equal(t, nil, err)
}

func TestIsoMsgEncodeBinary87(t *testing.T) {
	isoMsg := sampleIsoMsg()
	expected := "0210F2382201D8A1831A0200000000000001196734000000000000067F000000000000000101080608010282357910010008060528000109673005005011673901002000103132333435363738393000103132333435363738393030383036383537303030303043504F5339392020303635304D414553313135323532383030333030303033353430202020202020202020202020202000264D41455330313135393431323334353630343037202020202020097800985F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC980000084142434431323334000000233032303035383830303030303030303031352020303031002930323630303030303030303030383030353238333532364C422020202056662D2A98F12D2A98F1"
	bytes, err := isoMsg.Encode(DefaultBinary87)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, strings.ToUpper(hex.EncodeToString(bytes)))
}

func TestIsoMsgEncodeEbcdic87(t *testing.T) {
	isoMsg := sampleIsoMsg()
	expected := "F0F2F1F0F2382201D8A1831A0200000000000001F1F9F6F7F3F4F0F0F0F0F0F0F0F0F0F0F0F0F0F6F7F0F0F0F" +
		"0F0F0F0F0F0F0F0F0F0F0F0F1F0F1F0F8F0F6F0F8F0F1F0F2F8F2F3F5F7F9F1F0F0F1F0F0F0F8F0F6F5F2F8F0F0F1F0F9F" +
		"6F7F3F0F0F5F0F0F5F1F1F6F7F3F9F0F1F0F0F2F0F0F1F0F1F2F3F4F5F6F7F8F9F0F0F1F0F1F2F3F4F5F6F7F8F9F0F0F8F" +
		"0F6F8F5F7F0F0F0F0F0C3D7D6E2F9F94040F0F6F5F0D4C1C5E2F1F1F5F2F5F2F8F0F0F3F0F0F0F0F3F5F4F04040404040" +
		"404040404040404040F0F2F6D4C1C5E2F0F1F1F5F9F4F1F2F3F4F5F6F0F4F0F7404040404040F9F7F8F0F9F60209789A0" +
		"31409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA0000112233440458560" +
		"00000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800F0F0F8C1C2C3C" +
		"4F1F2F3F4F0F0F0F0F2F3F0F2F0F0F5F8F8F0F0F0F0F0F0F0F0F0F1F54040F0F0F1F0F2F9F0F2F6F0F0F0F0F0F0F0F0F0F" +
		"0F8F0F0F5F2F8F3F5F2F6D3C240404040F5F6F6F6F2C4F2C1F9F8C6F1F2C4F2C1F9F8C6F1"
	bytes, err := isoMsg.Encode(DefaultEbcdic87)
	s := strings.ToUpper(hex.EncodeToString(bytes))
	log.Printf("Hex: %s", s)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, s)
}

func TestIsoMsgParseBinary87(t *testing.T) {
	isoMsg := NewIsoMsg()
	s := "0210F2382201D8A1831A0200000000000001196734000000000000067F00000000000000010108060801028235791001000806052800010" +
		"9673005005011673901002000103132333435363738393000103132333435363738393030383036383537303030303043504F53393920203" +
		"03635304D414553313135323532383030333030303033353430202020202020202020202020202000264D414553303131353934313233343" +
		"53630343037202020202020097800985F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F" +
		"4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC98" +
		"000008414243443132333400000023303230303538383030303030303030303135202030303100293032363030303030303030303038303035" +
		"3238333532364C422020202056662D2A98F12D2A98F1"
	err := isoMsg.Parse(DefaultBinary87, s)
	assert.Equal(t, err, nil)

	mti, err := isoMsg.Get(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0210", mti)

	bitmap, err := isoMsg.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "F2382201D8A1831A0200000000000001", bitmap)

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
	assert.Equal(t, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAA" +
		"A000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800", f55)

	f56, err := isoMsg.Get(56)
	assert.Equal(t, nil, err)
	assert.Equal(t, "ABCD1234", f56)

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

}

func TestIsoMsgParseEbcdic87(t *testing.T) {
	isoMsg := NewIsoMsg()
	s := "F0F2F1F0F2382201D8A1831A0200000000000001F1F9F6F7F3F4F0F0F0F0F0F0F0F0F0F0F0F0F0F6F7F0F0F0F" +
		"0F0F0F0F0F0F0F0F0F0F0F0F1F0F1F0F8F0F6F0F8F0F1F0F2F8F2F3F5F7F9F1F0F0F1F0F0F0F8F0F6F5F2F8F0F0F1F0F9F" +
		"6F7F3F0F0F5F0F0F5F1F1F6F7F3F9F0F1F0F0F2F0F0F1F0F1F2F3F4F5F6F7F8F9F0F0F1F0F1F2F3F4F5F6F7F8F9F0F0F8F" +
		"0F6F8F5F7F0F0F0F0F0C3D7D6E2F9F94040F0F6F5F0D4C1C5E2F1F1F5F2F5F2F8F0F0F3F0F0F0F0F3F5F4F04040404040" +
		"404040404040404040F0F2F6D4C1C5E2F0F1F1F5F9F4F1F2F3F4F5F6F0F4F0F7404040404040F9F7F8F0F9F60209789A0" +
		"31409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA0000112233440458560" +
		"00000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800F0F0F8C1C2C3C" +
		"4F1F2F3F4F0F0F0F0F2F3F0F2F0F0F5F8F8F0F0F0F0F0F0F0F0F0F1F54040F0F0F1F0F2F9F0F2F6F0F0F0F0F0F0F0F0F0F" +
		"0F8F0F0F5F2F8F3F5F2F6D3C240404040F5F6F6F6F2C4F2C1F9F8C6F1F2C4F2C1F9F8C6F1"

	err := isoMsg.Parse(DefaultEbcdic87, s)
	assert.Equal(t, err, nil)

	mti, err := isoMsg.Get(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, "0210", mti)

	bitmap, err := isoMsg.Get(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, "F2382201D8A1831A0200000000000001", bitmap)

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
	assert.Equal(t, "0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F1"+
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
}

func TestIsoMsgEncodeBinaryField32(t *testing.T) {
	isoMsg := NewIsoMsg()
	isoMsg.Set(0, "100")
	isoMsg.Set(32, "1234567890123456789")

	expected := "010000000001000000001912345678901234567890"

	b, err := isoMsg.Encode(DefaultBinary87)
	assert.Equal(t, err, nil)
	s := strings.ToUpper(hex.EncodeToString(b))
	log.Println(s)
	assert.Equal(t, expected, s)
}

func TestIsoMsgBinary87Field55(t *testing.T) {
	log.Println("----> Encoding...")
	isoMsg := NewIsoMsg()
	isoMsg.Set(0, "0210")
	isoMsg.Set(55, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800")
	expected := "0210000000000000020000985F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800"

	b, err := isoMsg.Encode(DefaultBinary87)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, strings.ToUpper(hex.EncodeToString(b)))

	log.Println("----> Decoding...")
	err = isoMsg.Decode(DefaultBinary87, b)
	assert.Equal(t, err, nil)

	f0, _ := isoMsg.Get(0)
	assert.Equal(t, "0210", f0)

	f1, _ := isoMsg.Get(1)
	assert.Equal(t, "0000000000000200", f1)

	f55, _ := isoMsg.Get(55)
	assert.Equal(t, "5F2A0209789A031409029C01009F1A0209789F02063030323030309F03060000000000009F10201F4301AAAAAAAA000011223344045856000000000000000000000000000000009F2701809F3602F1039F37044DDF27A982025C009505F070AC9800", f55)
}

func TestIsoMsgBinary87Field43(t *testing.T) {
	log.Println("----> Encoding...")
	isoMsg := NewIsoMsg()
	isoMsg.Set(0, "0210")
	isoMsg.Set(43, "0650MAES115252800300003540")
	expected := "02100000000000200000303635304D4145533131353235323830303330303030333534302020202020202020202020202020"

	b, err := isoMsg.Encode(DefaultBinary87)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, strings.ToUpper(hex.EncodeToString(b)))

	log.Println("----> Decoding...")
	err = isoMsg.Decode(DefaultBinary87, b)
	assert.Equal(t, err, nil)

	f0, _ := isoMsg.Get(0)
	assert.Equal(t, "0210", f0)

	f1, _ := isoMsg.Get(1)
	assert.Equal(t, "0000000000200000", f1)

	f43, _ := isoMsg.Get(43)
	assert.Equal(t, "0650MAES115252800300003540              ", f43)
}

func TestIsoMsgBinary87Field36(t *testing.T) {
	log.Println("----> Encoding...")
	isoMsg := NewIsoMsg()
	isoMsg.Set(0, "0210")
	isoMsg.Set(36, "1234567890")
	expected := "02100000000010000000001031323334353637383930"

	b, err := isoMsg.Encode(DefaultBinary87)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, strings.ToUpper(hex.EncodeToString(b)))

	log.Println("----> Decoding...")
	err = isoMsg.Decode(DefaultBinary87, b)
	assert.Equal(t, err, nil)

	f0, _ := isoMsg.Get(0)
	assert.Equal(t, "0210", f0)

	f1, _ := isoMsg.Get(1)
	assert.Equal(t, "0000000010000000", f1)

	f36, _ := isoMsg.Get(36)
	assert.Equal(t, "1234567890", f36)
}

func TestIsoMsgBinary87Field2(t *testing.T) {
	log.Println("----> Encoding...")
	isoMsg := NewIsoMsg()
	isoMsg.Set(0, "0210")
	isoMsg.Set(2, "6734000000000000067")
	expected := "02104000000000000000196734000000000000067F"

	b, err := isoMsg.Encode(DefaultBinary87)
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, strings.ToUpper(hex.EncodeToString(b)))

	log.Println("----> Decoding...")
	err = isoMsg.Decode(DefaultBinary87, b)
	assert.Equal(t, err, nil)

	f0, _ := isoMsg.Get(0)
	assert.Equal(t, "0210", f0)

	f1, _ := isoMsg.Get(1)
	assert.Equal(t, "4000000000000000", f1)

	f2, _ := isoMsg.Get(2)
	assert.Equal(t, "6734000000000000067", f2)
}
