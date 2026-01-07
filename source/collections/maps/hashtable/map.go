package hashtable

func (table *HashTable[TKey, TValue]) Get(key TKey) (TValue, bool) {
	return table.getBucket(key).Get(key)
}

func (table *HashTable[TKey, TValue]) Set(key TKey, value TValue) {
	table.getBucket(key).Set(key, value)
}

func (table *HashTable[TKey, TValue]) Unset(key TKey) {
	table.getBucket(key).Unset(key)
}
