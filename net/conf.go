package net

import (
	"encoding/json"
	"log"
	"os"
)

var GtxConf GtxConfig

type GtxChannel struct {
	Host			string
	Port			int
	Type			int
	HeaderLen		int
	HeaderLenIncluded	bool
}

type GtxCryptogram struct {
	Enabled	bool
	Type	int
	Data	map[string]string

}
type GtxConfig struct {
	Name 		string
	Channel		GtxChannel
	IsoFields   map[string]interface{}
	Cryptogram	GtxCryptogram
}

func ReadConfig(confStr string) {
	file, err := os.Open(confStr)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	decoder := json.NewDecoder(file)
	GtxConf = GtxConfig{}
	err = decoder.Decode(&GtxConf)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}

