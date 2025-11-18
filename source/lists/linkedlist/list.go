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

func (list *LinkedList[TValue]) Get(index uint) (TValue, error) {
	if index >= list.Size() {
		var value TValue	
		return value, lists.ErrOutOfBounds
	}


	var at uint = 0
	currentNode := list.head.next
	for ; at < index ; at++ {
		currentNode = currentNode.next
	}

	return currentNode.value, nil
}

func (list *LinkedList[TValue]) Add(value TValue) {

}

func (list *LinkedList[TValue]) Remove(value TValue) {

}

func (list *LinkedList[TValue]) Size() uint {
	return list.size
}
