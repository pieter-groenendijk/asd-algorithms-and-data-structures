package binheap

import "cmp"

type Node[TPriority cmp.Ordered, TValue any] struct {
	order TPriority
	value TValue
}

func NewNode[TPriority cmp.Ordered, TValue any](priority TPriority, value TValue) *Node[TPriority, TValue] {
	return &Node[TPriority, TValue]{
		order: priority,
		value: value,
	}
}
