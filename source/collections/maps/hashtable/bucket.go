package hashtable

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists/linkedlist"
)

type Bucket[TKey string, TValue any] linkedlist.LinkedList[*node[TKey, TValue]]

func (b *Bucket[TKey, TValue]) AsList() *linkedlist.LinkedList[*node[TKey, TValue]] {
	return (*linkedlist.LinkedList[*node[TKey, TValue]])(b)
}

func (b *Bucket[TKey, TValue]) getNode(key TKey) (*node[TKey, TValue], error) {
	list := b.AsList()

	for bucketValue := range list.Iterator() {
		if bucketValue.key == key {
			return bucketValue, nil
		}
	}

	return nil, collections.ErrNotFound
}

func (b *Bucket[TKey, TValue]) Get(key TKey) (TValue, error) {
	node, err := b.getNode(key)
	if err != nil {
		var value TValue
		return value, err
	}

	return node.value, nil
}

func (b *Bucket[TKey, TValue]) Set(key TKey, value TValue) {
	list := b.AsList()

	for bucketValue := range list.Iterator() {
		if bucketValue.key == key {
			bucketValue.value = value
			return
		}
	}

	list.Prepend(newNode(key, value))
}

// can be further optimized, with a lower abstraction of a linked list
// now both b.getNode and list.Remove() does a round trip, i.e. now 2n instead of n.
func (b *Bucket[TKey, TValue]) Unset(key TKey) {
	list := b.AsList()

	node, err := b.getNode(key)
	if err != nil {
		return 
	}

	list.Remove(node)
}
