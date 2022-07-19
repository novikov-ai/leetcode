package palindrome_number

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := map[int]bool{
		-123:           false,
		0:              true,
		9:              true,
		10:             false,
		11:             true,
		111:            true,
		21112:          true,
		216112:         false,
		1000021:        false,
		1200021:        true,
		51231322313215: true,
		51231322713215: false,
	}

	for num, result := range testCases {
		num := num
		result := result

		t.Run(fmt.Sprintf("%v: %v", num, result), func(t *testing.T) {
			palindrome := IsPalindrome(num)
			require.Equal(t, result, palindrome)
		})
	}
}
