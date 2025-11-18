package linkedlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/lists"

// All List interface operations
/*
	// Returns the element at the specified position in this list.
	Get(index uint) (TValue, error)
	// Returns true if this list contains the specified element.
	// Contains(value TValue) bool
	// Returns true if this list contains no elements.
	// IsEmpty() bool
	// Returns an iterator over the elements in this list in proper sequence.
	// Iterator()

	// Replaces the element at the specified position in this list with the specified element.
	// SetAt(value TValue, index uint)
	// Appends the specified element to the end of this list.
	Add(value TValue)
	// Inserts the specified element at the specified position in this list.
	// AddAt(value TValue, index uint)

	// Removes the first occurrence of the specified element from this list, if it is present.
	Remove(value TValue)
	// Removes the element at the specified position in this list.
	// RemoveAt(index uint)
	// Removes all of the elements from this list.
	// Clear()
*/

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

func (list *LinkedList[TValue]) Get(index uint) (TValue, error) {
	node, err := list.getValueNode(index)
	if err != nil {
		var value TValue
		return value, err
	}

	return node.value, err
}

func (list *LinkedList[TValue]) Add(value TValue) {
	lastIndex := list.Size() - 1
	lastNode, _ := list.getValueNode(lastIndex) // We can safely ignore the error return value

	node := newValueNode(value)
	lastNode.next = node
}

func (list *LinkedList[TValue]) Remove(value TValue) {
	// Determine beforeNode
	beforeNode := list.head
	var removeNode *valueNode[TValue]
	for ; beforeNode.next != nil; {
		if beforeNode.next.value == value {
			removeNode = beforeNode.next
			break
		}

		beforeNode = &beforeNode.next.node // implicit conversion to be a *node, not *valueNode, to be compatible
	}
	afterNode := removeNode.next
	
	beforeNode.next = afterNode
}

func (list *LinkedList[TValue]) Size() uint {
	return list.size
}
