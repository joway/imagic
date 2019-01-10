package util

import (
	"reflect"
	"strings"
)

// Contains check if element in target
func Contains(in interface{}, elem interface{}) bool {
	inValue := reflect.ValueOf(in)
	elemValue := reflect.ValueOf(elem)
	inType := inValue.Type()

	switch inType.Kind() {
	case reflect.String:
		return strings.Contains(inValue.String(), elemValue.String())
	case reflect.Map:
		for _, key := range inValue.MapKeys() {
			if Equal(key.Interface(), elem) {
				return true
			}
		}
	case reflect.Slice:
		for i := 0; i < inValue.Len(); i++ {
			if Equal(inValue.Index(i).Interface(), elem) {
				return true
			}
		}
	}

	return false
}
