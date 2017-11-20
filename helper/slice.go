package helper

import "reflect"

// panic if s is not a slice
func ReverseSlice(s interface{}) {
	reflectValue := reflect.ValueOf(s)
	if reflectValue.Kind() != reflect.Slice {
		return
	}
	size := reflectValue.Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
