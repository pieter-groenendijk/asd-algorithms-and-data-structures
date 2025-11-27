package linkedlistv2

type LinkedList[TKey comparable, TValue any] struct {
	dummyHead Node[TKey, TValue] // first node; a dummy node to prevent conditionals
	appendAfter Node[TKey, TValue] // tail, not guaranteed to be a node containing actual values
	size int
}

func newLinkedList[TKey comparable, TValue any](newNode func() Node[TKey, TValue]) *LinkedList[TKey, TValue] {
	dummyHead := newNode()

	return &LinkedList[TKey, TValue]{
		dummyHead: dummyHead,
		appendAfter: dummyHead,
		size: 0,
	}
}
