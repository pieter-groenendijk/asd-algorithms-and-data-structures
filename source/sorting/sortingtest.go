package sorting

import (
	"cmp"
	"math"
	"math/rand"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestSort(t *testing.T, sortFunc SortFunc[int]) {
	type testCase struct {
		name string
		givenItems []int
		expectItems []int
	}

	testCases := []testCase{
		{
			name: "InverselySorted",
			givenItems: []int{8,5,3},
			expectItems: []int{3,5,8},
		},
		{
			name: "PartlySorted",
			givenItems: []int{5,3,8},
			expectItems: []int{3,5,8},
		},
		{
			name: "AlreadySorted",
			givenItems: []int{3,5,8},
			expectItems: []int{3,5,8},
		},
		{
			name: "ZeroItems",
			givenItems: []int{},
			expectItems: []int{},
		},
		{
			name: "OneItem", 
			givenItems: []int{5},
			expectItems: []int{5},
		},
		{
			name: "TwoItems",
			givenItems: []int{5,3},
			expectItems: []int{3,5},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			gotItems := sortFunc(test.givenItems)

			testutils.AssertEquals(t, test.expectItems, gotItems)
		})
	}
}

func AsSortFunc[T cmp.Ordered](sortInPlaceFunc SortInPlaceFunc[T]) SortFunc[T] {
	return func(list []T) []T {
		sortInPlaceFunc(list)

		return list
	}
}

func getRandomInts(length int, maxValue int) []int {
	list := make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = rand.Intn(maxValue)
	}

	return list
}

func benchmarkSortNInPlace(b *testing.B, sortFunc SortFunc[int], length int, maxValue int, doPreSort bool) {
	items := getRandomInts(length, maxValue)
	if doPreSort {
		sortFunc(items)
	}

	for b.Loop() {
		sortFunc(items)
	}
}

// Benchmark:
// length: TINY
// value range: large
// sorted: false
func BenchmarkSortTinyLength(b *testing.B, sortFunc SortFunc[int]) {
	benchmarkSortNInPlace(b, sortFunc, 10, math.MaxInt, false)
}

// Benchmark:
// length: SMALL
// value range: large
// sorted: false
func BenchmarkSortSmallLength(b *testing.B, sortFunc SortFunc[int]) {
	benchmarkSortNInPlace(b, sortFunc, 1_000, math.MaxInt, false)
}

// Benchmark:
// length: MEDIUM
// value range: large
// sorted: false
func BenchmarkSortMediumLength(b *testing.B, sortFunc SortFunc[int]) {
	benchmarkSortNInPlace(b, sortFunc, 10_000, math.MaxInt, false)
}

// Benchmark:
// length: LARGE
// value range: large
// sorted: false
func BenchmarkSortLargeLength(b *testing.B, sortFunc SortFunc[int]) {
	benchmarkSortNInPlace(b, sortFunc, 100_000, math.MaxInt, false)
}

// Benchmark:
// length: medium
// value range: medium
// sorted: TRUE
func BenchmarkSortAlreadySorted(b *testing.B, sortFunc SortFunc[int]) {
	benchmarkSortNInPlace(b, sortFunc, 10_000, math.MaxInt, true)
}

// Benchmark:
// length: medium
// value range: SMALL
// sorted: false
func BenchmarkSmallValueRange(b *testing.B, sortFunc SortFunc[int]) {
	benchmarkSortNInPlace(b, sortFunc, 10_000, 10, false)
}

// Benchmark:
// length: medium
// value range: LARGE
// sorted: false
func BenchmarkLargeValueRange(b *testing.B, sortFunc SortFunc[int]) {
	benchmarkSortNInPlace(b, sortFunc, 10_000, math.MaxInt, false)
}
