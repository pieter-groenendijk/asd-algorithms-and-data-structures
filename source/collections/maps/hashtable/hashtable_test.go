package hashtable

import (
	"testing"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestNew(t *testing.T) {
	t.Run("", func(t *testing.T) {
		got := New[string, int]()

		testutils.AssertNotNil(t, got, "")
	})
}
