package linkedlist

import "iter"

func (l *LinkedList[TValue]) GetAt(index int) (TValue, bool) {
	node, exists := l.GetNodeAt(index)
	if !exists {
		var zeroValue TValue
		return zeroValue, false
	}

	return node.value, true
}

func (l *LinkedList[TValue]) SetAt(index int, value TValue) bool {
	node, exists := l.GetNodeAt(index)
	if !exists {
		return false
	}

	node.value = value

	return true
}

func (l *LinkedList[TValue]) Append(value TValue) {
	l.InsertAfter(l.tail, NewNode(value))
}

func (l *LinkedList[TValue]) Prepend(value TValue) {
	l.InsertAfter(l.head, NewNode(value))
}

func (l *LinkedList[TValue]) RemoveAt(index int) error {
	beforeNode, exists := l.GetNodeBeforeAt(index)
	if !exists {
		return nil
	}

	l.RemoveAfter(beforeNode)

	return nil
}

func (l *LinkedList[TValue]) Remove(value TValue) {
	beforeNode, exists := l.GetNodeBefore(value)
	if !exists {
		return
	}

	l.RemoveAfter(beforeNode)
}

func (l *LinkedList[TValue]) All() iter.Seq[*TValue] {
	return func(yield func(*TValue) bool) {
		currNode := l.head.next
		for currNode != nil {
			if !yield(&currNode.value) {
				return
			}
			currNode = currNode.next
		}
	}
}
