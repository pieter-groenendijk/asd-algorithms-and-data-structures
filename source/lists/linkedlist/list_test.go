package linkedlist

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/lists"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestSet(t *testing.T) {
	type testCase struct {
		name string
		preparedValues []int
		execParamValue int
		execParamIndex int
		expectNext *node[int]
		expectErr err
	}

	cases := []testCase{
		{
			name: "firstValueSet",
			preparedValues: []int{5, 3, 5},
			execParamValue: 8,
			execParamIndex: 0,
			expectNext: &node[int]{
				value: 8,
				next: &node[int]{
					value: 3,
					next: &node[int]{
						value: 5,
						next: nil,
					},
				},
			},
			expectErr: nil,
		},
		{
			name: "middleValueSet",
			preparedValues: []int{8, 1, 3},
			execParamValue: 5,
			execParamIndex: 1,
			expectNext: &node[int]{
				value: 8,
				next: &node[int]{
					value: 5,
					next: &node[int]{
						value: 3,
						next: nil,
					},
				},
			},
			expectErr: nil,
		},
		{
			name: "lastValueSet",
			preparedValues: []int{8, 1, 3},
			execParamValue: 5,
			execParamIndex: 2,
			expectNext: &node[int]{
				value: 8,
				next: &node[int]{
					value: 1,
					next: &node[int]{
						value: 5,
						next: nil,
					},
				},
			},
			expectErr: nil,
		},
		{
			name: "singleAndLastValueSet",
			preparedValues: []int{1},
			execParamValue: 5,
			execParamIndex: 0,
			expectNext: &node[int]{
				value: 5,
				next: nil,
			},
			expectErr: nil,
		},
		{
			name: "paramIndexBeforeBounds",
			preparedValues: []int{},
			execParamValue: 5,
			execParamIndex: -5,
			expectNext: nil,
			expectErr: lists.ErrOutOfBounds,
		},
		{
			name: "paramIndexAfterBounds",
			preparedValues: []int{8, 1, 3},
			execParamValue: 5,
			execParamIndex: 3,
			expectNext: &node[int]{
				value: 8,
				next: &node[int]{
					value: 1,
					next: &node[int]{
						value: 3,
						next: nil,
					},
				},
			},
			expectErr: lists.ErrOutOfBounds,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int]()
			for _, preparedValue := range test.preparedValues {
				list.Append(preparedValue)
			}

			err := list.Set(test.execParamValue, test.execParamIndex)

			testutils.AssertEquals(t, test.expectNext, list.head.next)
			testutils.AssertEquals(t, test.expectErr, err)
		})
	}
}

func TestPrepend(t *testing.T) {
	type testCase struct {
		name string
		valueParams []int
		expectNext *node[int]
	}

	cases := []testCase{
		{
			name: "prepend",
			valueParams: []int{5},
			expectNext: &node[int]{
				value: 5,
				next: nil,
			},
		},
		{
			name: "prependMultiple",
			valueParams: []int{5, 3, 7},
			expectNext: &node[int]{
				value: 7,
				next: &node[int]{
					value: 3,
					next: &node[int]{
						value: 5,
						next: nil,
					},
				}, 
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int]()

			for _, valueParam := range test.valueParams {
				list.Prepend(valueParam)
			}

			testutils.AssertEquals(t, test.expectNext, list.head.next)
		})
	}
}

func TestAppend(t *testing.T) {
	t.Run("firstAdd", func(t *testing.T) {
		list := New[int]()
		value := 5

		list.Append(value)

		gotNode := list.head.next
		testutils.AssertEquals(t, *gotNode, *newNode(value))
		testutils.AssertEquals(t, 1, list.size)
	})

	t.Run("subsequentAdds", func(t *testing.T) {
		list := New[int]()
		value := 5

		list.Append(value)

		gotNode := list.head.next
		testutils.AssertEquals(t, *gotNode, *newNode(value))
		testutils.AssertEquals(t, 1, list.size)

		list.Append(value)

		gotNode = gotNode.next
		testutils.AssertEquals(t, *gotNode, *newNode(value))
		testutils.AssertEquals(t, 2, list.size)
	})
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
		{"before sequence", []int{}, 0, 0, lists.ErrOutOfBounds},
		{"after sequence", []int{5, 3, 8}, 3, 0, lists.ErrOutOfBounds},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int]()

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
		expectedValues []int
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
			list := New[int]()

			for _, value := range test.values {
				list.Append(value)
			}

			list.Remove(test.removeValue)

			for index, expectedValue := range test.expectedValues {
				gottenValue, _ := list.Get(index)
				testutils.AssertEquals(t, expectedValue, gottenValue)
			}
		})
	}
}
