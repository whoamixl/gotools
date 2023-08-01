package utils

import (
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	bitSet  []bool
	hashNum uint
}

func NewBloomFilter(size uint, hashNum uint) *BloomFilter {
	return &BloomFilter{
		bitSet:  make([]bool, size),
		hashNum: hashNum,
	}
}
func (bf *BloomFilter) Add(element string) {
	for i := uint(0); i < bf.hashNum; i++ {
		index := bf.hash(element, i) % uint(len(bf.bitSet))
		bf.bitSet[index] = true
	}
}
func (bf *BloomFilter) Contains(element string) bool {
	for i := uint(0); i < bf.hashNum; i++ {
		index := bf.hash(element, i) % uint(len(bf.bitSet))
		if !bf.bitSet[index] {
			return false
		}
	}
	return true
}
func (bf *BloomFilter) hash(element string, seed uint) uint {
	h := fnv.New32a()
	h.Write([]byte(element))
	return (uint(h.Sum32()) + seed) % math.MaxUint32
}
