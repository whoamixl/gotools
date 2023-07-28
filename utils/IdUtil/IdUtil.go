// Package IdUtil id工具
package IdUtil

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

// UUID generate an uuid of 16 length
func UUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// Snowflake struct
type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	workerID  int64
	sequence  int64
}

// NewSnowflake creates a new Snowflake instance
func NewSnowflake(workerID int64) int64 {
	sn := &Snowflake{
		timestamp: 0,
		workerID:  workerID,
		sequence:  0,
	}
	return sn.generateID()
}

// GenerateID generates a unique ID using Snowflake algorithm
func (s *Snowflake) generateID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	currentTimestamp := time.Now().UnixNano() / 1000000
	if s.timestamp == currentTimestamp {
		s.sequence = (s.sequence + 1) & 4095
		if s.sequence == 0 {
			// Wait until next millisecond
			for currentTimestamp <= s.timestamp {
				currentTimestamp = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = currentTimestamp
	id := (currentTimestamp << 22) | (s.workerID << 12) | s.sequence
	return id
}
