package palindrome_number

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x > 0 && x < 10 {
		return true
	}

	countLenX := 0
	multiplier := 1
	valueX := x

	for valueX > 0 {
		valueX /= 10
		countLenX++

		multiplier *= 10
	}

	multiplier /= 10

	for i := 0; i < countLenX/2; i++ {
		first := (x / multiplier) % 10
		last := x % 10

		if first != last {
			return false
		}

		x /= 10
		multiplier /= 100
	}

	return true
}
