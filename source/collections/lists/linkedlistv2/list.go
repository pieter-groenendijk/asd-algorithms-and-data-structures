package linkedlistv2

func (l *LinkedList[TValue]) GetFrom(index int) (TValue, bool) {
	node, exists := l.GetNodeAt(index)
	if !exists {
		var zeroValue TValue
		return zeroValue, exists
	}

	return node.value, exists
}

func (l *LinkedList[TValue]) Append(value TValue) {
	l.InsertAfter(l.tail, NewNode(value))
}

func (l *LinkedList[TValue]) Prepend(value TValue) {
	l.InsertAfter(l.head, NewNode(value))
}

func (l *LinkedList[TValue]) RemoveAt(index int) {
	beforeNode, exists := l.GetNodeBeforeAt(index)
	if !exists {
		return
	}

	l.RemoveAfter(beforeNode)
}

func (l *LinkedList[TValue]) Remove(value TValue) {
	beforeNode, exists := l.GetNodeBefore(value)
	if !exists {
		return
	}

	l.RemoveAfter(beforeNode)
}
