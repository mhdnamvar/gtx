package main

import (
	"fmt"
	"testing"
)

func Test_Ascii87_AChar(t *testing.T) {
	for _, codec := range ASCII1987 {
		fmt.Printf("%-8s%-35s%-3d%-3t\n", 
			codec.GetName(), 
			codec.GetDescription(),
			codec.GetLength(),
			codec.GetPadding())
	}
	fmt.Println()
	value := "AB"
	expected := []byte("AB  ")
	mtiCodec := ASCII1987[0]
	actual, err := mtiCodec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
