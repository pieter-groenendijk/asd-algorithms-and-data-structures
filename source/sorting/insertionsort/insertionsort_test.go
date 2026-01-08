package insertionsort

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/sorting"
)

func TestSort(t *testing.T) {
	sorting.TestSort(t, sorting.AsSortFunc[int](Sort))
}

// Wrapper functions for your benchmarks
func BenchmarkSortTinyLength(b *testing.B) {
	sorting.BenchmarkSortTinyLength(b, sorting.AsSortFunc[int](Sort))
}

func BenchmarkSortSmallLength(b *testing.B) {
	sorting.BenchmarkSortSmallLength(b, sorting.AsSortFunc[int](Sort))
}

func BenchmarkSortMediumLength(b *testing.B) {
	sorting.BenchmarkSortMediumLength(b, sorting.AsSortFunc[int](Sort))
}

func BenchmarkSortLargeLength(b *testing.B) {
	sorting.BenchmarkSortLargeLength(b, sorting.AsSortFunc[int](Sort))
}
