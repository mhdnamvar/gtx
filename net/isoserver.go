package net
//
//import (
//	"encoding/binary"
//	"fmt"
//	"log"
//	"net"
//)
//
//type IsoServer struct {
//	conn              net.Conn
//	host              string
//	port              int
//	channelType       int
//	headerLen         int
//	headerLenIncluded bool
//}
//
//func (isoServer *IsoServer) SetHost(str string) {
//	isoServer.host = str
//}
//
//func (isoServer *IsoServer) SetPort(i int) {
//	isoServer.port = i
//}
//
//func (isoServer *IsoServer) SetChannelType(i int) {
//	isoServer.channelType = i
//}
//
//func (isoServer *IsoServer) SetHeaderLen(i int) {
//	isoServer.headerLen = i
//}
//
//func (isoServer *IsoServer) SetHeaderLenIncluded(b bool) {
//	isoServer.headerLenIncluded = b
//}
//
//func (isoServer *IsoServer) New(host string, port int, channelType int, headerLen int, headerLenIncluded bool) {
//	isoServer.SetHost(host)
//	isoServer.SetPort(port)
//	isoServer.SetChannelType(channelType)
//	isoServer.SetHeaderLen(headerLen)
//	isoServer.SetHeaderLenIncluded(headerLenIncluded)
//}
//
//func (isoServer *IsoServer) Start() {
//	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", isoServer.host, isoServer.port))
//	if err != nil {
//		log.Fatalf("%s", err.Error())
//	}
//	defer l.Close()
//	channelName, err := iso8583.GetEncoderName(isoServer.channelType)
//	if err != nil {
//		log.Fatalf("%s", err.Error())
//	}
//	log.Printf("Channel: %s", channelName)
//	log.Printf("Listening on %s:%d", isoServer.host, isoServer.port)
//	for {
//		isoServer.conn, err = l.Accept()
//		if err != nil {
//			log.Fatalf("%s", err.Error())
//		}
//		go isoServer.HandleRequest()
//	}
//}
//
//func (isoServer *IsoServer) HandleRequest() {
//	buf, err := ReadMessage(isoServer.conn)
//	if err != nil {
//		log.Fatalf("%s", err.Error())
//	}
//	if len(buf) == 0 {
//		log.Printf("Invalid message received!")
//		return
//	}
//	log.Printf("%x", buf)
//	encoder, err := iso8583.GetEncoder(isoServer.channelType)
//	if err != nil {
//		log.Fatalf("%s", err.Error())
//	}
//	res := iso8583.Decode(buf, encoder)
//	res.HexDump()
//	res.Log()
//	res.SetResponseMti()
//	res.Set(39, "00")
//
//	buf, err = iso8583.Encode(&res, encoder)
//	checkError(err)
//	WriteMessage(isoServer.conn, buf)
//	res.Log()
//	res.HexDump()
//
//}
//
//func (isoServer *IsoServer) Stop() {
//	log.Println("Not implemented yet!")
//}
//
//func (isoServer *IsoServer) ReadMessage() []byte {
//	header := make([]byte, isoServer.headerLen)
//	isoServer.conn.Read(header)
//	//log.Printf("header=%x", header)
//	var data []byte
//
//	switch isoServer.headerLen {
//	case 2:
//		if isoServer.headerLenIncluded {
//			data = make([]byte, binary.BigEndian.Uint16(header)-2)
//		} else {
//			data = make([]byte, binary.BigEndian.Uint16(header))
//		}
//		break
//	case 4:
//		if isoServer.headerLenIncluded {
//			data = make([]byte, binary.BigEndian.Uint32(header)-4)
//		} else {
//			data = make([]byte, binary.BigEndian.Uint32(header))
//		}
//		break
//	case 8:
//		if isoServer.headerLenIncluded {
//			data = make([]byte, binary.BigEndian.Uint64(header)-8)
//		} else {
//			data = make([]byte, binary.BigEndian.Uint64(header))
//		}
//		break
//	default:
//		e := fmt.Errorf("Invalid header length: %d", isoServer.headerLen)
//		log.Fatal(e.Error())
//	}
//	_, err := isoServer.conn.Read(data)
//	checkError(err)
//
//	return data
//}
//
//func (isoServer *IsoServer) WriteMessage(msg []byte) {
//	header := make([]byte, isoServer.headerLen)
//	msgLen := len(msg)
//	if isoServer.headerLenIncluded {
//		msgLen += isoServer.headerLen
//	}
//
//	switch isoServer.headerLen {
//	case 2:
//		binary.BigEndian.PutUint16(header, uint16(msgLen))
//		break
//	case 4:
//		binary.BigEndian.PutUint32(header, uint32(msgLen))
//		break
//	case 8:
//		binary.BigEndian.PutUint64(header, uint64(msgLen))
//		break
//	default:
//		log.Fatalf("Invalid header header: %d", isoServer.headerLen)
//	}
//
//	isoServer.conn.Write(header)
//	isoServer.conn.Write(msg)
//}
