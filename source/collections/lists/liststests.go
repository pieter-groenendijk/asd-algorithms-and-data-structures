package lists

import (
	"math"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

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

func TestSetAt(t *testing.T, init func() List[int]) {
	type testCase struct {
		name string
		values []int
		setAt int
		newValue int
		expectValues []int
		expectErr error
	}
	cases := []testCase{
		{"first & last value, one value", []int{2}, 0, 3, []int{3}, nil},
		{"first value, two values", []int{3,6}, 0, 5, []int{5,6}, nil},
		{"first value, three values", []int{100,6,9}, 0, 6, []int{6,6,9}, nil},
		{"middle value, three values", []int{100,6,6}, 1, 9, []int{100,9,6}, nil},
		{"middle value, five values", []int{100,6,9,120,1325}, 3, 121, []int{100,6,9,121,1325}, nil},
		{"last value, two values", []int{3,6}, 1, 5, []int{3,5}, nil},
		{"last value, three values", []int{100,6,9}, 2, 100, []int{100,6,100}, nil},
		{"value not in list", []int{5,5,3,9}, 12, 51, []int{5,5,3,9}, ErrOutOfBounds},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := init()

			for _, value := range test.values {
				list.Append(value)
			}

			gotErr := list.SetAt(test.newValue, test.setAt)

			testutils.AssertEquals(t, test.expectErr, gotErr)

			index := 0
			for value := range list.All() {
				testutils.AssertEquals(t, test.expectValues[index], value)
				index++
			}
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

func TestAppend(t *testing.T, init func() List[int]) {
	type testCase struct {
		name string
		valuesToAdd []int
	}
	cases := []testCase{
		{"Append", []int{3, 5, 9, 1, 3, 4, 5, 1, math.MaxInt, math.MinInt}},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := init()

			for indexToAdd, valueToAdd := range test.valuesToAdd {
				list.Append(valueToAdd)

				index := 0
				for value := range list.All() {
					if index > indexToAdd {
						break
					}

					testutils.AssertEquals(t, test.valuesToAdd[index], value)
					index++
				}
			}
		})
	}
}

func TestPrepend(t *testing.T, init func() List[int]) {
	type testCase struct {
		name string
		valuesToAdd []int
		expectValues []int
	}
	cases := []testCase{
		{"Prepend", []int{3, 5, 9, 1, 3, 4, 5, 1, math.MaxInt, math.MinInt}, []int{math.MinInt, math.MaxInt, 1, 5, 4, 3, 1, 9, 5, 3}},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := init()

			for indexToAdd, valueToAdd := range test.valuesToAdd {
				list.Prepend(valueToAdd)

				length := len(test.valuesToAdd)
				index := 0
				for value := range list.All() {
					if index > indexToAdd {
						break
					}

					testutils.AssertEquals(t, test.expectValues[length - index - 1], value)
					index++
				}
			}
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
