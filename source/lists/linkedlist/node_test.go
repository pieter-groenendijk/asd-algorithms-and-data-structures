package linkedlist

import "testing"

func TestNewNode(t *testing.T) {
	t.Run("NewNode", func(t *testing.T) {
		node := newDummyNode[int]()

		if node == nil {
			t.Errorf("Node shouldn't be nil")
		}

		if node.next != nil {
			t.Errorf("node.next should be nil")
		}
	})
}

func TestNewValueNode(t *testing.T) {
	t.Run("NewValueNode", func(t *testing.T) {
		value := 5

		node := newNode(value)

		if node == nil {
			t.Errorf("Node shouldn't be nil")
		}

		if node.value != value {
			t.Errorf("Node.value, expected: %d, got %d", value, node.value)
		}
	})
}
