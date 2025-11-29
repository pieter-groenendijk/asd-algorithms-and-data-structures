package linkedlist

import "iter"

func (l *LinkedList[TKey, TValue]) All() iter.Seq[TValue] {
	return func(yield func(TValue) bool) {
		currentNode := l.dummyHead.Next()
		for {
			if currentNode == nil {
				return
			}
			if !yield(currentNode.Value()) {
				return
			}

			currentNode = currentNode.Next()
		}
	}
}
