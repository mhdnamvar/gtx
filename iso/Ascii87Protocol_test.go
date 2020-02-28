package main

import (
	"testing"
)

func Test_Ascii87_AChar(t *testing.T) {
	value := "AB"
	expected := []byte("AB          ")
	codec := ASCII1987[37] //RRN
	actual, err := codec.Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
