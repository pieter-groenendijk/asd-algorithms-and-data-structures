package linkedlist

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestAll(t *testing.T) {
	type testCase struct {
		name string
		values []int
	}

	testCases := []testCase{
		{
			name: "Empty",
			values: []int{},
		},
		{
			name: "OneValue",
			values: []int{5},
		},
		{
			name: "TwoValues",
			values: []int{5, 8},
		},
		{
			name: "ThreeValues",
			values: []int{5, 8, 3},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			values := test.values

			list := New[int]()
			for _, value := range values {
				list.Append(NewBasicNode(value, nil))
			}

			i := 0
			for gotValue := range list.All() {
				testutils.AssertEquals(t, values[i], gotValue)
				i++
			}

			valuesLen := len(values)
			if i != valuesLen {
				t.Errorf("Expected all values to be iterated, length: %d, length iterated: %d", valuesLen, i)
			}
		})
	}
}
