package linkedlist

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestAdd(t *testing.T) {
	t.Run("firstAdd", func(t *testing.T) {
		list := New[int]()
		value := 5

		list.Add(value)

		expectedNode := list.head.next
		testutils.AssertEquals(t, *expectedNode, *newValueNode(value))
		testutils.AssertEquals(t, uint(1), list.size)
	})

	t.Run("subsequentAdds", func(t *testing.T) {
		list := New[int]()
		value := 5

		list.Add(value)

		expectedNode := list.head.next
		testutils.AssertEquals(t, *expectedNode, *newValueNode(value))
		testutils.AssertEquals(t, uint(1), list.size)

		list.Add(value)

		expectedNode = expectedNode.next
		testutils.AssertEquals(t, *expectedNode, *newValueNode(value))
		testutils.AssertEquals(t, uint(2), list.size)
	})
}
