package arraylist

import (
	"iter"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"
)

func (list *ArrayList[TValue]) Get(index int) (TValue, error) {
	if index >= list.size || index < 0 {
		var value TValue
		return value, lists.ErrOutOfBounds
	}

	value := list.space[index]

	return value, nil
}

func (list *ArrayList[TValue]) Prepend(value TValue) {
	list.space[list.size] = value
	list.size++
	list.maybeGrowCapacity(1)
}

func (list *ArrayList[TValue]) Append(value TValue) {
	list.space[list.size] = value
	list.size++
	list.maybeGrowCapacity(0)
}

func (list *ArrayList[TValue]) IndexOf(value TValue) (int, error) {
	size := list.size
	space := list.space
	for i := 0; i < size; i++ {
		if space[i] == value {
			return i, nil
		}
	}

	return 0, collections.ErrNotFound 
}

func (list *ArrayList[TValue]) SetAt(value TValue, index int) error {
	if index >= list.size || index < 0 {
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
	if index >= list.size || index < 0 {
		return lists.ErrOutOfBounds
	}
	
	list.removeAt(index)

	return nil
}

func (list *ArrayList[TValue]) Size() int {
	return list.size
}

func (list *ArrayList[TValue]) All() iter.Seq[TValue] {
	return func(yield func(TValue) bool) {
		space := list.space
		size := list.size
		for i := 0; i < size; i++ {
			if !yield(space[i]) {
				return
			}
		}
	}
}
