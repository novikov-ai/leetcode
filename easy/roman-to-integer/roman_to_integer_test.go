package roman_to_integer

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := map[string]int{
		"III":     3,
		"LVIII":   58,
		"CXLIX":   149,
		"MCMXCIV": 1994,
		"MMV":     2005,
	}

	for romanNum, result := range testCases {
		romanNum := romanNum
		result := result

		t.Run(fmt.Sprintf("%s => %v", romanNum, result), func(t *testing.T) {
			number := RomanToInt(romanNum)
			require.Equal(t, result, number)
		})
	}
}
