package linkedlist

type LinkedList[TValue comparable] struct {
	head *node[TValue]
	size uint
}

func New[TValue comparable]() *LinkedList[TValue] {
	return &LinkedList[TValue]{
		head: newNode[TValue](),
		size: 0,
	}
}
