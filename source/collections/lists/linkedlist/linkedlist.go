package linkedlist

type LinkedList[TValue comparable] struct {
	head *node[TValue]
	size int
}

func New[TValue comparable]() *LinkedList[TValue] {
	return &LinkedList[TValue]{
		head: newDummyNode[TValue](),
		size: 0,
	}
}
