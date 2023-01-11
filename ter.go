package ter

import (
	"reflect"
)

// var falsyList = []any{
// 	false,
// 	0,
// 	"",
// 	nil,
// }

func True[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func Truthy[T comparable](truthyValue any, trueValue T, falseValue T) T {
	ifTrue := !isFalsy(truthyValue)
	return True(ifTrue, trueValue, falseValue)
}

func isFalsy(item any) bool {

	// case nil
	if item == nil {
		return true
	}

	v := reflect.ValueOf(item)

	switch v.Kind() {

	// case false
	case reflect.Bool:
		return !v.Bool()

	// case 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	// case ""
	case reflect.String:
		str := v.String()
		return str == ""

	default:
		return false
	}
}
