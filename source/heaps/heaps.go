package heaps

import "cmp"

type Heap[TOrder cmp.Ordered, TValue any] interface {
	// Inserts the value with associated priority into the collection
	Push(order TOrder, value TValue)
	// Removes and returns the value with the highest associated priority
	Pop() (TValue, bool)
}
