package linkedlist

import "iter"

func (list *LinkedList[TValue]) Iterator() iter.Seq[TValue] {
	return func(yield func(TValue) bool) {
		var at int = 0
		currentNode := list.head.next
		for ; at < list.size; at++ {
			if !yield(currentNode.value) {
				return
			}
			currentNode = currentNode.next
		}
	}
}
