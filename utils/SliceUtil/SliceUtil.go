package SliceUtil

import (
	"errors"
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
func Range(start, end, step int64) []int64 {
	var result []int64
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result
}

// Split splits a slice into multiple sub-slices of equal size.
// It takes a slice and a size parameter as input.
// The function returns a 2-dimensional slice containing the sub-slices.
func Split[T any](slices []T, size int64) [][]T {
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
	return result
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
