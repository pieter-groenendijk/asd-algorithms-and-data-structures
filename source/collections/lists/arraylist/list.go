package arraylist

import (
	"iter"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
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

func (list *ArrayList[TValue]) needToGrow(expectedLength int) bool {
	return expectedLength > len(list.space)
}

func (list *ArrayList[TValue]) grow(expectedLength int) []TValue {
	return make([]TValue, expectedLength*3/2+3)
}

func (list *ArrayList[TValue]) Append(value TValue) {
	list.space[list.length] = value
	list.length++
	list.maybeGrowCapacity(0)
}

func (list *ArrayList[TValue]) IndexOf(value TValue) (int, error) {
	size := list.length
	space := list.space
	for i := 0; i < size; i++ {
		if space[i] == value {
			return i, nil
		}
	}

	return 0, collections.ErrNotFound
}

func (list *ArrayList[TValue]) SetAt(value TValue, index int) error {
	if index >= list.length || index < 0 {
		return lists.ErrOutOfBounds
	}

	list.space[index] = value

	return nil
}

func (list *ArrayList[TValue]) Remove(value TValue) {
	index, err := list.IndexOf(value)
	if err != nil {
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

func (list *ArrayList[TValue]) All() iter.Seq[TValue] {
	return func(yield func(TValue) bool) {
		space := list.space
		size := list.length
		for i := 0; i < size; i++ {
			if !yield(space[i]) {
				return
			}
		}
	}
}
