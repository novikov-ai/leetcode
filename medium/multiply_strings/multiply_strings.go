package multiplystrings

import "fmt"

func multiply(num1 string, num2 string) string {
	n1, n2 := getNumber(num1), getNumber(num2)

	result := n1 * n2

	return fmt.Sprintf("%v", result)
}

func getNumber(n string) int {
	switch n {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	default:
		return -1
	}
}
