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
	if index > v.Len() || index < 0 {
		return SliceIndexRangeOut
	}
	if !v.CanSet() {
		return SliceMustPointer
	}

	itemValue := reflect.ValueOf(item)
	newSlice := reflect.Append(v, itemValue)

	reflect.Copy(newSlice.Slice(index+1, newSlice.Len()), newSlice.Slice(index, newSlice.Len()))
	newSlice.Index(index).Set(itemValue)
	v.Set(newSlice)
	return nil
}
