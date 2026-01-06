package arraylist

func (list *ArrayList[TValue]) needToGrow(expectedLength int) bool {
	return expectedLength > len(list.space)
}

func (list *ArrayList[TValue]) grow(expectedLength int) []TValue {
	return make([]TValue, expectedLength*3/2+3)
}

// TODO: Optimize space efficiency: Resize when lots is unused
func (list *ArrayList[TValue]) removeAt(index int) {
	newLength := list.length - 1
	i := 0
	for ; i < index; i++ {
		list.space[i] = list.space[i]
	}
	for ; i < newLength; i++ {
		list.space[i] = list.space[i+1]
	}

	list.length = newLength
}
