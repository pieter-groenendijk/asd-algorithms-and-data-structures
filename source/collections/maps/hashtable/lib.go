package hashtable

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/hashing/divhash"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/hashing/fnvhash"
)

func (table *HashTable[TKey, TValue]) hash(key TKey) uint32 {
	index, _ := divhash.Hash( // We can safely ignore size errors, as long we initiate a hashtable with a non-zero length.
		fnvhash.Hash([]byte(key)),
		uint32(len(table.buckets)),
	)

	return index
}

func (table *HashTable[TKey, TValue]) getBucket(key TKey) *Bucket[TKey, TValue] {
	return table.buckets[table.hash(key)]
}
