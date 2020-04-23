package crypto

import (
	"fmt"
	"encoding/hex"
	"encoding/binary"
	"bytes"
	"crypto/des"
)

func DeriveKey(key, pan, psn string) []byte{
	var b bytes.Buffer
	s := pan + psn
	fmt.Sprintf(s, "%016s", s)
	d, _ := hex.DecodeString(s[len(s)-16:])
	b.Write(d)
	binary.BigEndian.PutUint64(d, binary.BigEndian.Uint64(d)^0xFFFFFFFFFFFFFFFF)
	b.Write(d)
	k, _ := hex.DecodeString(key)
	return DesEncrypt(b.Bytes(), k)
}

func DesEncrypt(d, k []byte) []byte {
	var tk bytes.Buffer
	tk.Write(k)
	tk.Write(k[:8])
	b, _ := des.NewTripleDESCipher(tk.Bytes())
	bs := b.BlockSize()
	o := make([]byte, len(d))
	t := o
	for len(d) > 0 {
		b.Encrypt(t, d[:bs])
		d = d[bs:]
		t = t[bs:]
	}
	return o
}
