package linkedlist

// A dummy node placed at the front of every LinkedList
// By doing this, many special cases of linked list operations can be eliminated
type node[TValue comparable] struct {
	next *valueNode[TValue]
}

type valueNode[TValue comparable] struct {
	node[TValue]
	value TValue
}

func newNode[TValue comparable]() *node[TValue] {
	return &node[TValue]{}
}

func newValueNode[TValue comparable](value TValue) *valueNode[TValue] {
	return &valueNode[TValue]{
		value: value,
	}
}
