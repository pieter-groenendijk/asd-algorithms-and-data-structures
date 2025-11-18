package linkedlist

type node[TValue any] struct {
	value TValue
	next *node[TValue]
}

// A dummy node placed at the front of every LinkedList
// By doing this, many special cases of linked list operations can be eliminated
type headNode[TValue any] struct {
	next *node[TValue]
}

func newNode[TValue any](value TValue) *node[TValue] {
	return &node[TValue]{
		value: value,
	}
}

func newHeadNode[TValue any]() *headNode[TValue] {
	return &headNode[TValue]{}
}
