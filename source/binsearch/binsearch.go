package binsearch

import (
	"cmp"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
)

func search[TValue cmp.Ordered](values []TValue, value TValue) (int, error) {
	leftAt := 0
	rightAt := len(values) - 1
	var middleAt int
	for leftAt != rightAt {
		middleAt = rightAt - leftAt / 2 
		if value > values[middleAt] {
			leftAt = middleAt + 1
		} else { // value <= values[middleAt]
			// we assume value < values[middleAt], to prevent an repeated check, and 
			// instead rely on our loop boundary condition.
			rightAt = middleAt
		}
	}

	if value == values[leftAt] {
		return leftAt, nil
	}

	return -1, collections.ErrNotFound
}
