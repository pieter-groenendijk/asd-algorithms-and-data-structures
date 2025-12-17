package btree

import (
	"cmp"
	"iter"
)

// TODO: Make graph inteface compliant
// TODO: Make collection interface compliant
// TODO: Make actually balanced

type BTree[TValue cmp.Ordered] Node[TValue]

func New[TValue cmp.Ordered](value TValue) *BTree[TValue] {
	return (*BTree[TValue])(NewNode(value))
}

func (t *BTree[TValue]) Insert(value TValue) {
	currentNode := (*Node[TValue])(t)
	for {
		currentValue := currentNode.value

		if value < currentValue {
			if currentNode.left == nil {
				currentNode.left = NewNode(value)
				return
			} 

			currentNode = currentNode.left
		} else if value > currentValue {
			if currentNode.right == nil {
				currentNode.right = NewNode(value)
				return
			} 
				
			currentNode = currentNode.right
		} else { // implicitly newValue == value
			currentNode.value = value // Although the comparison is the same, the underlying data may not be (e.g. key-value structure), thus we copy.
			return
		}
	}
}

func (t *BTree[TValue]) Remove(value TValue) {
	currentNode := (*Node[TValue])(t)
	for {
		currentValue := currentNode.value

		if value < currentValue {
			if currentNode.left == nil {
				currentNode.left = NewNode(value)
				return
			} 

			currentNode = currentNode.left
		} else if value > currentValue {
			if currentNode.right == nil {
				currentNode.right = NewNode(value)
				return
			} 
				
			currentNode = currentNode.right
		} else { // implicitly newValue == value
			currentNode.value = value // Although the comparison is the same, the underlying data may not be (e.g. key-value structure), thus we copy.
			return
		}
	}
}

func (t *BTree[TValue]) Get(value TValue) {

}

func (t *BTree[TValue]) All() iter.Seq[TValue] {
	return func(yield func(TValue) bool) {

	}
}

type Node[TValue cmp.Ordered] struct {
	value TValue
	left *Node[TValue]
	right *Node[TValue]
}

func NewNode[TValue cmp.Ordered](value TValue) *Node[TValue] {
	return &Node[TValue]{
		value: value,
		left: nil,
		right: nil,
	}
}
