package hashtable

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists/linkedlist"
)

type pair[TKey string, TValue any] struct {
	key   TKey
	value TValue
}

type Bucket[TKey string, TValue any] struct {
	pairs linkedlist.LinkedList[pair[TKey, TValue]]
}

func NewBucket[TKey string, TValue any]() *Bucket[TKey, TValue] {
	return &Bucket[TKey, TValue]{
		pairs: *linkedlist.New[pair[TKey, TValue]](func(thisValue pair[TKey, TValue], thatValue pair[TKey, TValue]) bool {
			return thisValue.key == thatValue.key
		}),
	}
}

func (b *Bucket[TKey, TValue]) Get(key TKey) (TValue, bool) {
	for pair := range b.pairs.All() {
		if pair.key == key {
			return pair.value, true
		}
	}

	var zeroValue TValue
	return zeroValue, false
}

func (b *Bucket[TKey, TValue]) Set(key TKey, value TValue) {
	for pair := range b.pairs.All() {
		if pair.key == key {
			pair.value = value
			return
		}
	}

	b.pairs.Prepend(pair[TKey, TValue]{
		key:   key,
		value: value,
	})
}

func (b *Bucket[TKey, TValue]) Unset(key TKey) {
	b.pairs.Remove(pair[TKey, TValue]{
		key: key,
	})
}
