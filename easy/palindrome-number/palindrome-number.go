package palindrome_number

func IsPalindrome(x int) bool {
	tempX, reversedX := x, 0

	for tempX > 0 {
		lastNumber := tempX % 10
		reversedX = reversedX*10 + lastNumber
		tempX /= 10
	}

	return x == reversedX
}
