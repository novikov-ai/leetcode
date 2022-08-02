package implement_strStr

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		result   int
	}{
		{haystack: "a", needle: "a", result: 0},
		{haystack: "hello", needle: "ll", result: 2},
		{haystack: "aaaa", needle: "ba", result: -1},
		{haystack: "asdasddsa", needle: "", result: 0},
		{haystack: "mississippi", needle: "issip", result: 4},
	}

	for _, test := range testCases {
		test := test

		t.Run(fmt.Sprintf("haystack: %v | needle: %v | result: %v", test.haystack, test.needle, test.result), func(t *testing.T) {
			actual := StrStr(test.haystack, test.needle)
			require.Equal(t, test.result, actual)

			actualSlidingWindow := StrStrSlidingWindow(test.haystack, test.needle)
			require.Equal(t, test.result, actualSlidingWindow)
		})
	}
}
