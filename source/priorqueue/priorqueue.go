package priorqueue

import (
	"cmp"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/heaps/binheap"
)

// Alias of a binary heap
type PriorityQueue[TPriority cmp.Ordered, TValue any] = binheap.BinHeap[TPriority, TValue]

func New[TPriority cmp.Ordered, TValue any](capacity int) *PriorityQueue[TPriority, TValue] {
	return binheap.New[TPriority, TValue](capacity)
}
