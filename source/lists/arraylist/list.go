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

func (list *ArrayList[TValue]) maybeGrowCapacity() {
	sizeNow := list.size
	if sizeNow < len(list.space) {
		return
	}

	oldList := list.space
	newList := make([]TValue, getCapacityForSize(sizeNow)) 
	copy(newList, oldList)
	list.space = newList
}

func getCapacityForSize(size int) int {
	return size * 3 / 2 + 3
}

func (list *ArrayList[TValue]) Add(value TValue) {
	list.space[list.size] = value
	list.size++
	list.maybeGrowCapacity()
}

func (list *ArrayList[TValue]) indexOf(value TValue) (int, error) {
	size := list.size
	space := list.space
	for i := 0; i < size; i++ {
		if space[i] == value {
			return i, nil
		}
	}

	return 0, lists.ErrNotFound
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
