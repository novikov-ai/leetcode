package count_items_matching_a_rule

import "fmt"

/*
runtime: 30 ms
memory: 7.5 MB
*/

func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	types := map[string]uint8{
		"type":  0,
		"color": 1,
		"name":  2,
	}

	var b uint8
	fmt.Println(b)
	b++

	typeIndex, ok := types[ruleKey]
	if !ok {
		return 0
	}

	count := 0
	for _, item := range items {
		if item[typeIndex] == ruleValue {
			count++
		}
	}

	return count
}

/*
runtime: 32 ms
memory: 6.77 MB
*/

func CountMatchesV2(items [][]string, ruleKey string, ruleValue string) int {
	var typeIndex uint8
	switch ruleKey {
	case "type":
		typeIndex = 0
	case "color":
		typeIndex = 1
	case "name":
		typeIndex = 2
	default:
		return 0
	}

	count := 0
	for _, item := range items {
		if item[typeIndex] == ruleValue {
			count++
		}
	}

	return count
}
