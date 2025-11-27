package linkedlistv2 

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"

type Node[TKey comparable, TValue any] interface {
	Is(key TKey) bool
	Value() TValue
	SetValue(value TValue)
	Next() Node[TKey, TValue]
	SetNext(node Node[TKey, TValue])
}

func (l *LinkedList[TKey, TValue]) GetNodeBefore(key TKey) (Node[TKey, TValue], error) {
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

func (l *LinkedList[TKey, TValue]) GetNode(key TKey) (Node[TKey, TValue], error) {
	currentNode := l.dummyHead
	for {
		if currentNode == nil {
			return nil, collections.ErrNotFound
		}

		if currentNode.Is(key) {
			return currentNode, nil
		}

		currentNode = currentNode.Next()
	}
}

func (l *LinkedList[TKey, TValue]) RemoveAfter(beforeNode Node[TKey, TValue]) {
	afterNode := beforeNode.Next()
	if afterNode == nil {
		return 
	}

	beforeNode.SetNext(afterNode.Next())

	l.size--
}

func (l *LinkedList[TKey, TValue]) InsertAfter(beforeNode, newNode Node[TKey, TValue]) {
	afterNode := beforeNode.Next()

	beforeNode.SetNext(newNode)
	newNode.SetNext(afterNode)

	l.size++
}
