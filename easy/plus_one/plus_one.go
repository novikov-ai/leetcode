package plus_one

/*
runtime: 0 ms
memory: 2.1 MB
*/

func PlusOne(digits []int) []int {
	digitsPlusOne := make([]int, len(digits)+1)

	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if sum > 9 {
			carry = sum - 9
			sum %= 10
		} else {
			carry = 0
		}

		digitsPlusOne[i+1] = sum
	}

	digitsPlusOne[0] = carry
	if digitsPlusOne[0] == 0 {
		return digitsPlusOne[1:]
	}

	return digitsPlusOne
}
