package linkedlist

import (
	"fmt"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"
)

func createList() lists.List[int] {
	return New[int](func(thisValue int, thatValue int) bool {
		return thisValue == thatValue
	})
}

func BenchmarkSetAt(b *testing.B) {
	sizes := []int{100, 1_000, 10_000, 100_000}

	// First Element
	for _, size := range sizes {
		b.Run(fmt.Sprintf("FirstElement-%d", size), func(b *testing.B) {
			l := createList()

			for i := 0; i < size; i++ {
				l.Append(i)
			}

			for b.Loop() {
				l.SetAt(0, 5)
			}
		})
	}

	// Last Element
	for _, size := range sizes {
		b.Run(fmt.Sprintf("LastElement-%d", size), func(b *testing.B) {
			l := createList()

			for i := 0; i < size; i++ {
				l.Append(i)
			}

			for b.Loop() {
				l.SetAt(size-1, 5)
			}
		})
	}
}

func BenchmarkGetAt(b *testing.B) {
	sizes := []int{100, 1_000, 10_000, 100_000}

	// First Element
	for _, size := range sizes {
		b.Run(fmt.Sprintf("FirstElement-%d", size), func(b *testing.B) {
			l := createList()

			for i := 0; i < size; i++ {
				l.Append(i)
			}

			for b.Loop() {
				l.GetAt(0)
			}
		})
	}

	// Last Element
	for _, size := range sizes {
		b.Run(fmt.Sprintf("LastElement-%d", size), func(b *testing.B) {
			l := createList()

			for i := 0; i < size; i++ {
				l.Append(i)
			}

			for b.Loop() {
				l.GetAt(size - 1)
			}
		})
	}
}

func BenchmarkAppend(b *testing.B) {
	// Any cases
	sizes := []int{100, 1_000, 10_000, 100_000} // Sure to cause the underlying array to grow with an initial capacity of 8
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := createList()

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Append(213)
			}
		})
	}
}

func BenchmarkPrepend(b *testing.B) {
	// Any cases
	sizes := []int{100, 1_000, 10_000, 100_000} // Sure to cause the underlying array to grow with an initial capacity of 8
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := createList()

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Prepend(213)
			}
		})
	}
}

func BenchmarkRemoveAt(b *testing.B) {
	// Best cases
	sizes := []int{100, 1_000, 10_000, 100_000} // Sure to cause the underlying array to grow with an initial capacity of 8
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Best-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := createList()

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.RemoveAt(0)
			}
		})
	}

	// Worst cases
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Worst-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := createList()

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.RemoveAt(size - 1)
			}
		})
	}
}

func BenchmarkRemove(b *testing.B) {
	// Best cases
	sizes := []int{100, 1_000, 10_000, 100_000} // Sure to cause the underlying array to grow with an initial capacity of 8
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Best-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := createList()

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Remove(0)
			}
		})
	}

	// Worst cases
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Worst-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := createList()

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Remove(size - 1)
			}
		})
	}
}
