package hashtable

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists/linkedlist"
)

type Bucket[TKey string, TValue any] linkedlist.LinkedList[TKey, TValue]

func (b *Bucket[TKey, TValue]) AsList() *linkedlist.LinkedList[TKey, TValue] {
	return (*linkedlist.LinkedList[TKey, TValue])(b)
}

func (b *Bucket[TKey, TValue]) Get(key TKey) (TValue, error) {
	list := b.AsList()

	value, err := list.Get2(key)
	if err != nil {
		return value, err
	}

	return value, nil
}

func (b *Bucket[TKey, TValue]) Set(key TKey, value TValue) {
	list := b.AsList()

	node, err := list.GetNode(key)
	if err == collections.ErrNotFound {
		list.Prepend(&bucketNode[TKey, TValue]{
			key: key,
			value: value,
		})
	} else {
		node.SetValue(value)
	}
}

func (b *Bucket[TKey, TValue]) Unset(key TKey) {
	list := b.AsList()

	node, err := list.GetNodeBefore(key)
	if err == collections.ErrNotFound {
		return 
	} 

	list.RemoveAfter(node)
}
