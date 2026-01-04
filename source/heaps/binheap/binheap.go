package binheap

import "cmp"

type BinHeap[TOrder cmp.Ordered, TValue any] struct {
	nodes []Node[TOrder, TValue]
}

func New[TOrder cmp.Ordered, TValue any](capacity int) *BinHeap[TOrder, TValue] {
	return &BinHeap[TOrder, TValue]{
		nodes: make([]Node[TOrder, TValue], 0, capacity),
	}
}

func (h *BinHeap[TOrder, TValue]) Push(order TOrder, value TValue) {
	h.nodes = append(h.nodes, *NewNode(order, value))

	length := len(h.nodes)
	if length == 1 {
		return
	}

	// While "heap order" incorrect, swap inserted with parent
	insertedAt := length - 1
	nodes := h.nodes
	for {
		parentAt := (insertedAt - 1) / 2 // `>> 1` is the same as `/ 2`, but more performant
		if parentAt < 0 {
			break
		}

		if order <= nodes[parentAt].order {
			break
		}

		nodes[parentAt], nodes[insertedAt] = nodes[insertedAt], nodes[parentAt]
	}
}


func (h *BinHeap[TOrder, TValue]) Pop() (TValue, bool) {
	length := len(h.nodes)
	if length == 0 {
		var value TValue
		return value, false
	}

	// Extract first while we still can
	first := h.nodes[0]

	// Move "last" to "first", and remove "last"
	h.nodes[0] = h.nodes[length-1]
	h.nodes = h.nodes[:length-1]
	length--

	// While "heap order" incorrect -> swap largest child with inserted
	nodes := h.nodes
	insertedAt := 0
	for {
		leftChildAt := insertedAt*2 + 1
		rightChildAt := insertedAt*2 + 2
		largestAt := insertedAt

		if leftChildAt < length && nodes[leftChildAt].order > nodes[largestAt].order {
			largestAt = leftChildAt
		}
		if rightChildAt < length && nodes[rightChildAt].order > nodes[largestAt].order {
			largestAt = rightChildAt
		}

		if largestAt == insertedAt { // We didn't find a larger child
			break
		}

		nodes[largestAt], nodes[insertedAt] = nodes[insertedAt], nodes[largestAt]
		insertedAt = largestAt
	}

	return first.value, true
}
