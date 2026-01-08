package binsearch

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/sorting/mergesort"
)

func BenchmarkSearch(b *testing.B) {
	// Any case is log(N)
	sizes := []int{100, 1_000, 10_000, 100_000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Best-%d", size), func(b *testing.B) {
			l := make([]int, size)
			for i := 0; i < size; i++ {
				l[i] = rand.Int()
			}
			value := l[0]
			l = mergesort.Sort(l)

			for b.Loop() {
				Search(l, value)
			}
		})
	}
}
