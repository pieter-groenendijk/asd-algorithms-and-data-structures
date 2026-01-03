package binsearch

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func intPointer(i int) *int {
	return &i
}

func TestSearch(t *testing.T) {
	testCases := []struct{
		name string
		giveValues []int
		giveValue int
		expectIndexIn []int 
		expectErr error
	}{
		{
			name: "valueNotInZeroValues",
			giveValues: []int{},
			giveValue: 5,
			expectIndexIn: nil,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "valueNotInOneValue",
			giveValues: []int{8},
			giveValue: 5,
			expectIndexIn: nil,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "valueNotInTwoValues",
			giveValues: []int{5,7},
			giveValue: 6,
			expectIndexIn: nil,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "valueNotInThreeValues",
			giveValues: []int{5,7,9},
			giveValue: 8,
			expectIndexIn: nil,
			expectErr: collections.ErrNotFound,
		},
		{
			name: "valueNotInDuplicateValues",
			giveValues: []int{5,5,5},
			giveValue: 8,
			expectIndexIn: nil,
			expectErr: collections.ErrNotFound,
		},

		{
			name: "valueInOneValue",
			giveValues: []int{5},
			giveValue: 5,
			expectIndexIn: []int{0},
			expectErr: nil,
		},
		{
			name: "valueInTwoValuesAtStart",
			giveValues: []int{5,7},
			giveValue: 5,
			expectIndexIn: []int{0},
			expectErr: nil,
		},
		{
			name: "valueInTwoValuesAtEnd",
			giveValues: []int{6,8},
			giveValue: 8,
			expectIndexIn: []int{1},
			expectErr: nil,
		},
		{
			name: "valueInThreeValuesAtStart",
			giveValues: []int{3,4,8},
			giveValue: 3,
			expectIndexIn: []int{0},
			expectErr: nil,
		},
		{
			name: "valueInThreeValuesAtEnd",
			giveValues: []int{4,5,8},
			giveValue: 8,
			expectIndexIn: []int{2},
			expectErr: nil,
		},
		{
			name: "valueInThreeValuesAtMiddle",
			giveValues: []int{3,4,5},
			giveValue: 4,
			expectIndexIn: []int{1},
			expectErr: nil,
		},
		{
			name: "valueInDuplicateValues",
			giveValues: []int{3,3,4,4,4,4,7,8},
			giveValue: 4,
			expectIndexIn: []int{2,3,4},
			expectErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			gotIndex, gotErr := Search(testCase.giveValues, testCase.giveValue)

			if testCase.expectIndexIn != nil {
				gotIndexIn := false
				for _, expectIndex := range testCase.expectIndexIn {
					if expectIndex == gotIndex {
						gotIndexIn = true
					}
				}

				testutils.AssertEquals(t, gotIndexIn, true)
			}

			testutils.AssertEquals(t, testCase.expectErr, gotErr)
		})
	}
}

func BenchSearch(b *testing.B) {
	b.Run("smallLength", func(b *testing.B) {
		
	})

	b.Run("mediumLength", func(b *testing.B) {

	})

	b.Run("largeLength", func(b *testing.B) {
	
	})

	b.Run("extremeLength", func(b *testing.B) {

	})
}
