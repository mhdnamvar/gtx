package net
//
//import (
//	"net"
//	"fmt"
//	"log"
//	"strconv"
//	"../encoding/iso8583"
//	"../encoding/tlv"
//)
//
//type IsoClient struct {
//	conn net.Conn
//	host	string
//	port	int
//	channelType	int
//	headerLen	int
//	headerLenIncluded	bool
//	encoder []iso8583.IsoEncoder
//	isoFields map[string]interface{}
//}
//
//func (isoClient *IsoClient) New(host string, port int, channelType int,
//								headerLen int,
//								headerLenIncluded bool,
//								isoFields map[string]interface{}) {
//	isoClient.host = host
//	isoClient.port = port
//	isoClient.channelType = channelType
//	isoClient.headerLen = headerLen
//	isoClient.headerLenIncluded = headerLenIncluded
//	encoder, err := iso8583.GetEncoder(channelType)
//	checkError(err)
//	isoClient.encoder = encoder
//	isoClient.isoFields = isoFields
//	encoderName, err := iso8583.GetEncoderName(channelType)
//	checkError(err)
//	log.Printf("Channel type is: %s", encoderName)
//}
//
//func (isoClient *IsoClient) SetEncoder(encoder []iso8583.IsoEncoder) {
//	isoClient.encoder = encoder
//}
//
//func (isoClient *IsoClient) SetHost(host string) {
//	isoClient.host = host
//}
//
//func (isoClient *IsoClient) SetPort(port int) {
//	isoClient.port = port
//}
//
//func (isoClient *IsoClient) SetChannelType(channelType int) {
//	isoClient.channelType = channelType
//}
//
//func (isoClient *IsoClient) SetHeaderLen(headerLen int) {
//	isoClient.headerLen = headerLen
//}
//
//func (isoClient *IsoClient) SetHeaderLenIncluded(headerLenIncluded bool) {
//	isoClient.headerLenIncluded = headerLenIncluded
//}
//
//func (isoClient *IsoClient) SendAndReceive() {
//	isoClient.connect()
//	req, err  := isoClient.getRequest()
//	checkError(err)
//	buf, err := iso8583.Encode(&req, isoClient.encoder)
//	checkError(err)
//	WriteMessage(isoClient.conn, buf)
//	req.Log()
//	req.HexDump()
//
//	data, err := ReadMessage(isoClient.conn)
//	checkError(err)
//	res := iso8583.Decode(data, isoClient.encoder)
//	res.HexDump()
//	res.Log()
//
//	isoClient.conn.Close()
//}
//
//func (isoClient *IsoClient) connect() {
//	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", isoClient.host, isoClient.port))
//	if err != nil {
//		log.Fatal(err)
//	}else{
//		isoClient.conn = conn
//		log.Printf("Connected to %s:%d\n", isoClient.host, isoClient.port)
//	}
//}
//
//func (isoClient *IsoClient) getRequest() (iso8583.IsoMsg, error) {
//	var isoMsg  iso8583.IsoMsg
//	isoMsg.IsoEncoder = isoClient.encoder
//	for i := range isoClient.isoFields {
//		index , _ := strconv.Atoi(i)
//		//TODO: Need to determine binary fields!
//		if index == 52 || index == 64 || index ==128 {
//			isoMsg.SetBinary(index, isoClient.isoFields[i].(string))
//		} else if index == 55 {
//			DE55 := isoClient.isoFields[i].(map[string]interface{})
//			tlvParser := tlv.NewTlvParser()
//			data := ""
//			//TODO: How to encode nested TLV data?!
//			for k, v := range DE55 {
//				data += tlvParser.Encode(k, v.(string))
//			}
//			isoMsg.SetBinary(index, data)
//		} else {
//			isoMsg.Set(index, isoClient.isoFields[i].(string))
//		}
//	}
//	return isoMsg, nil
//}