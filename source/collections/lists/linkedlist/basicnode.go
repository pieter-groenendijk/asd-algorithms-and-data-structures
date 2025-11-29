package linkedlist

type BasicNode[TValue comparable] struct {
	value TValue
	next Node[TValue, TValue]
}

func NewBasicNode[TValue comparable](value TValue, next Node[TValue, TValue]) *BasicNode[TValue] {
	return &BasicNode[TValue]{
		value: value,
		next: next,
	}
}

func (n *BasicNode[TValue]) Is(value TValue) bool {
	return n.value == value
}

func (n *BasicNode[TValue]) Value() TValue {
	return n.value
}

func (n *BasicNode[TValue]) SetValue(value TValue) {
	n.value = value
}

func (n *BasicNode[TValue]) Next() Node[TValue, TValue] {
	return n.next
}

func (n *BasicNode[TValue]) SetNext(node Node[TValue, TValue]) {
	n.next = node
}
