package utils

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

var ebcdic2ascii = []byte{
	'\x00', '\x01', '\x02', '\x03', '\xdc', '\x09', '\xc3', '\x7f',
	'\xca', '\xb2', '\xd5', '\x0b', '\x0c', '\x0d', '\x0e', '\x0f',
	'\x10', '\x11', '\x12', '\x13', '\xdb', '\xda', '\x08', '\xc1',
	'\x18', '\x19', '\xc8', '\xf2', '\x1c', '\x1d', '\x1e', '\x1f',
	'\xc4', '\xb3', '\xc0', '\xd9', '\xbf', '\x0a', '\x17', '\x1b',
	'\xb4', '\xc2', '\xc5', '\xb0', '\xb1', '\x05', '\x06', '\x07',
	'\xcd', '\xba', '\x16', '\xbc', '\xbb', '\xc9', '\xcc', '\x04',
	'\xb9', '\xcb', '\xce', '\xdf', '\x14', '\x15', '\xfe', '\x1a',
	' ', '\xa0', '\xe2', '\xe4', '\xe0', '\xe1', '\xe3', '\xe5',
	'\xe7', '\xf1', '\xa2', '.', '<', '(', '+', '|',
	'&', '\xe9', '\xea', '\xeb', '\xe8', '\xed', '\xee', '\xef',
	'\xec', '\xdf', '!', '$', '*', ')', ';', '\xac',
	'-', '/', '\xc2', '\xc4', '\xc0', '\xc1', '\xc3', '\xc5',
	'\xc7', '\xd1', '\xa6', ',', '%', '_', '>', '?',
	'\xf8', '\xc9', '\xca', '\xcb', '\xc8', '\xcd', '\xce', '\xcf',
	'\xcc', '`', ':', '#', '@', '\'', '=', '"',
	'\xd8', 'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', '\xab', '\xbb', '\xf0', '\xfd', '\xfe', '\xb1',
	'\xb0', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
	'q', 'r', '\xaa', '\xba', '\xe6', '\xb8', '\xc6', '\xa4',
	'\xb5', '~', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z', '\xa1', '\xbf', '\xd0', '\xdd', '\xde', '\xae',
	'^', '\xa3', '\xa5', '\xb7', '\xa9', '\xa7', '\xb6', '\xbc',
	'\xbd', '\xbe', '[', ']', '\xaf', '\xa8', '\xb4', '\xd7',
	'{', 'A', 'B', 'C', 'D', 'E', 'F', 'G',
	'H', 'I', '\xad', '\xf4', '\xf6', '\xf2', '\xf3', '\xf5',
	'}', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', '\xb9', '\xfb', '\xfc', '\xf9', '\xfa', '\xff',
	'\\', '\xf7', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z', '\xb2', '\xd4', '\xd6', '\xd2', '\xd3', '\xd5',
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', '\xb3', '\xdb', '\xdc', '\xd9', '\xda', '\x1a',
}

var ascii2ebcdic = []byte{
	'\x00', '\x01', '\x02', '\x03', '7', '-', '.', '/',
	'\x16', '\x05', '%', '\x0b', '\x0c', '\x0d', '\x0e', '\x0f',
	'\x10', '\x11', '\x12', '\x13', '<', '=', '2', '&',
	'\x18', '\x19', '?', '\'', '\x1c', '\x1d', '\x1e', '\x1f',
	'@', 'Z', '\x7f', '{', '[', 'l', 'P', '}',
	'M', ']', '\\', 'N', 'k', '`', 'K', 'a',
	'\xf0', '\xf1', '\xf2', '\xf3', '\xf4', '\xf5', '\xf6', '\xf7',
	'\xf8', '\xf9', 'z', '^', 'L', '~', 'n', 'o',
	'|', '\xc1', '\xc2', '\xc3', '\xc4', '\xc5', '\xc6', '\xc7',
	'\xc8', '\xc9', '\xd1', '\xd2', '\xd3', '\xd4', '\xd5', '\xd6',
	'\xd7', '\xd8', '\xd9', '\xe2', '\xe3', '\xe4', '\xe5', '\xe6',
	'\xe7', '\xe8', '\xe9', '\xba', '\xe0', '\xbb', '\xb0', 'm',
	'y', '\x81', '\x82', '\x83', '\x84', '\x85', '\x86', '\x87',
	'\x88', '\x89', '\x91', '\x92', '\x93', '\x94', '\x95', '\x96',
	'\x97', '\x98', '\x99', '\xa2', '\xa3', '\xa4', '\xa5', '\xa6',
	'\xa7', '\xa8', '\xa9', '\xc0', 'O', '\xd0', '\xa1', '\x07',
	'?', '?', '?', '?', '?', '?', '?', '?',
	'?', '?', '?', '?', '?', '?', '?', '?',
	'?', '?', '?', '?', '?', '?', '?', '?',
	'?', '?', '?', '?', '?', '?', '?', '?',
	'A', '\xaa', 'J', '\xb1', '\x9f', '\xb2', 'j', '\xb5',
	'\xbd', '\xb4', '\x9a', '\x8a', '_', '\xca', '\xaf', '\xbc',
	'\x90', '\x8f', '\xea', '\xfa', '\xbe', '\xa0', '\xb6', '\xb3',
	'\x9d', '\xda', '\x9b', '\x8b', '\xb7', '\xb8', '\xb9', '\xab',
	'd', 'e', 'b', 'f', 'c', 'g', '\x9e', 'h',
	't', 'q', 'r', 's', 'x', 'u', 'v', 'w',
	'\xac', 'i', '\xed', '\xee', '\xeb', '\xef', '\xec', '\xbf',
	'\x80', '\xfd', '\xfe', '\xfb', '\xfc', '\xad', '\xae', 'Y',
	'D', 'E', 'B', 'F', 'C', 'G', '\x9c', 'H',
	'T', 'Q', 'R', 'S', 'X', 'U', 'V', 'W',
	'\x8c', 'I', '\xcd', '\xce', '\xcb', '\xcf', '\xcc', '\xe1',
	'p', '\xdd', '\xde', '\xdb', '\xdc', '\x8d', '\x8e', '\xdf',
}

func AsciiToEbcdic(s string) []byte {
	b := []byte(s)
	var ebcdic = make([]byte, len(b))
	for i, v := range b {
		ebcdic[i] = ascii2ebcdic[v]
	}
	return ebcdic
}

func EbcdicToAscii(s string) []byte {
	b := []byte(s)
	var ascii = make([]byte, len(b))
	for i, v := range b {
		ascii[i] = ebcdic2ascii[v]
	}
	return ascii
}

func EbcdicToAsciiBytes(b []byte) []byte {
	var ascii = make([]byte, len(b))
	for i, v := range b {
		ascii[i] = ebcdic2ascii[v]
	}
	return ascii
}
