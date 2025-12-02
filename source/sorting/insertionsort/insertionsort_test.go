package insertionsort

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/sorting"
)

func TestSort(t *testing.T) {
	sorting.TestSortInPlace(t, Sort)
}
