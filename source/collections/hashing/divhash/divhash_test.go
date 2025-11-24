package divhash

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/hashing"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestHash(t *testing.T) {
	type testCase struct {
		name string
		paramValue uint32 
		paramSize uint32 
		expectResult uint32 
		expectErr error
	}

	cases := []testCase{
		{
			name: "firstOfNaturalRange",
			paramValue: 0,
			paramSize: 10,
			expectResult: 0,
			expectErr: nil,
		},
		{
			name: "lastOfNaturalRange",
			paramValue: 9,
			paramSize: 10,
			expectResult: 9,
			expectErr: nil,
		},
		{
		 	name: "overNaturalRange",
			paramValue: 10,
			paramSize: 10,
			expectResult: 0,
		},
		{
			name: "smallestRoom",
			paramValue: 1512,
			paramSize: 1,
			expectResult: 0,
			expectErr: nil,
		},
		{
			name: "zeroRange",
			paramValue: 0,
			paramSize: 0,
			expectResult: 0,
			expectErr: hashing.ErrNoHashRoom,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			gotResult, gotErr := Hash(test.paramValue, test.paramSize)
			
			testutils.AssertEquals(t, test.expectResult, gotResult)
			testutils.AssertEquals(t, test.expectErr, gotErr)
		})
	}
}
