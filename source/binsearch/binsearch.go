package binsearch

import (
	"cmp"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
)

func Search[TValue cmp.Ordered](values []TValue, value TValue) (int, error) {
	rightAt := len(values) - 1
	if rightAt < 0 {
		return -1, collections.ErrNotFound
	}

	leftAt := 0
	var middleAt int
	for leftAt != rightAt {
		middleAt = (rightAt - leftAt) / 2 + leftAt
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
