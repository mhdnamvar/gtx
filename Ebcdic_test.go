package main

import (
	"testing"
)

func Test_EncodeEbcdic(t *testing.T) {
	expected := []byte{0x2B, 0x2F, 0x5F, 0xCE, 0x2F, 0xCA}
	actual := ASCIIToEbcdic("Namvar")
	checkEncodeResult(t, expected, actual, nil)
}

func Test_EncodeAscii(t *testing.T) {
	value := string([]byte{0x2B, 0x2F, 0x5F, 0xCE, 0x2F, 0xCA})
	expected := "Namvar"
	actual := string(EbcdicToASCII(value))
	checkDecodeResult(t, expected, actual, nil)
}
