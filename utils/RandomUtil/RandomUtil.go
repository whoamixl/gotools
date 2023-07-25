package RandomUtil

import (
	cRand "crypto/rand"
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
