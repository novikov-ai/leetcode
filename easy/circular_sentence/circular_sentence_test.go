package circular_sentence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCircularSentence(t *testing.T) {
	testCases := []struct {
		in     string
		result bool
	}{
		{
			in:     "leetcode exercises sound delightful",
			result: true,
		},
		{
			in:     "eetcode",
			result: true,
		},
		{
			in:     "leetcode",
			result: false,
		},
		{
			in:     "Leetcode is cool",
			result: false,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.result, IsCircularSentence(testCase.in))
		})
	}
}
