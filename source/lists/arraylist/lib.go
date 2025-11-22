package arraylist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/lists"

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
