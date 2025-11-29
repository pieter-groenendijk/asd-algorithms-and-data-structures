package linkedlist

import (
	"fmt"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestGetNodeBefore(t *testing.T) {
	type testCase struct {
		// administrative
		name string
		// preconditions
		values []int	
		// conditions
		givenKey int
		// postconditions
		doExpectNode bool
		expectNextNode Node[int, int]
		doExpectErr bool
		expectErr error
	}

	testCases := []testCase{
		{
			name: "NotFoundInEmptyList",
			values: []int{},
			givenKey: 5,
			doExpectNode: false,
			doExpectErr: true,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "NotFoundInNonZeroList",
			values: []int{3,5,8},
			givenKey: 2,
			doExpectNode: false,
			doExpectErr: true,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "FoundFirst",
			values: []int{3,5,8},
			givenKey: 3,
			doExpectNode: true,
			expectNextNode: NewBasicNode(3, NewBasicNode(5, NewBasicNode(8, nil))),
			doExpectErr: false,
		},
		{
			name: "FoundMiddle",
			values: []int{3,5,8},
			givenKey: 5,
			doExpectNode: true,
			expectNextNode: NewBasicNode(5, NewBasicNode(8, nil)),
			doExpectErr: false,
		},
		{
			name: "FoundLast",
			values: []int{3,5,8},
			givenKey: 8,
			doExpectNode: true,
			expectNextNode: NewBasicNode(8, nil),
			doExpectErr: false,
		},
		{
			name: "FoundFirstOfDuplicates",
			values: []int{3,3,8},
			givenKey: 3,
			doExpectNode: true,
			expectNextNode: NewBasicNode(3, NewBasicNode(3, NewBasicNode(8, nil))),
			doExpectErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range test.values {
				list.Append(NewBasicNode(value, nil))
			}

			gotNode, gotErr := list.GetNodeBefore(test.givenKey)

			if test.doExpectNode {
				testutils.AssertEquals(t, test.expectNextNode, gotNode.Next())
			}
			if test.doExpectErr {
				testutils.AssertEquals(t, test.expectErr, gotErr)
			}
		})
	}
}

func TestGetNode(t *testing.T) {
	type testCase struct {
		// administrative
		name string
		// preconditions
		values []int	
		// conditions
		givenKey int
		// postconditions
		doExpectNode bool
		expectNode Node[int, int]
		doExpectErr bool
		expectErr error
	}

	testCases := []testCase{
		{
			name: "NotFoundInEmptyList",
			values: []int{},
			givenKey: 5,
			doExpectNode: false,
			doExpectErr: true,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "NotFoundInNonZeroList",
			values: []int{3,5,8},
			givenKey: 2,
			doExpectNode: false,
			doExpectErr: true,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "FoundFirst",
			values: []int{3,5,8},
			givenKey: 3,
			doExpectNode: true,
			expectNode: NewBasicNode(3, NewBasicNode(5, NewBasicNode(8, nil))),
			doExpectErr: false,
		},
		{
			name: "FoundMiddle",
			values: []int{3,5,8},
			givenKey: 5,
			doExpectNode: true,
			expectNode: NewBasicNode(5, NewBasicNode(8, nil)),
			doExpectErr: false,
		},
		{
			name: "FoundLast",
			values: []int{3,5,8},
			givenKey: 8,
			doExpectNode: true,
			expectNode: NewBasicNode(8, nil),
			doExpectErr: false,
		},
		{
			name: "FoundFirstOfDuplicates",
			values: []int{3,3,8},
			givenKey: 3,
			doExpectNode: true,
			expectNode: NewBasicNode(3, NewBasicNode(3, NewBasicNode(8, nil))),
			doExpectErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range test.values {
				list.Append(NewBasicNode(value, nil))
			}

			gotNode, gotErr := list.GetNode(test.givenKey)

			if test.doExpectNode {
				testutils.AssertEquals(t, test.expectNode, gotNode)
			}
			if test.doExpectErr {
				testutils.AssertEquals(t, test.expectErr, gotErr)
			}
		})
	}
}

func TestRemoveAfter(t *testing.T) {
	type testCase struct {
		name string
		values []int
		givenBeforeNodeValue int 
		expectValues []int
	}

	testCases := []testCase{
		{
			name: "NoAfterNode",
			values: []int{3,5},
			givenBeforeNodeValue: 5,
			expectValues: []int{3,5},
		},
		{
			name: "AfterNodeRemoved",
			values: []int{3,5,8},
			givenBeforeNodeValue: 3,
			expectValues: []int{3},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range test.values {
				list.Append(NewBasicNode(value, nil))
			}
			givenBeforeNode, _ := list.GetNodeBefore(test.givenBeforeNodeValue)

			fmt.Println(test.name)
			fmt.Printf("%#v\n", list.dummyHead)

			list.RemoveAfter(givenBeforeNode)

			fmt.Printf("%#v\n", list.dummyHead)

			i := 0
			for value := range list.All() {
				testutils.AssertEquals(t, test.expectValues[i], value)
				i++
			}

			valuesLen := len(test.expectValues)
			if i != valuesLen {
				t.Errorf("Expected all values to be iterated, length: %d, length iterated: %d", valuesLen, i)
			}
		})
	}
}

func TestInsertAfter(t *testing.T) {

}
