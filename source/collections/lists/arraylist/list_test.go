package arraylist

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestAppend(t *testing.T) {
	type testCase struct {
		name string
		initialCapacity int
		addValues []int
		expected *ArrayList[int]
	}

	cases := []testCase{
		{
			"no growth needed", 
			2,
			[]int{2}, 
			&ArrayList[int]{
				space: []int{2,0},
				size: 1,
			},
		},
		{
			"growth needed", 
			2,
			[]int{2,5}, 
			&ArrayList[int]{
				space: []int{2,5,0,0,0,0},
				size: 2,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int](test.initialCapacity)

			for _, addValue := range test.addValues {
				list.Append(addValue)
			}

			testutils.AssertEquals(t, test.expected, list)
		})
	}
}

func TestGet(t *testing.T) {
	type testCase struct {
		name string
		values []int
		index int
		expectedValue int
		expectedError error
	}

	cases := []testCase{
		{"first & last value, one value", []int{2}, 0, 2, nil},
		{"first value, two values", []int{3,6}, 0, 3, nil},
		{"first value, three values", []int{100,6,9}, 0, 100, nil},
		{"middle value, three values", []int{100,6,9}, 1, 6, nil},
		{"middle value, five values", []int{100,6,9,120,1325}, 3, 120, nil},
		{"last value, two values", []int{3,6}, 1, 6, nil},
		{"last value, three values", []int{100,6,9}, 2, 9, nil},
		{"before sequence", []int{1,5,2,1}, -1, 0, lists.ErrOutOfBounds},
		{"before sequence", []int{}, 0, 0, lists.ErrOutOfBounds},
		{"after sequence", []int{5, 3, 8}, 3, 0, lists.ErrOutOfBounds},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int](10)

			for _, value := range test.values {
				list.Append(value)
			}

			gotValue, gotError := list.Get(test.index)

			testutils.AssertEquals(t, test.expectedValue, gotValue)
			testutils.AssertEquals(t, test.expectedError, gotError)
		})
	}
}

func TestRemove(t *testing.T) {
	type testCase struct {
		name string
		values []int
		removeValue int
		expected *ArrayList[int]
	}

	cases := []testCase{
		{"first & last value, one value", []int{2}, 2, &ArrayList[int]{[]int{0,0,0}, 0}},
		{"first value, two values", []int{3,6}, 3, &ArrayList[int]{[]int{6,0,0,0}, 1}},
		{"first value, three values", []int{100,6,9}, 100, &ArrayList[int]{[]int{6,9,0,0,0,0}, 2}},
		{"middle value, three values", []int{100,6,6}, 6, &ArrayList[int]{[]int{100,6,0,0,0,0}, 2}},
		{"middle value, five values", []int{100,6,9,120,1325}, 120, &ArrayList[int]{[]int{100,6,9,1325,0,0,0,0,0}, 4}},
		{"last value, two values", []int{3,6}, 6, &ArrayList[int]{[]int{3,0,0,0}, 1}},
		{"last value, three values", []int{100,6,9}, 9, &ArrayList[int]{[]int{100,6,0,0,0,0}, 2}},
		{"value not in list", []int{5,5,3,9}, 12, &ArrayList[int]{[]int{5,5,3,9,0,0}, 4}},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int](6)

			for _, value := range test.values {
				list.Append(value)
			}

			list.Remove(test.removeValue)

			testutils.AssertEquals(t, test.expected, list)
		})
	}
}
