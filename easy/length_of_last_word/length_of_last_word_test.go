package length_of_last_word

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		input  string
		result int
	}{
		{input: "Hello World", result: 5},
		{input: "   fly me   to   the moon  ", result: 4},
		{input: "luffy is still joyboy", result: 6},
	}

	for _, test := range testCases {
		test := test

		t.Run(fmt.Sprintf("input: %v | result: %v", test.input, test.result), func(t *testing.T) {
			result := LengthOfLastWord(test.input)
			require.Equal(t, test.result, result)
		})
	}
}
