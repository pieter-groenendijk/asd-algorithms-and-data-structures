package lists

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"

// A sequence of values
type List[TValue any] interface {
	collections.Collection

	// Returns the element at the specified position in this list.
	Get(index int) (TValue, error)
	// Returns true if this list contains the specified element.

	// Replaces the element at the specified position in this list with the specified element. 
	SetAt(value TValue, index int) error
	// Adds the specified element to the beginning of this list. 
	Prepend(value TValue)
	// Adds the specified element to the end of this list. 
	Append(value TValue)
	// Inserts the specified element at the specified position in this list.
	// InsertAt(value TValue, index uint)

	// Removes the first occurrence of the specified element from this list, if it is present.
	Remove(value TValue)
	// Removes the element at the specified position in this list.
	// RemoveAt(index uint)
}
