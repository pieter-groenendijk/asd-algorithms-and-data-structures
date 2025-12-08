package lists

import (
	"testing"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

// // A sequence of values
// type List[TValue any] interface {
// 	// Returns the element at the specified position in this list.
// 	Get(index int) (TValue, error)
// 	// Returns true if this list contains the specified element.
//
// 	// Replaces the element at the specified position in this list with the specified element.
// 	SetAt(value TValue, index int) error
// 	// Adds the specified element to the beginning of this list.
// 	Prepend(value TValue)
// 	// Adds the specified element to the end of this list.
// 	Append(value TValue)
// 	// Inserts the specified element at the specified position in this list.
// 	// InsertAt(value TValue, index uint)
//
// 	// Removes the first occurrence of the specified element from this list, if it is present.
// 	Remove(value TValue)
// 	// Removes the element at the specified position in this list.
// 	RemoveAt(index uint) error
//
// 	All() iter.Seq[TValue]
// }

// func TestAppend[T List[int]](t *testing.T) {
// 	type testCase struct {
// 		name string
// 		initialCapacity int
// 		addValues []int
// 		expected *T
// 	}
//
// 	cases := []testCase{
// 		{
// 			"no growth needed",
// 			2,
// 			[]int{2},
// 			&T{
// 				space: []int{2,0},
// 				size: 1,
// 			},
// 		},
// 		{
// 			"growth needed",
// 			2,
// 			[]int{2,5},
// 			&ArrayList[int]{
// 				space: []int{2,5,0,0,0,0},
// 				size: 2,
// 			},
// 		},
// 	}
//
// 	for _, test := range cases {
// 		t.Run(test.name, func(t *testing.T) {
// 			list := New[int](test.initialCapacity)
//
// 			for _, addValue := range test.addValues {
// 				list.Append(addValue)
// 			}
//
// 			testutils.AssertEquals(t, test.expected, list)
// 		})
// 	}
// }

func TestGet(t *testing.T, init func() List[int]) {
	type testCase struct {
		name string
		values []int
		index int
		expectedValue int
		expectedError error
	}
	testCases := []testCase{
		{"first & last value, one value", []int{2}, 0, 2, nil},
		{"first value, two values", []int{3,6}, 0, 3, nil},
		{"first value, three values", []int{100,6,9}, 0, 100, nil},
		{"middle value, three values", []int{100,6,9}, 1, 6, nil},
		{"middle value, five values", []int{100,6,9,120,1325}, 3, 120, nil},
		{"last value, two values", []int{3,6}, 1, 6, nil},
		{"last value, three values", []int{100,6,9}, 2, 9, nil},
		{"before sequence", []int{1,5,2,1}, -1, 0, ErrOutOfBounds},
		{"before sequence", []int{}, 0, 0, ErrOutOfBounds},
		{"after sequence", []int{5, 3, 8}, 3, 0, ErrOutOfBounds},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			list := init()

			for _, value := range test.values {
				list.Append(value)
			}

			gotValue, gotError := list.Get(test.index)

			testutils.AssertEquals(t, test.expectedValue, gotValue)
			testutils.AssertEquals(t, test.expectedError, gotError)
		})
	}
}

func TestRemove(t *testing.T, init func() List[int]) {
	type testCase struct {
		name string
		values []int
		removeValue int
		expectValues []int
	}

	cases := []testCase{
		{"first & last value, one value", []int{2}, 2, []int{}},
		{"first value, two values", []int{3,6}, 3, []int{6}},
		{"first value, three values", []int{100,6,9}, 100, []int{6,9}},
		{"middle value, three values", []int{100,6,6}, 6, []int{100,6}},
		{"middle value, five values", []int{100,6,9,120,1325}, 120, []int{100,6,9,1325}},
		{"last value, two values", []int{3,6}, 6, []int{3}},
		{"last value, three values", []int{100,6,9}, 9, []int{100,6}},
		{"value not in list", []int{5,5,3,9}, 12, []int{5,5,3,9}},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := init()

			for _, value := range test.values {
				list.Append(value)
			}

			list.Remove(test.removeValue)


			testutils.AssertEquals(t, len(test.expectValues), list.Size())

			index := 0
			for value := range list.All() {
				testutils.AssertEquals(t, test.expectValues[index], value)
				index++
			}
		})
	}
}

func TestRemoveAt(t *testing.T, init func() List[int]) {
	type testCase struct {
		name string
		values []int
		removeAt int
		expectValues []int
		expectErr error
	}

	cases := []testCase{
		{"FirstAndLastValueOneValue", []int{2}, 0, []int{}, nil},
		{"FirstValueTwoValues", []int{3,6}, 0, []int{6}, nil},
		{"FirstValueThreeValues", []int{100,6,9}, 0, []int{6,9}, nil},
		{"MiddleValueThreeValues", []int{100,6,6}, 1, []int{100,6}, nil},
		{"MiddleValueFiveValues", []int{100,6,9,120,1325}, 3, []int{100,6,9,1325}, nil},
		{"LastValueTwoValues", []int{3,6}, 1, []int{3}, nil},
		{"LastValueThreeValues", []int{100,6,9}, 2, []int{100,6}, nil},
		{"ValueNotInList", []int{5,5,3,9}, 12, []int{5,5,3,9}, ErrOutOfBounds},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := init()

			for _, value := range test.values {
				list.Append(value)
			}

			gotErr := list.RemoveAt(test.removeAt)


			testutils.AssertEquals(t, len(test.expectValues), list.Size())

			index := 0
			for value := range list.All() {
				testutils.AssertEquals(t, test.expectValues[index], value)
				index++
			}

			testutils.AssertEquals(t, test.expectErr, gotErr)
		})
	}
}

func TestSize(t *testing.T, init func() List[int]) {
	type testCase struct {
		name string
		values []int
		expectSize int
	}

	testCases := []testCase{
		{
			name: "zeroItems",
			values: []int{},
			expectSize: 0,
		},
		{
			name: "oneItem",
			values: []int{5},
			expectSize: 1,
		},
		{
			name: "twoItems",
			values: []int{5,3},
			expectSize: 2,
		},
		{	
			name: "threeItems",
			values: []int{5,3,9},
			expectSize: 3,
		},
		{
			name: "128Items",
			values: make([]int, 128),
			expectSize: 128,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			list := init()
			for _, value := range test.values {
				list.Append(value)
			}

			gotSize := list.Size()

			testutils.AssertEquals(t, test.expectSize, gotSize)
		})
	}
}
