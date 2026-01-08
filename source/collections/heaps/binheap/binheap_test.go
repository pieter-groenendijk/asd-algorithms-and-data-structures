package binheap

import (
	"testing"

	heaps2 "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/heaps"
)

func newHeap() heaps2.Heap[int, int] {
	return New[int, int](32)
}

func TestPushAndPop(t *testing.T) {
	heaps2.TestPushAndPop(t, newHeap)
}
