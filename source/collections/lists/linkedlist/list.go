package linkedlist

func (l *LinkedList[TKey, TValue]) Get2(key TKey) (TValue, error) {
	_node, err := l.GetNode(key)
	if err != nil {
		var value TValue
		return value, err
	}

	return _node.Value(), nil
}

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
