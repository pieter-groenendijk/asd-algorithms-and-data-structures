package lists

// List : a sequence of values
type List[TValue any] interface {
	// SetAt replaces the element at the specified position in this list with the specified element.
	SetAt(value TValue, index int) bool
	// GetAt returns the element at the specified position in this list.
	GetAt(index int) (TValue, bool)

	// Prepend adds the specified element to the beginning of this list.
	Prepend(value TValue)
	// Append adds the specified element to the end of this list.
	Append(value TValue)

	// Remove removes the first occurrence of the specified element from this list, if it is present.
	Remove(value TValue)
	// RemoveAt removes the element at the specified position in this list.
	RemoveAt(index int) error
}
