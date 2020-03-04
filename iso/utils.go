package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Hex2Bin ...
func Hex2Bin(s string) (binString string) {
	for i := 0; i < len(s); i += 2 {
		str := s[i : i+2]
		if i64, err := strconv.ParseInt(str, 16, 64); err != nil {
			fmt.Println(err)
		} else {
			binString = fmt.Sprintf("%s%08b", binString, i64)
		}
	}
	return
}

// Hex2Dec ...
func Hex2Dec(s string) (dec int) {
	if dec, err := strconv.ParseInt(s, 16, 64); err != nil {
		fmt.Println(err)
	} else {
		return int(dec)
	}
	return
}

// Bin2Hex ...
func Bin2Hex(s string) (hexString string) {
	for i := 0; i < len(s); i += 4 {
		str := s[i : i+4]
		if i64, err := strconv.ParseInt(str, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			hexString = fmt.Sprintf("%s%X", hexString, i64)
		}
	}
	return
}

// LeftPad ...
func LeftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}

// RightPad ...
func RightPad(s string, padStr string, pLen int) string {
	return s + strings.Repeat(padStr, pLen)
}

// RightPad2Len ...
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// LeftPad2Len ...
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// IntToBcd ...
func IntToBcd(i uint64) []byte {
	var bcd []byte
	for i > 0 {
		low := i % 10
		i /= 10
		hi := i % 10
		i /= 10
		var x []byte
		x = append(x, byte((hi&0xf)<<4)|byte(low&0xf))
		bcd = append(x, bcd[:]...)
	}
	return bcd
}

// BcdToInt ...
func BcdToInt(bcd []byte) uint64 {
	var i uint64
	for k := range bcd {
		r0 := bcd[k] & 0xf
		r1 := bcd[k] >> 4 & 0xf
		r := r1*10 + r0
		i = i*uint64(100) + uint64(r)
	}
	return i
}

// StrToBcd ...
func StrToBcd(s string) []byte {
	return BytesToBcd([]byte(s))
}

// BytesToBcd ...
func BytesToBcd(b []byte) []byte {
	if len(b)%2 != 0 {
		//b = append(b, []byte{0x0}...)
		b = append([]byte{0x0}, b...)
	}
	var slice = make([]byte, len(b)/2)
	for i := 0; i < len(b); i++ {
		step := 4
		if (i & 1) == 1 {
			step = 0
		}
		slice[i>>1] = slice[i>>1] | (b[i]-48)<<uint(step)
	}
	return slice
}

func Btoi(b []byte) (int, error) {
	s := string(b)
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Btoi() failed: %v", err)
		return 0, err
	}
	return n, nil
}
