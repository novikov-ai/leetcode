package valid_parentheses

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := map[string]bool{
		"()":           true,
		"()[]{}":       true,
		"(]":           false,
		"(":            false,
		")":            false,
		"[{(([{}]))}]": true,
		"[{(([{]}))}]": false,
	}

	for testCase, result := range testCases {
		testCase := testCase
		expected := result

		t.Run(fmt.Sprintf("%s => %v", testCase, expected), func(t *testing.T) {
			isValid := IsValid(testCase)
			require.Equal(t, expected, isValid)
		})
	}
}
