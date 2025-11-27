package hashtable

type node[TKey string, TValue any] struct {
	key TKey
	value TValue
}

func newNode[TKey string, TValue any](key TKey, value TValue) *node[TKey, TValue] {
	return &node[TKey, TValue]{
		key: key,
		value: value,
	}
}
