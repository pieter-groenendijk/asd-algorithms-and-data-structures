package hashtable

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists/linkedlist/linkedlistv2"
)

type Bucket[TKey string, TValue any] linkedlistv2.LinkedList[TKey, TValue]

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
		return // nothing to do
	} 

	list.RemoveAfter(node)
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
