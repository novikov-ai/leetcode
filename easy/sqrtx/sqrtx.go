package sqrtx

/*
runtime: 19 ms
memory: 2.16 MB
*/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	maxValue := 1
	n := 31 // limitation ((x^31)-1)

	for n > 0 {
		maxValue *= 2
		n--
	}

	maxValue -= 1

	// every row of numbers before next flat sqr is N+2
	// [1 2 3] [4 5 6 7 8] [9 10 11 12 13 14 15] [16 17 18 19 20 21 22 23 24] [25 ... 35]
	//    3  		5				7							9					11

	magic := 3
	magicIncr := 3

	prev := 1

	for i := 1; i < maxValue; i++ {
		v := i * i
		magicIncr--

		if v == x {
			return i
		}

		if v > x {
			break
		}

		prev = i

		if magicIncr == 0 {
			magic += 2
			magicIncr = magic
		}
	}

	return prev
}
