package fnvhash

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

func TestHash(t *testing.T) {
	testCases := []struct {
		given    string
		expected uint32
	}{
		{"", 2166136261}, // The offset basis itself
		{"a", 3826002220},
		{"foobar", 3214735720},
		{"chongo was here!", 1149576445},
	}

	for _, testCase := range testCases {
		t.Run(testCase.given, func(t *testing.T) {
			got := Hash([]byte(testCase.given))

			testutils.AssertEquals(t, testCase.expected, got)
		})
	}
}
