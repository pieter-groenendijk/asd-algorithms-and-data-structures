package hashtable

type HashTable[TKey string, TValue comparable] struct {
	buckets []*Bucket[TKey, TValue]
	size int
}

func New[TKey string, TValue comparable]() *HashTable[TKey, TValue] {
	return &HashTable[TKey, TValue]{
		buckets: make([]*Bucket[TKey, TValue], 10),
		size: 0,
	}
}
