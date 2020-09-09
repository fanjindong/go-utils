package utils

import (
	"math/rand"
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

//func SliceElemsFormat(elems ...interface{}) []string {
//	ss := make([]string, 0, len(elems))
//	for i, elem := range elems {
//		switch e := elem.(type) {
//		case string:
//			ss[i] = e
//		case int:
//			ss[i] = strconv.FormatInt(int64(e), 10)
//		case int8:
//			ss[i] = strconv.FormatInt(int64(e), 10)
//		case int16:
//			ss[i] = strconv.FormatInt(int64(e), 10)
//		case int32:
//			ss[i] = strconv.FormatInt(int64(e), 10)
//		}
//	}
//	return ss
//}

//SliceShuffle 用于打乱一个切片
func SliceShuffle(slice interface{}) {
	v := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	for n := v.Len() - 1; n > 0; n-- {
		rIndex := rand.Intn(n)
		swap(rIndex, n)
	}
}

//SliceContain 判断一个数组(切片)是否包含某一元素(支持[]int, []string, []Struct, []*Struct等)
func SliceContain(slice interface{}, value interface{}) bool {
	return SliceContainsAny(slice, value)
}

//SliceContainsAny 判断一个数组(切片)是否包含任一元素(支持[]int, []string, []Struct, []*Struct等)
// value in slice: return true
// value not in slice: return false
// slice not in (array, slice): return false
func SliceContainsAny(slice interface{}, values ...interface{}) bool {
	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)
	rtk := rt.Kind()
	if rtk == reflect.Ptr {
		rtk = rt.Elem().Kind()
		rv = rv.Elem()
	}
	switch rt.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			rvi := rv.Index(i)
			switch rvi.Kind() {
			case reflect.Uintptr, reflect.Array, reflect.Slice, reflect.Interface, reflect.Ptr, reflect.Struct, reflect.Map, reflect.Func, reflect.UnsafePointer:
				for _, value := range values {
					if reflect.DeepEqual(rvi.Interface(), value) {
						return true
					}
				}
			default:
				for _, value := range values {
					if rvi.Interface() == value {
						return true
					}
				}
			}
		}
	}
	return false
}

//SliceContainsAll 判断一个数组(切片)是否包含所有元素(支持[]int, []string, []Struct, []*Struct等)
func SliceContainsAll(slice interface{}, values ...interface{}) bool {
	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)
	rtk := rt.Kind()
	if rtk == reflect.Ptr {
		rtk = rt.Elem().Kind()
		rv = rv.Elem()
	}
	switch rt.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			rvi := rv.Index(i)
			switch rvi.Kind() {
			case reflect.Uintptr, reflect.Array, reflect.Slice, reflect.Interface, reflect.Ptr, reflect.Struct, reflect.Map, reflect.Func, reflect.UnsafePointer:
				for i, value := range values {
					if reflect.DeepEqual(rvi.Interface(), value) {
						values = append(values[:i], values[i+1:]...)
						break
					}
				}
			default:
				for i, value := range values {
					if rvi.Interface() == value {
						values = append(values[:i], values[i+1:]...)
						break
					}
				}
			}
			if len(values) == 0 {
				return true
			}
		}
	}
	return false
}
