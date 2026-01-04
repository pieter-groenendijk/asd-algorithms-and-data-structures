package heaps

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

type HeapCreator func() Heap[int, int]

type node struct {
	order int
	value int
}

// Blackbox testing will create circular test "true" dependencies between Push and Pop, therefore
// one set of cases tests both.
func TestPushAndPop(t *testing.T, heapCreator HeapCreator) {
	testCases := []struct{
		name string
		prePushed []node
		expectValueIn []int 
		expectValueExist bool
	}{
		{
			name: "zeroElements",
			prePushed: []node{},	
			expectValueIn: []int{},
			expectValueExist: false,
		},
		{
			name: "oneElement",
			prePushed: []node{{5, 9}},
			expectValueIn: []int{9},
			expectValueExist: true,
		},
		{
			name: "twoElementsGetLast",
			prePushed: []node{{3, 4}, {5, 3}},
			expectValueIn: []int{3},
			expectValueExist: true,
		},
		{
			name: "twoElementsGetFirst",
			prePushed: []node{{8, 3}, {5, 1}},
			expectValueIn: []int{3},
			expectValueExist: true,
		},
		{
			name: "threeElementsGetFirst",
			prePushed: []node{{8, 3}, {4, 7}, {6, 4}},
			expectValueIn: []int{3},
			expectValueExist: true,
		},
		{
			name: "threeElementsGetMiddle",
			prePushed: []node{{8, 3}, {10, 7}, {6, 4}},
			expectValueIn: []int{7},
			expectValueExist: true,
		},
		{
			name: "threeElementsGetLast",
			prePushed: []node{{8, 3}, {4, 7}, {11, 4}},
			expectValueIn: []int{4},
			expectValueExist: true,
		},
		{
			name: "duplicateOrders",
			prePushed: []node{{8, 3}, {4, 7}, {8, 4}},
			expectValueIn: []int{3, 4},
			expectValueExist: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			heap := heapCreator()

			for _, nodeToPush := range testCase.prePushed {
				heap.Push(nodeToPush.order, nodeToPush.value)
			}

			gotValue, gotValueExists := heap.Pop()

			if testCase.expectValueExist {
				testutils.AssertIn(t, testCase.expectValueIn, gotValue)
			}

			testutils.AssertEquals(t, testCase.expectValueExist, gotValueExists)
		})
	}
}
