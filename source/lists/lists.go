package lists

import "errors"

var ErrOutOfBounds = errors.New("Out of bounds")

// A sequence of values
type List[TValue any] interface {
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

	// Returns the number of elements in this list.
	Size() uint 
}
