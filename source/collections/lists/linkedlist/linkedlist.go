package linkedlist 

// TKey is used for comparison checks. Usually TKey and TValue are one.
type LinkedList[TKey comparable, TValue any] struct {
	dummyHead Node[TKey, TValue] // first node; a dummy node to prevent conditionals
	appendAfter Node[TKey, TValue] // tail, not guaranteed to be a node containing actual values
}

func create[TKey comparable, TValue any](dummyHead Node[TKey, TValue]) *LinkedList[TKey, TValue] {
	return &LinkedList[TKey, TValue]{
		dummyHead: dummyHead,
		appendAfter: dummyHead,
	}
}

func NewWithCustomNode[TKey comparable, TValue any](newNode func() Node[TKey, TValue]) *LinkedList[TKey, TValue] {
	dummyHead := newNode()

	return create(dummyHead)
}

func New[TValue comparable]() *LinkedList[TValue, TValue] {
	var dummyValue TValue

	return create(NewBasicNode(dummyValue, nil))
}
