package collections

type Collection interface {
	// Returns true if this collection contains no values.
	IsEmpty() bool
	// Returns an iterator over the elements in this collection.
	Iterator() 
	// Removes all of the values from this list.
	Clear()
	// Returns the number of values in this list.
	Size() int 
}
