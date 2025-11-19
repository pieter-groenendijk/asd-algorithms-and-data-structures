package linkedlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/lists"

func (list *LinkedList[TValue]) getValueNode(index uint) (*valueNode[TValue], error) {
	if index >= list.Size() {
		return nil, lists.ErrOutOfBounds
	}


	var at uint = 0
	currentNode := list.head.next
	for ; at < index ; at++ {
		currentNode = currentNode.next
	}

	return currentNode, nil
}

func (list *LinkedList[TValue]) getNode(index uint) (*node[TValue], error) {
	if index > list.Size() {
		return nil, lists.ErrOutOfBounds
	}

	var at uint = 0
	currentNode := list.head
	for ; at < index ; at++ {
		currentNode = &currentNode.next.node
	}

	return currentNode, nil
}

func (list *LinkedList[TValue]) Get(index uint) (TValue, error) {
	node, err := list.getValueNode(index)
	if err != nil {
		var value TValue
		return value, err
	}

	return node.value, err
}

func (list *LinkedList[TValue]) Add(value TValue) {
	lastIndex := list.size 
	lastNode, _ := list.getNode(lastIndex) // We can safely ignore the error return value

	node := newValueNode(value)
	lastNode.next = node

	list.size++
}

func (list *LinkedList[TValue]) Remove(value TValue) {
	// Determine beforeNode
	beforeNode := list.head
	var removeNode *valueNode[TValue]
	for ;; {
		if beforeNode.next == nil {
			return 
		}

		if beforeNode.next.value == value {
			removeNode = beforeNode.next
			break
		}

		beforeNode = &beforeNode.next.node // implicit conversion, only compiler work, not execution
	}
	afterNode := removeNode.next
	
	beforeNode.next = afterNode
}

func (list *LinkedList[TValue]) Size() uint {
	return list.size
}
