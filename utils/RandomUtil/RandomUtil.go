// Package RandomUtil 随机工具
package RandomUtil

import (
	cRand "crypto/rand"
	"errors"
	"math/rand"
	"time"
)

// RandomInt generates a random integer between the start and end values (inclusive).
// Parameters:
// - start: the starting value of the range
// - end: the ending value of the range
// Returns:
// - int64: a random integer between the start and end values
func RandomInt(start, end int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return start + r.Int63n(end-start)
}

// RandomBytes generates a random byte slice of the specified length.
// Parameters:
// - length: the length of the byte slice to generate
// Returns:
// - []byte: a random byte slice of the specified length
func RandomBytes(length int64) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := cRand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// RandomEle returns a random element from the given slice.
func RandomEle[T any](slice []T) (T, error) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(slice) == 0 {
		errors.New("input slice cannot be empty")
	}
	index := rand.Intn(len(slice))
	return slice[index], nil
}

// RandomUniqueEleSet Select a specified number of unique elements at random from a given slice
func RandomUniqueEleSet[T comparable](slice []T, length int) ([]T, error) {
	if length > len(slice) {
		return nil, errors.New("Requested number of elements is greater than the length of the slice")
	}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	uniqueSet := make(map[T]bool)
	result := make([]T, 0, length)
	for len(uniqueSet) < length {
		index := rand.Intn(len(slice))
		element := slice[index]
		if !uniqueSet[element] {
			uniqueSet[element] = true
			result = append(result, element)
		}
	}
	return result, nil
}

// RandomString generates a random string of given length
func RandomString(length int64) string {
	return randomString("abcdefghijklmnopqrstuvwxyz0123456789", length)
}

// RandomNumbers generates a random string of given length with only number
func RandomNumbers(length int64) string {
	return randomString("0123456789", length)
}

func randomString(str string, length int64) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	var i int64
	for i = 0; i < length; i++ {
		index := rand.Intn(len(str))
		result[i] = str[index]
	}
	return string(result)
}
