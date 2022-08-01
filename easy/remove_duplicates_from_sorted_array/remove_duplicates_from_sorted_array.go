package remove_duplicates_from_sorted_array

/*
runtime: 14 ms
memory: 4.4 MB
*/
func RemoveDuplicates(nums []int) int {
	lag := 0
	for i, n := range nums {
		if i == 0 || n == nums[i-1] {
			continue
		}

		lag++
		nums[lag] = n
	}

	return lag + 1
}
