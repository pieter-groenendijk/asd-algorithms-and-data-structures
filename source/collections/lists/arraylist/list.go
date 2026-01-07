package arraylist

import (
	"iter"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"
)

func (l *ArrayList[TValue]) SetAt(index int, value TValue) bool {
	if index >= l.length || index < 0 {
		return false
	}

	l.space[index] = value

	return true
}

func (l *ArrayList[TValue]) GetAt(index int) (TValue, bool) {
	if index >= l.length || index < 0 {
		var zeroValue TValue
		return zeroValue, false
	}

	return l.space[index], true
}

func (l *ArrayList[TValue]) Prepend(value TValue) {
	newLength := l.length + 1
	if l.needToGrow(newLength) {
		newSpace := l.grow(newLength)
		for at, value := range l.space {
			newSpace[at+1] = value
		}
		l.space = newSpace
	} else {
		for i := l.length - 1; i >= 0; i-- {
			l.space[i+1] = l.space[i]
		}
	}

	l.space[0] = value
	l.length = newLength
}

func (l *ArrayList[TValue]) Append(value TValue) {
	newLength := l.length + 1
	if l.needToGrow(newLength) {
		newSpace := l.grow(newLength)
		copy(newSpace, l.space[:l.length])
		l.space = newSpace
	}

	l.space[newLength-1] = value
	l.length = newLength
}

func (l *ArrayList[TValue]) IndexOf(value TValue) (int, bool) {
	size := l.length
	space := l.space
	for i := 0; i < size; i++ {
		if space[i] == value {
			return i, true
		}
	}

	return 0, false
}

func (l *ArrayList[TValue]) Remove(value TValue) {
	index, exists := l.IndexOf(value)
	if !exists {
		return
	}

	l.removeAt(index) // can safely ignore returned error
}

func (l *ArrayList[TValue]) RemoveAt(index int) error {
	if index >= l.length || index < 0 {
		return lists.ErrOutOfBounds
	}

	l.removeAt(index)

	return nil
}

func (l *ArrayList[TValue]) Size() int {
	return l.length
}

func (l *ArrayList[TValue]) All() iter.Seq2[int, TValue] {
	return func(yield func(int, TValue) bool) {
		space := l.space
		size := l.length
		for i := 0; i < size; i++ {
			if !yield(i, space[i]) {
				return
			}
		}
	}
}
