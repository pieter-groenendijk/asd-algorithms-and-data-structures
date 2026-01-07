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
	for ; index < newLength; index++ {
		list.space[index] = list.space[index+1]
	}

	list.length = newLength
}
