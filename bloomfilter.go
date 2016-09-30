package bloomfilter

import (
	"hash/fnv"

	"github.com/nukr/bloomfilter/bitset"
)

// BloomFilter ...
type BloomFilter struct {
	bits      *bitset.Bitset
	numBits   int
	numHashes int
}

// New ...
func New(numBits, numHashes int) BloomFilter {
	bs := bitset.MakeBitSet(uint(numBits))
	return BloomFilter{
		bits:      &bs,
		numBits:   numBits,
		numHashes: numHashes,
	}
}

// Add ...
func (bf BloomFilter) Add(item []byte) {
	hashA, hashB := hashFNV1a(item)
	for i := 0; i < bf.numHashes; i++ {
		hash := (hashA + hashB*uint32(i)) % uint32(bf.numBits)
		bf.bits.SetBit(uint(hash))
	}
}

// MayContain ...
func (bf BloomFilter) MayContain(item []byte) bool {
	hashA, hashB := hashFNV1a(item)
	for i := 0; i < bf.numHashes; i++ {
		hash := (hashA + hashB*uint32(i)) % uint32(bf.numBits)
		if !bf.bits.Bit(uint(hash)) {
			return false
		}
	}
	return true
}

func hashFNV1a(input []byte) (uint32, uint32) {
	hash := fnv.New64a()
	hash.Write(input)
	value64 := hash.Sum64()
	return uint32(value64 & 0xFFFFFFFF), uint32(value64 >> 32)
}
