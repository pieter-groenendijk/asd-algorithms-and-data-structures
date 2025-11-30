package collections

import "iter"

type Collection[TValue any] interface {
	// Returns an iterator over the elements in this collection.
	All() iter.Seq[TValue]
}
