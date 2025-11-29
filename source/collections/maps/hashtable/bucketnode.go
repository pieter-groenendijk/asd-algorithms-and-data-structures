package hashtable

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists/linkedlist/linkedlistv2"

type bucketNode[TKey string, TValue any] struct {
	key TKey
	value TValue
	next linkedlistv2.Node[TKey, TValue]
}

func (n *bucketNode[TKey, TValue]) Is(key TKey) bool {
	return n.key == key
}

func (n *bucketNode[TKey, TValue]) Value() TValue {
	return n.value
}

func (n *bucketNode[TKey, TValue]) SetValue(value TValue) {
	n.value = value
}

func (n *bucketNode[TKey, TValue]) Next() linkedlistv2.Node[TKey, TValue] {
	return n.next
}

func (n *bucketNode[TKey, TValue]) SetNext(node linkedlistv2.Node[TKey, TValue]) {
	n.next = node
}
