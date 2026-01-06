package linkedlistv2

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"

type Node[TValue any] struct {
	value TValue
	next  *Node[TValue]
}

func NewNode[TValue any](value TValue) *Node[TValue] {
	return &Node[TValue]{
		value: value,
		next:  nil,
	}
}

func (l *LinkedList[TValue]) GetNodeBefore(value TValue) (*Node[TValue], error) {
	currNode := l.head
	afterNode := currNode.next
	for afterNode != nil {
		if l.equalsFunc(value, afterNode.value) {
			return currNode, nil
		}

		currNode = afterNode
		afterNode = afterNode.next
	}

	return nil, collections.ErrNotFound
}

func (l *LinkedList[TValue]) GetNode(index int) *Node[TValue] {
	currNode := l.head.next
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}

	return currNode
}

func (l *LinkedList[TValue]) RemoveAfter(beforeNode *Node[TValue]) {
	afterNode := beforeNode.next
	if afterNode == nil {
		l.tail = afterNode
		return
	}

	beforeNode.next = afterNode.next
}

func (l *LinkedList[TValue]) InsertAfter(beforeNode, newNode *Node[TValue]) {
	afterNode := beforeNode.next
	if afterNode == nil {
		l.tail = newNode
	}

	beforeNode.next = newNode
	newNode.next = afterNode
}
