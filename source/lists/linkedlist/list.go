package linkedlist

func (list *LinkedList[TValue]) Get(index int) (TValue, error) {
	node, err := list.getNode(index + 1)
	if err != nil {
		var value TValue
		return value, err
	}

	return node.value, err
}

func (list *LinkedList[TValue]) Prepend(value TValue) {
	beforeNode := list.head
	afterNode := beforeNode.next
	node := newNode(value)

	beforeNode.next = node
	node.next = afterNode

	list.size++
}

func (list *LinkedList[TValue]) Append(value TValue) {
	lastIndex := list.size 
	lastNode, _ := list.getNode(lastIndex) // We can safely ignore the error return value; we're never out of bounds

	node := newNode(value)
	lastNode.next = node

	list.size++
}

func (list *LinkedList[TValue]) Remove(value TValue) {
	beforeNode := list.head
	var removeNode *node[TValue]
	for ;; {
		if beforeNode.next == nil {
			return 
		}

		if beforeNode.next.value == value {
			removeNode = beforeNode.next
			break
		}

		beforeNode = beforeNode.next 
	}
	afterNode := removeNode.next
	
	beforeNode.next = afterNode
}

func (list *LinkedList[TValue]) Size() int {
	return list.size
}
