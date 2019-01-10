package util

import "reflect"

// Equal return if actual equal expected
func Equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	return reflect.DeepEqual(expected, actual)
}
