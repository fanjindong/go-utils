package utils

import (
	"reflect"
)

func SliceInsert(s interface{}, index int, item interface{}) error {
	t := reflect.TypeOf(s).Elem()
	switch t.Kind() {
	case reflect.Slice:
	default:
		return InvalidSlice
	}

	v := reflect.ValueOf(s).Elem()
	if index >= v.Len() || index < 0 {
		return SliceIndexRangeOut
	}
	if !v.CanSet() {
		return SliceMustPointer
	}

	newSlice := reflect.MakeSlice(t, 0, v.Cap()+1)
	for i := 0; i < v.Len(); i++ {
		if i == index {
			newSlice = reflect.Append(newSlice, reflect.ValueOf(item))
		}
		newSlice = reflect.Append(newSlice, v.Index(i))
	}
	v.Set(newSlice)
	return nil
}
