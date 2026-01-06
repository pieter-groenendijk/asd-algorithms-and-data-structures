package lists

import (
	"math"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func ptr[T any](value T) *T {
	return &value
}

func TestGet(t *testing.T, init func() List[int]) {
	type testCase struct {
		name          string
		values        []int
		index         int
		expectedValue *int
	}
	testCases := []testCase{
		{"first & last value, one value", []int{2}, 0, ptr(2)},
		{"first value, two values", []int{3, 6}, 0, ptr(3)},
		{"first value, three values", []int{100, 6, 9}, 0, ptr(100)},
		{"middle value, three values", []int{100, 6, 9}, 1, ptr(6)},
		{"middle value, five values", []int{100, 6, 9, 120, 1325}, 3, ptr(120)},
		{"last value, two values", []int{3, 6}, 1, ptr(6)},
		{"last value, three values", []int{100, 6, 9}, 2, ptr(9)},
		{"before sequence", []int{1, 5, 2, 1}, -1, nil},
		{"before sequence", []int{}, 0, nil},
		{"after sequence", []int{5, 3, 8}, 3, nil},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			list := init()

			for _, value := range test.values {
				list.Append(value)
			}

			if test.expectedValue == nil {
				defer func() {
					recover()
				}()
				list.Get(test.index)
			} else {
				gotValue := list.Get(test.index)

				testutils.AssertEquals(t, *test.expectedValue, gotValue)
			}
		})
	}
}

