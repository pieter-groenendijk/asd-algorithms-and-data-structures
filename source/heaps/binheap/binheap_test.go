package binheap

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/heaps"
)

func newHeap() heaps.Heap[int, int] {
	return New[int, int](32)
}

func TestPushAndPop(t *testing.T) {
	heaps.TestPushAndPop(t, newHeap)
}
