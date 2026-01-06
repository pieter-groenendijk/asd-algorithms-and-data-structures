package arraylist

import (
	"iter"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"
)

func (list *ArrayList[TValue]) Get(index int) TValue {
	return list.space[index]
}

func (list *ArrayList[TValue]) Prepend(value TValue) {
	newLength := list.length + 1
	if list.needToGrow(newLength) {
		newSpace := list.grow(newLength)
		for at, value := range list.space {
			newSpace[at+1] = value
		}
		list.space = newSpace
	} else {
		for i := list.length - 1; i >= 0; i-- {
			list.space[i+1] = list.space[i]
		}
	}

	list.space[0] = value
	list.length = newLength
}

func (list *ArrayList[TValue]) Append(value TValue) {
	newLength := list.length + 1
	if list.needToGrow(newLength) {
		newSpace := list.grow(newLength)
		copy(newSpace, list.space[:list.length])
		list.space = newSpace
	}

	list.space[newLength-1] = value
	list.length = newLength
}

func (list *ArrayList[TValue]) IndexOf(value TValue) (int, bool) {
	size := list.length
	space := list.space
	for i := 0; i < size; i++ {
		if space[i] == value {
			return i, true
		}
	}

	return 0, false
}

func (list *ArrayList[TValue]) SetAt(value TValue, index int) error {
	if index >= list.length || index < 0 {
		return lists.ErrOutOfBounds
	}

	list.space[index] = value

	return nil
}

func (list *ArrayList[TValue]) Remove(value TValue) {
	index, exists := list.IndexOf(value)
	if !exists {
		return
	}

	list.removeAt(index) // can safely ignore returned error
}

func (list *ArrayList[TValue]) RemoveAt(index int) error {
	if index >= list.length || index < 0 {
		return lists.ErrOutOfBounds
	}

	list.removeAt(index)

	return nil
}

func (list *ArrayList[TValue]) Size() int {
	return list.length
}

func (list *ArrayList[TValue]) All() iter.Seq2[int, TValue] {
	return func(yield func(int, TValue) bool) {
		space := list.space
		size := list.length
		for i := 0; i < size; i++ {
			if !yield(i, space[i]) {
				return
			}
		}
	}
}
