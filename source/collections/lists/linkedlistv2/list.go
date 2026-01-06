package linkedlistv2

func (l *LinkedList[TValue]) Get(index int) TValue {
	return l.GetNode(index).value
}

func (l *LinkedList[TValue]) Append(value TValue) {
	l.InsertAfter(l.tail, NewNode(value))
}

func (l *LinkedList[TValue]) Prepend(value TValue) {
	l.InsertAfter(l.head, NewNode(value))
}

func (l *LinkedList[TValue]) Remove(value TValue) {
	beforeNode, err := l.GetNodeBefore(value)
	if err != nil {
		return
	}

	l.RemoveAfter(beforeNode)
}
