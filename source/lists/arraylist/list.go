package arraylist

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/lists"
)

func (list *ArrayList[TValue]) Get(index int) (TValue, error) {
	if index >= list.size || index < 0 {
		var value TValue
		return value, lists.ErrOutOfBounds
	}

	value := list.space[index]

	return value, nil
}

func (list *ArrayList[TValue]) Append(value TValue) {
	list.space[list.size] = value
	list.size++
	list.maybeGrowCapacity()
}

func (list *ArrayList[TValue]) Remove(value TValue) {
	index, err := list.indexOf(value)
	if err != nil {
		return
	}

	oldList := list.space
	newSize := list.size - 1
	newCapacity := getCapacityForSize(newSize)
	newList := make([]TValue, newCapacity) 

	i := 0
	for ; i < index; i++ {
		newList[i] = oldList[i]
	}
	for ; i < newSize; i++ {
		newList[i] = oldList[i + 1]
	}

	list.space = newList
	list.size = newSize
}

func (list *ArrayList[TValue]) Size() int {
	return list.size
}
