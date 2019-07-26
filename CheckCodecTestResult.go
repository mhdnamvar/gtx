package main

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %v(type=%T), expected: %v(type=%T)\n", actual, actual, expected, expected)
	}
}

func checkEncodeResult(t *testing.T, expected []byte, actual []byte, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %X, expected: %X\n", actual, expected)
	}
}

func checkDecodeResult(t *testing.T, expected string, actual string, err error) {
	checkEncodeResult(t, []byte(expected), []byte(actual), err)
}

func checkEncodeError(t *testing.T, actual []byte, err error, errType int) {
	if err == nil || !reflect.DeepEqual(err.Error(), Errors[errType].message) {
		t.Errorf("Should return error\n")
	}
	if actual != nil {
		t.Errorf("Nil expected but received: %X\n", actual)
	}
}

func checkDecodeError(t *testing.T, actual string, err error, errType int) {
	if err == nil || !reflect.DeepEqual(err.Error(), Errors[errType].message) {
		t.Errorf("Should return error\n")
	}
	if actual != "" {
		t.Errorf("Empty string expected but received: %s\n", actual)
	}
}
