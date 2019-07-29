package main

import (
	"fmt"
	"testing"
)

func Test_Ascii87_DE000(t *testing.T) {
	for _, c := range ASCII1987 {
		switch c.(type) {
		case *AChar:
			fmt.Printf("%-8s%-50s%-3d%-5t\n",
				c.(*AChar).Name,
				c.(*AChar).Description,
				c.(*AChar).Length,
				c.(*AChar).Padding)
		}
	}
	fmt.Println()
	value := "AB"
	expected := []byte("AB  ")
	actual, err := ASCII1987[0].Encode(value)
	assertEqual(t, nil, err)
	assertEqual(t, expected, actual)
}
