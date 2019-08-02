package main

import (
	"testing"
)

func Test_Bitmap_Encode_Primary(t *testing.T) {
	var bits Bitmap
	bits.Set(2)
	bits.Set(12)
	expected := "4010000000000000"	
	actual := bits.String()
	assertEqual(t, expected, actual)
}

func Test_Bitmap_Encode_SetSecondary(t *testing.T) {
	var bits Bitmap
	bits.Set(2)
	bits.Set(12)
	bits.Set(128)
	expected := "C0100000000000000000000000000001"	
	actual := bits.String()
	assertEqual(t, expected, actual)
}

func Test_Bitmap_Encode(t *testing.T) {
	var bits Bitmap
	bits.Sets(2, 3, 23, 36, 64, 65, 90, 128)
	expected := "E0000200100000018000004000000001"	
	actual := bits.String()
	assertEqual(t, expected, actual)
}

func Test_Bitmap_DecodeWithSecondary(t *testing.T) {
	var bits Bitmap
	bits.Sets(2, 3, 23, 36, 64, 65, 90, 128)
	expected := []int{1, 2, 3, 23, 36, 64, 65, 90, 128}
	actual := bits.Array()
	assertEqual(t, expected, actual)
}

func Test_Bitmap_DecodeOnlyPrimary(t *testing.T) {
	var bits Bitmap
	bits.Sets(2, 3, 23, 36, 64)
	expected := []int{2, 3, 23, 36, 64}
	actual := bits.Array()
	assertEqual(t, expected, actual)
}

func Test_Bitmap_Array(t *testing.T) {	
	var bits Bitmap
	value := "767F4601A8E1A20A"
	expected := []int{2, 3, 4, 6, 7, 10, 11, 12, 13, 14, 15, 16, 18,
		22, 23, 32, 33, 35, 37, 41, 42, 43, 48, 49, 51, 55, 61, 63}
	bits.Decode(value)	
	actual := bits.Array()
	assertEqual(t, expected, actual)
}

func Test_Bitmap_Empty(t *testing.T) {
	var bits Bitmap
	bits.Set(2)
	bits.Clear(2)
	expected := "0000000000000000"	
	actual := bits.String()
	assertEqual(t, expected, actual)
}

func Test_Bitmap_Secondary(t *testing.T) {
	var bits Bitmap
	bits.Set(2)
	bits.Set(12)
	bits.Set(128)
	expected := "C0100000000000000000000000000001"	
	actual := bits.String()
	assertEqual(t, expected, actual)
}
