package linkedlistv2

import (
	"math"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func new() *LinkedList[int] {
	return New(func(thisValue int, thatValue int) bool {
		return thisValue == thatValue
	})
}

// helper method to get an easy conceptual view of a list
func (l *LinkedList[TValue]) toSlice() []TValue {
	slice := make([]TValue, 0, 8)

	currNode := l.head.next
	for currNode != nil {
		slice = append(slice, currNode.value)

		currNode = currNode.next
	}

	return slice
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name       string
		operations func(l *LinkedList[int])
		expected   []int
	}{
		{
			name: "RemoveFirstElement",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.Remove(20)
			},
			expected: []int{123, 3},
		},
		{
			name: "RemoveMiddleElement",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.Remove(123)
			},
			expected: []int{20, 3},
		},
		{
			name: "RemoveLastElement",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.Remove(3)
			},
			expected: []int{20, 123},
		},
		{
			name: "RemoveNonExistent",
			operations: func(l *LinkedList[int]) {
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
		operations func(l *LinkedList[int])
		expected   []int
	}{
		{
			name: "RemoveFirstElement",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(0)
			},
			expected: []int{123, 3},
		},
		{
			name: "RemoveMiddleElement",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(1)
			},
			expected: []int{20, 3},
		},
		{
			name: "RemoveLastElement",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(2)
			},
			expected: []int{20, 123},
		},
		{
			name: "RemoveBeforeBounds",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
				l.Append(123)
				l.Append(3)

				l.RemoveAt(-1)
			},
			expected: []int{20, 123, 3},
		},
		{
			name: "RemoveAfterBounds",
			operations: func(l *LinkedList[int]) {
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
		operations func(l *LinkedList[int]) []returnValue
		expected   []returnValue
	}{
		{
			name: "GetFirst",
			operations: func(l *LinkedList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetFrom(0)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{10, true}},
		},
		{
			name: "GetMiddle",
			operations: func(l *LinkedList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetFrom(1)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{20, true}},
		},
		{
			name: "GetLast",
			operations: func(l *LinkedList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetFrom(2)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{5, true}},
		},
		{
			name: "GetFrom",
			operations: func(l *LinkedList[int]) []returnValue {
				l.Append(10)

				valueOne, existsOne := l.GetFrom(0)

				l.Append(20)

				valueTwo, existsTwo := l.GetFrom(0)

				l.Append(123)

				valueThree, existsThree := l.GetFrom(0)

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
			operations: func(l *LinkedList[int]) []returnValue {
				l.Append(10)
				l.Append(20)
				l.Append(5)

				value, exists := l.GetFrom(-1)

				valueTwo, existsTwo := l.GetFrom(-5)

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
			operations: func(l *LinkedList[int]) []returnValue {
				value, exists := l.GetFrom(0)
				return []returnValue{
					{value, exists},
				}
			},
			expected: []returnValue{{0, false}},
		},
		{
			name: "GetAfterBounds",
			operations: func(l *LinkedList[int]) []returnValue {
				l.Append(10)
				l.Append(20)

				value, exists := l.GetFrom(2)

				l.Append(5)

				valueTwo, existsTwo := l.GetFrom(3)

				valueThree, existsThree := l.GetFrom(math.MaxInt)

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
		operations func(l *LinkedList[int])
		expected   []int
	}{
		{
			name: "AppendOne",
			operations: func(l *LinkedList[int]) {
				l.Append(20)
			},
			expected: []int{20},
		},
		{
			name: "AppendTwo",
			operations: func(l *LinkedList[int]) {
				l.Append(10)
				l.Append(20)
			},
			expected: []int{10, 20},
		},
		{
			name: "AppendThree",
			operations: func(l *LinkedList[int]) {
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
		operations func(l *LinkedList[int])
		expected   []int
	}{
		{
			name: "PrependOne",
			operations: func(l *LinkedList[int]) {
				l.Prepend(20)
			},
			expected: []int{20},
		},
		{
			name: "PrependTwo",
			operations: func(l *LinkedList[int]) {
				l.Prepend(10)
				l.Prepend(20)
			},
			expected: []int{20, 10},
		},
		{
			name: "PrependThree",
			operations: func(l *LinkedList[int]) {
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
