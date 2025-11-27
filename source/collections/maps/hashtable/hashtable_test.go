package hashtable

import (
	"testing"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name string
		expectReturn *HashTable[string, int]
	}

	cases := []testCase{
		{
			name: "",
			expectReturn: &HashTable[string, int]{
				buckets: make([]*Bucket[int], 0, 10),
				size: 0,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			gotReturn := New[string, int]()

			testutils.AssertEquals(t, test.expectReturn.buckets, gotReturn.buckets)
			testutils.AssertEquals(t, test.expectReturn.size, gotReturn.size)
			testutils.AssertNotNil(t, gotReturn.hash, "hash")
		})
	}
}
