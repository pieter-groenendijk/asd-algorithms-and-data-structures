package hashtable

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func Test(t *testing.T) {
	hashTable := New[string, int]()

	hashTable.Get("any")
	value, exists := hashTable.Get("any")
	testutils.AssertEquals(t, false, exists)
	// We don't care about the value here

	hashTable.Set("horse", 10)
	value, exists = hashTable.Get("horse")
	testutils.AssertEquals(t, true, exists)
	testutils.AssertEquals(t, 10, value)

	hashTable.Set("blue", 1234)
	value, exists = hashTable.Get("blue")
	testutils.AssertEquals(t, true, exists)
	testutils.AssertEquals(t, 1234, value)
	value, exists = hashTable.Get("horse")
	testutils.AssertEquals(t, true, exists)
	testutils.AssertEquals(t, 10, value)

	hashTable.Set("horse", 5)
	value, exists = hashTable.Get("horse")
	testutils.AssertEquals(t, true, exists)
	testutils.AssertEquals(t, 5, value)

	hashTable.Unset("horse")
	value, exists = hashTable.Get("horse")
	testutils.AssertEquals(t, false, exists)
	// We don't care about the value here
}
