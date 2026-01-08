package arraylist

import (
	"fmt"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"
)

func createList() lists.List[int] {
	return New[int](8)
}

func BenchmarkSetAt(b *testing.B) {
	sizes := []int{100, 1_000, 10_000, 100_000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			l := new()

			for i := 0; i < size; i++ {
				l.Append(i)
			}

			index := size / 2

			for b.Loop() {
				l.SetAt(index, 5)
			}
		})
	}
}

func BenchmarkGetAt(b *testing.B) {
	sizes := []int{100, 1_000, 10_000, 100_000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			l := createList()

			for i := 0; i < size; i++ {
				l.Append(i)
			}

			index := size / 2

			for b.Loop() {
				l.GetAt(index)
			}
		})
	}

}

func BenchmarkAppend(b *testing.B) {
	// Worst cases
	sizes := []int{100, 1_000, 10_000, 100_000} // Sure to cause the underlying array to grow with an initial capacity of 8
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Worst-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := New[int](size)

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Append(213)
			}
		})
	}

	// Best cases
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Best-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := New[int](size * 2)

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
	// Worst cases
	sizes := []int{100, 1_000, 10_000, 100_000} // Sure to cause the underlying array to grow with an initial capacity of 8
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Worst-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := New[int](size)

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Prepend(213)
			}
		})
	}

	// Best cases
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Best-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := New[int](size * 2)

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
				l := New[int](size)

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.RemoveAt(size - 1)
			}
		})
	}

	// Worst cases
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Worst-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := New[int](size * 2)

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.RemoveAt(0)
			}
		})
	}
}

func BenchmarkRemove(b *testing.B) {
	//
	sizes := []int{100, 1_000, 10_000, 100_000} // Sure to cause the underlying array to grow with an initial capacity of 8
	for _, size := range sizes {
		b.Run(fmt.Sprintf("LastElement-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := New[int](size)

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Remove(size - 1)
			}
		})
	}

	// Worst cases
	for _, size := range sizes {
		b.Run(fmt.Sprintf("FirstElement-%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				l := New[int](size * 2)

				for i := 0; i < size; i++ {
					l.Append(i)
				}

				b.StartTimer()
				l.Remove(0)
			}
		})
	}
}
