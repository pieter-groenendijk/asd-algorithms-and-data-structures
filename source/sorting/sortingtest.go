package sorting

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestSortInPlace(t *testing.T, sortInPlaceFunc SortInPlaceFunc[int]) {
	type testCase struct {
		name string
		givenItems []int
		expectItems []int
	}

	testCases := []testCase{
		{
			name: "InverselySorted",
			givenItems: []int{8,5,3},
			expectItems: []int{3,5,8},
		},
		{
			name: "PartlySorted",
			givenItems: []int{5,3,8},
			expectItems: []int{3,5,8},
		},
		{
			name: "AlreadySorted",
			givenItems: []int{3,5,8},
			expectItems: []int{3,5,8},
		},
		{
			name: "ZeroItems",
			givenItems: []int{},
			expectItems: []int{},
		},
		{
			name: "OneItem", 
			givenItems: []int{5},
			expectItems: []int{5},
		},
		{
			name: "TwoItems",
			givenItems: []int{5,3},
			expectItems: []int{3,5},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			sortInPlaceFunc(test.givenItems)

			testutils.AssertEquals(t, test.expectItems, test.givenItems)
		})
	}
}
