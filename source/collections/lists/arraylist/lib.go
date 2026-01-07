package arraylist

func (l *ArrayList[TValue]) needToGrow(expectedLength int) bool {
	return expectedLength > len(l.space)
}

func (l *ArrayList[TValue]) grow(expectedLength int) []TValue {
	return make([]TValue, expectedLength*3/2+3)
}

// TODO: Optimize space efficiency: Resize when lots is unused
func (l *ArrayList[TValue]) removeAt(index int) {
	newLength := l.length - 1
	for ; index < newLength; index++ {
		l.space[index] = l.space[index+1]
	}

	l.length = newLength
}
