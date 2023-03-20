package length_of_last_word

/*
runtime: 0 ms
memory: 6.5 MB
*/

func LengthOfLastWord(s string) int {
	words := make([]string, 0, len(s))
	word := ""

	for _, ch := range s {
		if string(ch) == " " {
			if word != "" {
				words = append(words, word)
				word = ""
			}
			continue
		}

		word += string(ch)
	}

	if word != "" {
		return len(word)
	}

	return len(words[len(words)-1])
}
