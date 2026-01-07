package hashtable

type HashTable[TKey string, TValue any] struct {
	buckets []*Bucket[TKey, TValue]
	size    int
}

func New[TKey string, TValue any]() *HashTable[TKey, TValue] {
	t := &HashTable[TKey, TValue]{
		buckets: make([]*Bucket[TKey, TValue], 10),
		size:    0,
	}

	numOfBuckets := len(t.buckets)
	for i := 0; i < numOfBuckets; i++ {
		t.buckets[i] = NewBucket[TKey, TValue]()
	}

	return t
}
