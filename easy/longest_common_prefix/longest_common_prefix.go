package two_sum

/*
runtime: 0 ms
memory: 2.4 MB
*/
func LongestCommonPrefix(strs []string) string {

	commonPrefix := strs[0]
	for i := 0; i < len(strs); i++ {
		if commonPrefix == "" {
			return ""
		}

		if len(strs[i]) < len(commonPrefix) {
			commonPrefix = commonPrefix[:len(strs[i])]
		}

		for j := 0; j < len(commonPrefix); j++ {
			if commonPrefix[j] == strs[i][j] {
				continue
			}
			if j == 0 {
				return ""
			}

			commonPrefix = commonPrefix[:j]
			break
		}
	}

	return commonPrefix
}
