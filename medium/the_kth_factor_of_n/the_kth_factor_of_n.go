package the_kth_factor_of_n

/*
runtime: 2 ms - 54.31%
memory: 2.0 MB - 84.48%
*/

func kthFactor(n int, k int) int {
	count := 1
	for i := 1; i <= n; i++ {
		if n%i != 0 {
			continue
		}

		if count == k {
			return i
		}
		count++
	}

	return -1
}
