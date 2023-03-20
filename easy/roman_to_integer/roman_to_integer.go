package roman_to_integer

/*
runtime: 7 ms
memory: 3.8 MB
*/

func RomanToInt(s string) int {

	romanNums := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	constrNums := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	stack := make([]int, len(s))

	cacheStr := ""
	round := 0

	for _, b := range s {
		num, ok := romanNums[string(b)]
		if !ok {
			return -1
		}

		round++

		if round == 1 {
			stack = append(stack, num)
			cacheStr = string(b)
			continue
		}

		constr := cacheStr + string(b)
		numConstr, ok := constrNums[constr]
		if ok {
			stack = stack[:len(stack)-1]
			stack = append(stack, numConstr)
			round = 0
		} else {
			stack = append(stack, num)
			cacheStr = string(b)
			round = 1
		}
	}

	sum := 0
	for _, v := range stack {
		sum += v
	}

	return sum
}
