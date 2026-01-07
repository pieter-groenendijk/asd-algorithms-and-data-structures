package arraylist

import (
	"math"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func new() *ArrayList[int] {
	return New[int](16)
}

// helper method to get an easy conceptual view of a list
func (l *ArrayList[TValue]) toSlice() []TValue {
	return l.space[0:l.length]
}

func TestSetAt(t *testing.T) {
	testCases := []struct {
		name           string
		operations     func(l *ArrayList[int]) bool
		expectedReturn bool
		expectedState  []int
	}{
		{
			name: "SetBeforeBounds",
			operations: func(l *ArrayList[int]) bool {
				return l.SetAt(0, 5)
			},
			expectedReturn: false,
			expectedState:  []int{},
		},
		{
			name: "SetAfterBounds",
			operations: func(l *ArrayList[int]) bool {
				l.Append(5)
				l.Append(1)

				return l.SetAt(2, 3)
			},
			expectedReturn: false,
			expectedState:  []int{5, 1},
		},
		{
			name: "SetFirstElement",
			operations: func(l *ArrayList[int]) bool {
				l.Append(10)
				l.Append(123)
				l.Append(16)

				return l.SetAt(0, 5)
			},
			expectedReturn: true,
			expectedState:  []int{5, 123, 16},
		},
		{
			name: "SetMiddleElement",
			operations: func(l *ArrayList[int]) bool {
				l.Append(10)
				l.Append(123)
				l.Append(16)

				return l.SetAt(1, 5)
			},
			expectedReturn: true,
			expectedState:  []int{10, 5, 16},
		},
		{
			name: "SetLastElement",
			operations: func(l *ArrayList[int]) bool {
				l.Append(10)
				l.Append(123)
				l.Append(16)

				return l.SetAt(2, 5)
			},
			expectedReturn: true,
			expectedState:  []int{10, 123, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l := new()

			gotReturn := testCase.operations(l)

			testutils.AssertEquals(t, testCase.expectedReturn, gotReturn)
			testutils.AssertEquals(t, testCase.expectedState, l.toSlice())
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name       string
		operations func(l *ArrayList[int])
		expected   []int
	}{
		{
			name: "RemoveFirstElement",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.Remove(20)
			},
			expected: []int{123, 3},
		},
		{
			name: "RemoveMiddleElement",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.Remove(123)
			},
			expected: []int{20, 3},
		},
		{
			name: "RemoveLastElement",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.Remove(3)
			},
			expected: []int{20, 123},
		},
		{
			name: "RemoveNonExistent",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.Remove(143)
			},
			expected: []int{20, 123, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l := new()

			testCase.operations(l)

			testutils.AssertEquals(t, testCase.expected, l.toSlice())
		})
	}
}

func TestRemoveAt(t *testing.T) {
	testCases := []struct {
		name       string
		operations func(l *ArrayList[int])
		expected   []int
	}{
		{
			name: "RemoveFirstElement",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(0)
			},
			expected: []int{123, 3},
		},
		{
			name: "RemoveMiddleElement",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(1)
			},
			expected: []int{20, 3},
		},
		{
			name: "RemoveLastElement",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(2)
			},
			expected: []int{20, 123},
		},
		{
			name: "RemoveBeforeBounds",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(-1)
			},
			expected: []int{20, 123, 3},
		},
		{
			name: "RemoveAfterBounds",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(3)
			},
			expected: []int{20, 123, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l := new()

			testCase.operations(l)

			testutils.AssertEquals(t, testCase.expected, l.toSlice())
		})
	}
}

func TestGet(t *testing.T) {
	type returnValue struct {
		value  int
		exists bool
	}

	testCases := []struct {
		name       string
		operations func(l *ArrayList[int]) []returnValue
		expected   []returnValue
	}{
		{
			name: "GetFirst",
			operations: func(l *ArrayList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetAt(0)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{10, true}},
		},
		{
			name: "GetMiddle",
			operations: func(l *ArrayList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetAt(1)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{20, true}},
		},
		{
			name: "GetLast",
			operations: func(l *ArrayList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetAt(2)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{5, true}},
		},
		{
			name: "GetAt",
			operations: func(l *ArrayList[int]) []returnValue {
				l.Append(10)

				valueOne, existsOne := l.GetAt(0)

				l.Append(20)

				valueTwo, existsTwo := l.GetAt(0)

				l.Append(123)

				valueThree, existsThree := l.GetAt(0)

				return []returnValue{
					{valueOne, existsOne},
					{valueTwo, existsTwo},
					{valueThree, existsThree},
				}
			},
			expected: []returnValue{
				{10, true},
				{10, true},
				{10, true},
			},
		},
		{
			name: "GetAtNegativeIndex",
			operations: func(l *ArrayList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetAt(-1)

				valueTwo, existsTwo := l.GetAt(-5)

				return []returnValue{
					{value, exists},
					{valueTwo, existsTwo},
				}
			},
			expected: []returnValue{
				{0, false},
				{0, false},
			},
		},
		{
			name: "GetBeforeBounds",
			operations: func(l *ArrayList[int]) []returnValue {
				value, exists := l.GetAt(0)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{0, false}},
		},
		{
			name: "GetAfterBounds",
			operations: func(l *ArrayList[int]) []returnValue {
				l.Append(10)
				l.Append(20)

				value, exists := l.GetAt(2)

				l.Append(5)

				valueTwo, existsTwo := l.GetAt(3)

				valueThree, existsThree := l.GetAt(math.MaxInt)

				return []returnValue{
					{value, exists},
					{valueTwo, existsTwo},
					{valueThree, existsThree},
				}
			},
			expected: []returnValue{
				{0, false},
				{0, false},
				{0, false},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l := new()

			results := testCase.operations(l)

			testutils.AssertEquals(t, testCase.expected, results)
		})
	}
}

func TestAppend(t *testing.T) {
	testCases := []struct {
		name       string
		operations func(l *ArrayList[int])
		expected   []int
	}{
		{
			name: "AppendOne",
			operations: func(l *ArrayList[int]) {
				l.Append(20)
			},
			expected: []int{20},
		},
		{
			name: "AppendTwo",
			operations: func(l *ArrayList[int]) {
				l.Append(10)
				l.Append(20)
			},
			expected: []int{10, 20},
		},
		{
			name: "AppendThree",
			operations: func(l *ArrayList[int]) {
				l.Append(10)
				l.Append(20)
				l.Append(5)
			},
			expected: []int{10, 20, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l := new()

			testCase.operations(l)

			testutils.AssertEquals(t, testCase.expected, l.toSlice())
		})
	}
}

func TestPrepend(t *testing.T) {
	testCases := []struct {
		name       string
		operations func(l *ArrayList[int])
		expected   []int
	}{
		{
			name: "PrependOne",
			operations: func(l *ArrayList[int]) {
				l.Prepend(20)
			},
			expected: []int{20},
		},
		{
			name: "PrependTwo",
			operations: func(l *ArrayList[int]) {
				l.Prepend(10)
				l.Prepend(20)
			},
			expected: []int{20, 10},
		},
		{
			name: "PrependThree",
			operations: func(l *ArrayList[int]) {
				l.Prepend(10)
				l.Prepend(20)
				l.Prepend(5)
			},
			expected: []int{5, 20, 10},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l := new()

			testCase.operations(l)

			testutils.AssertEquals(t, testCase.expected, l.toSlice())
		})
	}
}
