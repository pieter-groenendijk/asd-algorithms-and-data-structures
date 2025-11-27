package v2

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"

type Node[TKey comparable, TValue any] interface {
	Is(key TKey) bool
	Value() TValue
	Next() Node[TKey, TValue]
	SetNext(node Node[TKey, TValue])
}



// linked list operations
type LinkedList[TKey comparable, TValue any, TNode Node[TKey, TValue]] struct {
	dummyHead TNode // first node; a dummy node to prevent conditionals
	appendAfter TNode // tail, not guaranteed to be a node containing actual values
	size int
}

func newLinkedList[TKey comparable, TValue any, TNode Node[TKey, TValue]](newNode func() TNode) *LinkedList[TKey, TValue, TNode] {
	dummyHead := newNode()

	return &LinkedList[TKey, TValue, TNode]{
		dummyHead: dummyHead,
		appendAfter: dummyHead,
		size: 0,
	}
}

// context operations
func (l *LinkedList[TKey, TValue, TNode]) GetNodeBefore(key TKey) (Node[TKey, TValue], error) {
	currentNode := l.dummyHead
	afterNode := currentNode.Next()
	for {
		if afterNode == nil {
			return nil, collections.ErrNotFound
		}

		if afterNode.Is(key) {
			return currentNode, nil
		}

		currentNode = afterNode
		afterNode = afterNode.Next()
	}
}

func NodeBefore[TKey comparable, TValue any](key TKey, startFrom Node[TKey, TValue]) (Node[TKey, TValue], error) {
	currentNode := startFrom
	afterNode := currentNode.Next()
	for {
		if afterNode == nil {
			return nil, collections.ErrNotFound
		}

		if afterNode.Is(key) {
			return currentNode, nil
		}

		currentNode = afterNode
		afterNode = afterNode.Next()
	}
}

func RemoveAfter[TKey comparable, TValue any](beforeNode Node[TKey, TValue]) {
	afterNode := beforeNode.Next()
	if afterNode == nil {
		return 
	}

	beforeNode.SetNext(afterNode.Next())
}

func InsertAfter[TKey comparable, TValue any, TNode Node[TKey, TValue]](beforeNode, newNode TNode) {
	afterNode := beforeNode.Next()

	beforeNode.SetNext(newNode)
	newNode.SetNext(afterNode)
}

// List operations
func (l *LinkedList[TKey, TValue, TNode]) Append(node TNode) {
	InsertAfter(l.appendAfter, node)
}

func (l *LinkedList[TKey, TValue, TNode]) Prepend(node TNode) {
	InsertAfter(l.dummyHead, node)
}

func (l *LinkedList[TKey, TValue, TNode]) Remove(key TKey) {
	beforeNode, err := NodeBefore(key, l.dummyHead)
	if err != nil {
		return
	}

	RemoveAfter(beforeNode)
}

func (l *LinkedList[TKey, TValue, TNode]) Size() int {
	return l.size
}
