package net

import (
	"log"
	"fmt"
	"encoding/binary"
	"net"
)



func ReadMessage(conn net.Conn) ([]byte, error) {

	header := make([]byte, GtxConf.Channel.HeaderLen)
	conn.Read(header)
	log.Printf("header=%x", header)
	var data []byte

	switch GtxConf.Channel.HeaderLen {
		case 2:
			if GtxConf.Channel.HeaderLenIncluded {
				data = make([]byte, binary.BigEndian.Uint16(header)-2)
			} else {
				data = make([]byte, binary.BigEndian.Uint16(header))
			}
			break
		case 4:
			if GtxConf.Channel.HeaderLenIncluded {
				data = make([]byte, binary.BigEndian.Uint32(header)-4)
			} else {
				data = make([]byte, binary.BigEndian.Uint32(header))
			}
			break
		case 8:
			if GtxConf.Channel.HeaderLenIncluded {
				data = make([]byte, binary.BigEndian.Uint64(header)-8)
			} else {
				data = make([]byte, binary.BigEndian.Uint64(header))
			}
			break
		default:
			e := fmt.Errorf("invalid header length: %d", GtxConf.Channel.HeaderLen)
			//log.Fatal(e.Error())
			return nil, e
	}
	_, err := conn.Read(data)
	log.Printf("data=%X", data)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	return data, nil
}

func WriteMessage(conn net.Conn, msg []byte) {
	header := make([]byte, GtxConf.Channel.HeaderLen)
	msgLen := len(msg)
	if GtxConf.Channel.HeaderLenIncluded {
		msgLen += GtxConf.Channel.HeaderLen
	}

	switch GtxConf.Channel.HeaderLen {
		case 2:
			binary.BigEndian.PutUint16(header, uint16(msgLen))
			break
		case 4:
			binary.BigEndian.PutUint32(header, uint32(msgLen))
			break
		case 8:
			binary.BigEndian.PutUint64(header, uint64(msgLen))
			break
		default:
			log.Fatalf("Invalid header header: %d", GtxConf.Channel.HeaderLen)
		}

	conn.Write(header)
	conn.Write(msg)
}