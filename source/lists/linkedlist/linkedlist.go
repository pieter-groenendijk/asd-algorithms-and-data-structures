package linkedlist


type LinkedList[TValue any] struct {
	head *headNode[TValue]
	size uint
}

func New[TValue any]() *LinkedList[TValue] {
	return &LinkedList[TValue]{
		head: newHeadNode[TValue](),
		size: 0,
	}
}
