package linkedlist

import (
	"testing"
)

func TestSomething(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		ll := New[int]()

		if ll == nil {
			t.Errorf("Linked list shouldn't be nil")
		}

		if ll.head == nil {
			t.Errorf("Linked list head shouldn't be nil")
		}

		if ll.size != 0 {
			t.Errorf("Size should be 0, but is %d", ll.size)
		}
	})
}
