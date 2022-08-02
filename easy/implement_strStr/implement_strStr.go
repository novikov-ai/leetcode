package implement_strStr

/*
runtime: 3 ms
memory: 2 MB
*/
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	index := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[index] {
			i -= index
			index = 0
			continue
		}

		if index == len(needle)-1 {
			return i - index
		}

		index++
	}

	return -1
}

/*
runtime: 4 ms
memory: 1.9 MB
*/
func StrStrSlidingWindow(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
