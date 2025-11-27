package hashtable

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists/linkedlist/linkedlistv2"
)

type Bucket[TKey string, TValue any] linkedlistv2.LinkedList[TKey, TValue]

type bucketNode[TKey string, TValue any] struct {
	key TKey
	value TValue
	next linkedlistv2.Node[TKey, TValue]
}

func (n bucketNode[TKey, TValue]) Is(key TKey) bool {
	return n.key == key
}

func (n bucketNode[TKey, TValue]) Value() TValue {
	return n.value
}

func (n bucketNode[TKey, TValue]) SetValue(value TValue) {
	n.value = value
}

func (n bucketNode[TKey, TValue]) Next() linkedlistv2.Node[TKey, TValue] {
	return n.next
}

func (n bucketNode[TKey, TValue]) SetNext(node linkedlistv2.Node[TKey, TValue]) {
	n.next = node
}

// 	is(key TKey) bool
// 	value() TValue
// 	SetValue(value TValue)
// 	Next() node[TKey, TValue]
// 	SetNext(node node[TKey, TValue])

func (b *Bucket[TKey, TValue]) AsList() *linkedlistv2.LinkedList[TKey, TValue] {
	return (*linkedlistv2.LinkedList[TKey, TValue])(b)
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
		list.Prepend(bucketNode[TKey, TValue]{
			key: key,
			value: value,
		})
	} else {
		node.SetValue(value)
	}
}

// func (b *Bucket[TKey, TValue]) getNode(key TKey) (*node[TKey, TValue], error) {
// 	list := b.AsList()
// 
// 	for bucketValue := range list.Iterator() {
// 		if bucketValue.key == key {
// 			return bucketValue, nil
// 		}
// 	}
// 
// 	return nil, collections.ErrNotFound
// }

// func (b *Bucket[TKey, TValue]) Get(key TKey) (TValue, error) {
// 	node, err := b.getNode(key)
// 	if err != nil {
// 		var value TValue
// 		return value, err
// 	}
// 
// 	return node.value, nil
// }
// 
// func (b *Bucket[TKey, TValue]) Set(key TKey, value TValue) {
// 	list := b.AsList()
// 
// 	for bucketValue := range list.Iterator() {
// 		if bucketValue.key == key {
// 			bucketValue.value = value
// 			return
// 		}
// 	}
// 
// 	list.Prepend(newNode(key, value))
// }
// 
// // can be further optimized, with a lower abstraction of a linked list
// // now both b.getNode and list.Remove() does a round trip, i.e. now 2n instead of n.
// func (b *Bucket[TKey, TValue]) Unset(key TKey) {
// 	list := b.AsList()
// 
// 	node, err := b.getNode(key)
// 	if err != nil {
// 		return 
// 	}
// 
// 	list.Remove(node)
// }
