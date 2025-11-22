package linkedlist

// A dummy node placed at the front of every LinkedList
// By doing this, many special cases of linked list operations can be eliminated
type node[TValue comparable] struct {
	value TValue
	next *node[TValue]
}

func newDummyNode[TValue comparable]() *node[TValue] {
	return &node[TValue]{}
}

func newNode[TValue comparable](value TValue) *node[TValue] {
	return &node[TValue]{
		value: value,
		next: nil,
	}
}
