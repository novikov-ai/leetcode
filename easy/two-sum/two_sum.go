package two_sum

// runtime: 63 ms
// memory: 3.5 MB
func TwoSum(nums []int, target int) []int {

	for i, num := range nums {
		for j, n := range nums {
			if i == j {
				continue
			}
			if (num + n) == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
