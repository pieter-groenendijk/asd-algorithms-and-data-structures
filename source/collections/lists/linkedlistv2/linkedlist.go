package linkedlistv2

type EqualsFunc[TValue any] func(thisValue TValue, thatValue TValue) bool

type LinkedList[TValue any] struct {
	head       *Node[TValue]
	tail       *Node[TValue]
	equalsFunc EqualsFunc[TValue]
}

func New[TValue any](equalsFunc EqualsFunc[TValue]) *LinkedList[TValue] {
	var zeroValue TValue
	head := NewNode(zeroValue)

	return &LinkedList[TValue]{
		head:       head,
		tail:       head,
		equalsFunc: equalsFunc,
	}
}
