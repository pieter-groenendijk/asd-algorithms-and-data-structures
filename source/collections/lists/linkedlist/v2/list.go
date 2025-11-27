package v2

func (l *LinkedList[TKey, TValue]) Append(node Node[TKey, TValue]) {
	l.InsertAfter(l.appendAfter, node)
	l.appendAfter = node
}

func (l *LinkedList[TKey, TValue]) Prepend(node Node[TKey, TValue]) {
	l.InsertAfter(l.dummyHead, node)
}

func (l *LinkedList[TKey, TValue]) Remove(key TKey) {
	beforeNode, err := l.GetNodeBefore(key)
	if err != nil {
		return
	}

	l.RemoveAfter(beforeNode)
}

func (l *LinkedList[TKey, TValue]) Size() int {
	return l.size
}