func TestSetAt(t *testing.T, init func() List[int]) {
	type testCase struct {
		name         string
		values       []int
		setAt        int
		newValue     int
		expectValues []int
		expectErr    error
	}
	cases := []testCase{
		{"first & last value, one value", []int{2}, 0, 3, []int{3}, nil},
		{"first value, two values", []int{3, 6}, 0, 5, []int{5, 6}, nil},
		{"first value, three values", []int{100, 6, 9}, 0, 6, []int{6, 6, 9}, nil},
		{"middle value, three values", []int{100, 6, 6}, 1, 9, []int{100, 9, 6}, nil},
		{"middle value, five values", []int{100, 6, 9, 120, 1325}, 3, 121, []int{100, 6, 9, 121, 1325}, nil},
		{"last value, two values", []int{3, 6}, 1, 5, []int{3, 5}, nil},
		{"last value, three values", []int{100, 6, 9}, 2, 100, []int{100, 6, 100}, nil},
		{"value not in list", []int{5, 5, 3, 9}, 12, 51, []int{5, 5, 3, 9}, ErrOutOfBounds},
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
		name         string
		values       []int
		removeValue  int
		expectValues []int
	}

	cases := []testCase{
		{"first & last value, one value", []int{2}, 2, []int{}},
		{"first value, two values", []int{3, 6}, 3, []int{6}},
		{"first value, three values", []int{100, 6, 9}, 100, []int{6, 9}},
		{"middle value, three values", []int{100, 6, 6}, 6, []int{100, 6}},
		{"middle value, five values", []int{100, 6, 9, 120, 1325}, 120, []int{100, 6, 9, 1325}},
		{"last value, two values", []int{3, 6}, 6, []int{3}},
		{"last value, three values", []int{100, 6, 9}, 9, []int{100, 6}},
		{"value not in list", []int{5, 5, 3, 9}, 12, []int{5, 5, 3, 9}},
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
		name         string
		values       []int
		removeAt     int
		expectValues []int
		expectErr    error
	}

	cases := []testCase{
		{"FirstAndLastValueOneValue", []int{2}, 0, []int{}, nil},
		{"FirstValueTwoValues", []int{3, 6}, 0, []int{6}, nil},
		{"FirstValueThreeValues", []int{100, 6, 9}, 0, []int{6, 9}, nil},
		{"MiddleValueThreeValues", []int{100, 6, 6}, 1, []int{100, 6}, nil},
		{"MiddleValueFiveValues", []int{100, 6, 9, 120, 1325}, 3, []int{100, 6, 9, 1325}, nil},
		{"LastValueTwoValues", []int{3, 6}, 1, []int{3}, nil},
		{"LastValueThreeValues", []int{100, 6, 9}, 2, []int{100, 6}, nil},
		{"ValueNotInList", []int{5, 5, 3, 9}, 12, []int{5, 5, 3, 9}, ErrOutOfBounds},
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
		name        string
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
	t.Run("Prepend", func(t *testing.T) {
		list := init()

		valueOne := 3
		valueTwo := 5
		valueThree := 9
		valueFour := 1
		valueFive := 3
		valueSix := 4
		valueSeven := 5
		valueEight := 1
		valueNine := math.MaxInt
		valueTen := math.MinInt

		list.Prepend(valueOne)

		testutils.AssertEquals(t, valueOne, list.Get(0))

		list.Prepend(valueTwo)

		testutils.AssertEquals(t, valueTwo, list.Get(0))
		testutils.AssertEquals(t, valueOne, list.Get(1))

		list.Prepend(valueThree)

		testutils.AssertEquals(t, valueThree, list.Get(0))
		testutils.AssertEquals(t, valueTwo, list.Get(1))
		testutils.AssertEquals(t, valueOne, list.Get(2))

		list.Prepend(valueFour)

		testutils.AssertEquals(t, valueFour, list.Get(0))
		testutils.AssertEquals(t, valueThree, list.Get(1))
		testutils.AssertEquals(t, valueTwo, list.Get(2))
		testutils.AssertEquals(t, valueOne, list.Get(3))

		list.Prepend(valueFive)

		testutils.AssertEquals(t, valueFive, list.Get(0))
		testutils.AssertEquals(t, valueFour, list.Get(1))
		testutils.AssertEquals(t, valueThree, list.Get(2))
		testutils.AssertEquals(t, valueTwo, list.Get(3))
		testutils.AssertEquals(t, valueOne, list.Get(4))

		list.Prepend(valueSix)

		testutils.AssertEquals(t, valueSix, list.Get(0))
		testutils.AssertEquals(t, valueFive, list.Get(1))
		testutils.AssertEquals(t, valueFour, list.Get(2))
		testutils.AssertEquals(t, valueThree, list.Get(3))
		testutils.AssertEquals(t, valueTwo, list.Get(4))
		testutils.AssertEquals(t, valueOne, list.Get(5))

		list.Prepend(valueSeven)

		testutils.AssertEquals(t, valueSeven, list.Get(0))
		testutils.AssertEquals(t, valueSix, list.Get(1))
		testutils.AssertEquals(t, valueFive, list.Get(2))
		testutils.AssertEquals(t, valueFour, list.Get(3))
		testutils.AssertEquals(t, valueThree, list.Get(4))
		testutils.AssertEquals(t, valueTwo, list.Get(5))
		testutils.AssertEquals(t, valueOne, list.Get(6))

		list.Prepend(valueEight)

		testutils.AssertEquals(t, valueEight, list.Get(0))
		testutils.AssertEquals(t, valueSeven, list.Get(1))
		testutils.AssertEquals(t, valueSix, list.Get(2))
		testutils.AssertEquals(t, valueFive, list.Get(3))
		testutils.AssertEquals(t, valueFour, list.Get(4))
		testutils.AssertEquals(t, valueThree, list.Get(5))
		testutils.AssertEquals(t, valueTwo, list.Get(6))
		testutils.AssertEquals(t, valueOne, list.Get(7))

		list.Prepend(valueNine)

		testutils.AssertEquals(t, valueNine, list.Get(0))
		testutils.AssertEquals(t, valueEight, list.Get(1))
		testutils.AssertEquals(t, valueSeven, list.Get(2))
		testutils.AssertEquals(t, valueSix, list.Get(3))
		testutils.AssertEquals(t, valueFive, list.Get(4))
		testutils.AssertEquals(t, valueFour, list.Get(5))
		testutils.AssertEquals(t, valueThree, list.Get(6))
		testutils.AssertEquals(t, valueTwo, list.Get(7))
		testutils.AssertEquals(t, valueOne, list.Get(8))

		list.Prepend(valueTen)

		testutils.AssertEquals(t, valueTen, list.Get(0))
		testutils.AssertEquals(t, valueNine, list.Get(1))
		testutils.AssertEquals(t, valueEight, list.Get(2))
		testutils.AssertEquals(t, valueSeven, list.Get(3))
		testutils.AssertEquals(t, valueSix, list.Get(4))
		testutils.AssertEquals(t, valueFive, list.Get(5))
		testutils.AssertEquals(t, valueFour, list.Get(6))
		testutils.AssertEquals(t, valueThree, list.Get(7))
		testutils.AssertEquals(t, valueTwo, list.Get(8))
		testutils.AssertEquals(t, valueOne, list.Get(9))
	})
}

func TestSize(t *testing.T, init func() List[int]) {
	type testCase struct {
		name       string
		values     []int
		expectSize int
	}

	testCases := []testCase{
		{
			name:       "zeroItems",
			values:     []int{},
			expectSize: 0,
		},
		{
			name:       "oneItem",
			values:     []int{5},
			expectSize: 1,
		},
		{
			name:       "twoItems",
			values:     []int{5, 3},
			expectSize: 2,
		},
		{
			name:       "threeItems",
			values:     []int{5, 3, 9},
			expectSize: 3,
		},
		{
			name:       "128Items",
			values:     make([]int, 128),
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
