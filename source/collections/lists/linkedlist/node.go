package linkedlist

type Node[TValue any] struct {
	value TValue
	next  *Node[TValue]
}

func NewNode[TValue any](value TValue) *Node[TValue] {
	return &Node[TValue]{
		value: value,
		next:  nil,
	}
}

func (l *LinkedList[TValue]) GetNodeBefore(value TValue) (*Node[TValue], bool) {
	currNode := l.head
	afterNode := currNode.next
	for afterNode != nil {
		if l.equalsFunc(value, afterNode.value) {
			return currNode, true
		}

		currNode = afterNode
		afterNode = afterNode.next
	}

	return nil, false
}

func (l *LinkedList[TValue]) GetNodeBeforeAt(index int) (*Node[TValue], bool) {
	if index < 0 || index >= l.length {
		return nil, false
	}

	currNode := l.head
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}

	if currNode == nil {
		return nil, false
	}

	return currNode, true
}

func (l *LinkedList[TValue]) GetNodeAt(index int) (*Node[TValue], bool) {
	if index < 0 || index >= l.length {
		return nil, false
	}

	currNode := l.head.next
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}

	if currNode == nil {
		return nil, false
	}

	return currNode, true
}

func (l *LinkedList[TValue]) RemoveAfter(beforeNode *Node[TValue]) {
	afterNode := beforeNode.next
	if afterNode == nil {
		return
	}

	beforeNode.next = afterNode.next

	if beforeNode.next == nil {
		l.tail = beforeNode
	}

	l.length--
}

func (l *LinkedList[TValue]) InsertAfter(beforeNode, newNode *Node[TValue]) {
	afterNode := beforeNode.next
	if afterNode == nil {
		l.tail = newNode
	}

	beforeNode.next = newNode
	newNode.next = afterNode

	l.length++
}
