package fnvhash

import (
	"encoding/binary"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/hashing"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/hashing/divhash"
)

const offset uint32 = 2166136261
const prime uint32 = 16777619

func Hash(value uint32, size uint32) (uint32, error) {
	if size == 0 {
		return 0, hashing.ErrNoHashRoom
	} 

	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, value)

	hash := offset
	for _, dataByte := range data {
		hash ^= uint32(dataByte)
		hash *= prime
	}

	return divhash.Hash(hash, size)
}
