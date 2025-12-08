package arraylist

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/lists"
)

var initList = func() lists.List[int] {
	return New[int](10)
}

func TestGet(t *testing.T) {
	lists.TestGet(t, initList)
}

func TestRemove(t *testing.T) {
	lists.TestRemove(t, initList)
}

func TestRemoveAt(t *testing.T) {
	lists.TestRemoveAt(t, initList)
}

func TestSize(t *testing.T) {
	lists.TestSize(t, initList)
}

func TestSetAt(t *testing.T) {
	lists.TestSetAt(t, initList)
}

func TestAppend(t *testing.T) {
	lists.TestAppend(t, initList)
}

func TestPrepend(t *testing.T) {
	lists.TestPrepend(t, initList)
}
