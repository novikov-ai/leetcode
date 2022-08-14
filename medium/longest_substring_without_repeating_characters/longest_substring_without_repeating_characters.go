package longest_substring_without_repeating_characters

/*
runtime: 27 ms
memory: 4.6 MB
*/
func LengthOfLongestSubstring(s string) int {
	uniqueLen := 0
	sub := make([]string, 0, len(s))

	for i := 0; i < len(s); i++ {
		ch := string(s[i])

		same := false
		for j := 0; j < len(sub); j++ {
			if sub[j] == ch {
				same = true
				break
			}
		}

		if len(sub) > uniqueLen {
			uniqueLen = len(sub)
		}

		sub = append(sub, ch)

		if same {
			sub = sub[1 : len(sub)-1]
			i--
		}
	}

	if len(sub) > uniqueLen {
		uniqueLen = len(sub)
	}

	return uniqueLen
}
