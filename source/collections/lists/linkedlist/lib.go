package linkedlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"

func (list *LinkedList[TValue]) getNode(index int) (*node[TValue], error) {
	if index > list.size || index < 0 {
		return nil, lists.ErrOutOfBounds
	}

	var at int = 0
	currentNode := list.head
	for ; at < index ; at++ {
		currentNode = currentNode.next
	}

	return currentNode, nil
}
