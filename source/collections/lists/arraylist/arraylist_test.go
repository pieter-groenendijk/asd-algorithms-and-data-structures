package arraylist

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name            string
		initialCapacity int
		expected        *ArrayList[int]
	}

	cases := []testCase{
		{
			"initialCapacity of 0",
			0,
			&ArrayList[int]{
				space:  []int{},
				length: 0,
			},
		},
		{
			"initialCapacity of 5",
			5,
			&ArrayList[int]{
				space:  []int{0, 0, 0, 0, 0},
				length: 0,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			list := New[int](test.initialCapacity)

			testutils.AssertEquals(t, list, test.expected)
		})
	}
}
