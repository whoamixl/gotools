// Package SliceUtil 切片工具
package SliceUtil

import (
	"errors"
	"fmt"
	"reflect"
)

// IsEmpty checks if a string slice is empty.
// Parameter slice is the string slice to be checked.
// The return value bool indicates whether the slice is empty (true) or not (false).
func IsEmpty(slice []string) bool {
	return slice == nil || len(slice) == 0
}

// AddAll is a generic function that takes multiple arguments of type T and returns a slice of type T.
func AddAll[T any](slices ...[]T) []T {
	var mergedSlice []T
	for _, slice := range slices {
		mergedSlice = append(mergedSlice, slice...)
	}
	return mergedSlice
}

// Range generate a int64 slice, params is start number, end number and step
func Range(start, end, step int64) ([]int64, error) {
	if step <= 0 {
		return nil, errors.New("step is not less than zero")
	}
	if !((end - start) > 0) {
		return nil, errors.New("end must greater than start")
	}
	var result []int64
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result, nil
}

// Split splits a slice into multiple sub-slices of equal size.
// It takes a slice and a size parameter as input.
// The function returns a 2-dimensional slice containing the sub-slices.
func Split[T any](slices []T, size int64) ([][]T, error) {
	if size <= 0 {
		return nil, errors.New("size is not less than zero")
	}
	length := int64(len(slices))
	chunks := (length + size - 1) / size
	result := make([][]T, chunks)
	for i := int64(0); i < chunks; i++ {
		start := i * size
		end := (i + 1) * size
		if end > length {
			end = length
		}
		result[i] = slices[start:end]
	}
	return result, nil
}

// Zip takes two arrays as input, the first array as keys and the second array as values.
// It returns a map with interface{} as key type and error.
func Zip[T any, Y any](keys []T, values []Y) (map[interface{}]Y, error) {
	result := make(map[interface{}]Y)
	if len(keys) != len(values) {
		return result, errors.New("error：The length of keys is not equal to the length of value")
	}
	for i := 0; i < len(keys); i++ {
		result[keys[i]] = values[i]
	}
	return result, nil
}

// Contains checks if an element exists in a slice
func Contains(slice interface{}, element interface{}) bool {
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(item, element) {
			return true
		}
	}
	return false
}

// ToString make a slice to string
func ToString[T any](slice []T) string {
	str := ""
	for i, v := range slice {
		if i == len(slice)-1 {
			str = str + fmt.Sprintf("%v", v)
		} else {
			str = str + fmt.Sprintf("%v", v) + ","
		}
	}
	return str
}

// Join make a slice to string with join a string
func Join[T any](slice []T, join string) string {
	str := ""
	for i, v := range slice {
		if i == len(slice)-1 {
			str = str + fmt.Sprintf("%v", v)
		} else {
			str = str + fmt.Sprintf("%v", v) + join
		}
	}
	return str
}
