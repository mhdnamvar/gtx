package crypto

import (
	"fmt"
	"testing"
)

func TestDeriveKey(t *testing.T) {
	mdk := "11223344556677888877665544332211"
	pan := "1122334455667788"
	psn := "01"

	fmt.Printf("MDK: %s\n", mdk)
	fmt.Printf("PAN: %s\n", pan)
	fmt.Printf("PSN: %s\n", psn)

	udk := DeriveKey(mdk, pan, psn)
	fmt.Printf("UDK: %X\n", udk)
}
