package remove_element

/*
runtime: 0 ms
memory: 2.1 MB
*/
func RemoveElement(nums []int, val int) int {
	lag := 0
	for _, n := range nums {
		if n == val {
			continue
		}

		nums[lag] = n
		lag++
	}

	return lag
}
