package iso

import (
	"reflect"
	"testing"
)

/*
	SHOULD BE REPLACED WITH https://github.com/stretchr/testify
*/

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected == nil {
		assertNil(t, actual)
	} else if actual == nil {
		assertNil(t, expected)
	} else if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual: %v(type=%T), expected: %v(type=%T)\n", actual, actual, expected, expected)
	}
}

func assertNil(t *testing.T, i interface{}) {
	if i != nil && (reflect.ValueOf(i).Kind() != reflect.Ptr || !reflect.ValueOf(i).IsNil()) {
		switch i.(type) {
		case []byte:
			if len(i.([]byte)) > 0 {
				t.Errorf("Expected nil but found: %v(type=%T)\n", i, i)
			}
		case string:
			if len(i.(string)) > 0 {
				t.Errorf("Expected nil but found: %v(type=%T)\n", i, i)
			}
		}
	}
}

func assertEqualError(t *testing.T, expected interface{}, actual interface{}, err error) {
	if err != nil && (actual.([]byte) == nil || actual.(string) == "") {
		assertEqual(t, expected, actual)
	} else {
		t.Errorf("Expected error: %s", err.Error())
	}
}

func assertEqualErrorCode(t *testing.T, actual interface{}, err error, errType int) {
	if err != nil && (actual.([]byte) == nil || actual.(string) == "") {
		assertEqual(t, err.Error(), Errors[errType].Error())
	} else {
		t.Errorf("Expected error: %s", Errors[errType].Error())
	}
}
