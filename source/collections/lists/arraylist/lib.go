package arraylist

func (list *ArrayList[TValue]) maybeGrowCapacity(insertOffset int) {
	sizeNow := list.size
	if sizeNow < len(list.space) {
		return
	}

	oldList := list.space
	newList := make([]TValue, getCapacityForSize(sizeNow)) 
	copy(newList[insertOffset:], oldList)
	list.space = newList
}

func getCapacityForSize(size int) int {
	return size * 3 / 2 + 3
}


func (list *ArrayList[TValue]) removeAt(index int) {
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
