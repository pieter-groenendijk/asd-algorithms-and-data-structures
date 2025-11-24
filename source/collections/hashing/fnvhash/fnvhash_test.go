package fnvhash

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
			name: "zero",
			paramValue: 0,
			paramSize: 100,
			expectResult: 5, // from 1268118805
			expectErr: nil,
		},
		{
			name: "one",
			paramValue: 1,
			paramSize: 100,
			expectResult: 92, // from 4218009092
			expectErr: nil,
		},
		{
			name: "two",
			paramValue: 2,
			paramSize: 100,
			expectResult: 23, // from 3958272823
			expectErr: nil,
		},
		{
			name: "three",
			paramValue: 3,
			paramSize: 100,
			expectResult: 14, // from 2613195814
			expectErr: nil,
		},
		{
			name: "maxValue",
			paramValue: 4294967295,
			paramSize: 100,
			expectResult: 41, // from 3809873841
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
			// value := test.paramValue

			// hasher := fnv.New32a()
			// hasher.Write([]byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24)})

			// hashValue := hasher.Sum32()
			// var err error = nil

			// hashValue, err = divhash.Hash(hashValue, test.paramSize)

			gotResult, gotErr := Hash(test.paramValue, test.paramSize)
			
			testutils.AssertEquals(t, test.expectResult, gotResult)
			testutils.AssertEquals(t, test.expectErr, gotErr)

			// testutils.AssertEquals(t, test.expectResult, hashValue)
			// testutils.AssertEquals(t, test.expectErr, err)
		})
	}
}
