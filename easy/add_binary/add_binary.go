package add_binary

import "fmt"

/*
runtime: 0 ms
memory: 4,54 MB
*/

func addBinary(a string, b string) string {
	// 0 + 0 = 0
	// 1 + 1 = 0 (1)
	// 1 + 0 = 1

	biggest, smallest := "", ""

	if len(a) >= len(b) {
		biggest = a
		smallest = b
	} else {
		biggest = b
		smallest = a
	}

	if diff := len(biggest) - len(smallest); diff > 0 {
		for i := 0; i < diff; i++ {
			smallest = "0" + smallest
		}
	}

	sum := ""
	carry := false
	for i := len(biggest) - 1; i >= 0; i-- {
		if biggest[i] == smallest[i] {
			if string(biggest[i]) == "1" {
				// a=b=1
				if carry {
					// 1+1+1
					sum = "1" + sum
					carry = true
				} else {
					sum = "0" + sum
					carry = true
				}
			} else {
				// a=b=0
				if carry {
					sum = "1" + sum
					carry = false
				} else {
					sum = "0" + sum
				}
			}
		} else {
			// a<>b=(0/1)
			if carry {
				sum = "0" + sum
				carry = true
			} else {
				sum = "1" + sum
			}
		}
	}

	if carry {
		sum = "1" + sum
	}

	return sum
}

/*
SOLUTION BELOW PROVIDED BY OPEN AI
runtime: 2 ms
memory: 3,74 MB
*/

func addBinaryV2(a string, b string) string {
	result := ""
	carry := 0

	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = fmt.Sprintf("%d%s", sum%2, result)
		carry = sum / 2
	}

	return result
}
