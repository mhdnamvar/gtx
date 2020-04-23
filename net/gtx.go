package net
//
//import (
//	"fmt"
//	"os"
//	"log"
//)
//
//func usage() {
//	fmt.Printf("%s\n", "usage: gtx [-c <config-file>]")
//	fmt.Printf("%s\n", "options:")
//	fmt.Printf("%5s\t%-50s\n", "-c", "specify configuration file in json format")
//	fmt.Printf("%5s\t%-50s\n", "-h", "print help")
//}
//
//func NotMain() {
//
//	configFileName := "gtx.json"
//	if len(os.Args) == 2 {
//		configFileName = os.Args[1]
//	}
//
//	_, err := os.Open(configFileName)
//	if err != nil {
//		log.Fatal("Please specify the configuration file")
//		usage()
//	}else {
//		ReadConfig(configFileName)
//		log.Printf("Configuration: %s", configFileName)
//		channelType := GtxConf.Channel.Type
//
//		var isoServer IsoServer
//		isoServer.New(
//			GtxConf.Channel.Host,
//			GtxConf.Channel.Port,
//			channelType,
//			GtxConf.Channel.HeaderLen,
//			GtxConf.Channel.HeaderLenIncluded)
//
//		isoServer.Start()
//
//		//var isoClient IsoClient
//		//isoClient.New(
//		//	GtxConf.Channel.Host,
//		//	GtxConf.Channel.Port,
//		//	channelType,
//		//	GtxConf.Channel.HeaderLen,
//		//	GtxConf.Channel.HeaderLenIncluded,
//		//	GtxConf.IsoFields)
//		//
//		//isoClient.SendAndReceive()
//
//	}
//}
