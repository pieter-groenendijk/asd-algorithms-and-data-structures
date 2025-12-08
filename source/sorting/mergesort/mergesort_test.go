package mergesort

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/sorting"
)

func TestSort(t *testing.T) {
	sorting.TestSort(t, Sort)
}

// Wrapper functions for your benchmarks
func BenchmarkSortTinyLength(b *testing.B) {
	sorting.BenchmarkSortTinyLength(b, Sort)
}

func BenchmarkSortSmallLength(b *testing.B) {
	sorting.BenchmarkSortSmallLength(b, Sort)
}

func BenchmarkSortMediumLength(b *testing.B) {
	sorting.BenchmarkSortMediumLength(b, Sort)
}

func BenchmarkSortLargeLength(b *testing.B) {
	sorting.BenchmarkSortLargeLength(b, Sort)
}

func BenchmarkSortAlreadySorted(b *testing.B) {
	sorting.BenchmarkSortAlreadySorted(b, Sort)
}

func BenchmarkSmallValueRange(b *testing.B) {
	sorting.BenchmarkSmallValueRange(b, Sort)
}

func BenchmarkLargeValueRange(b *testing.B) {
	sorting.BenchmarkLargeValueRange(b, Sort)
}
