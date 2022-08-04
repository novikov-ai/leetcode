package two_sum

/*
runtime: 9 ms
memory: 5.6 MB
*/
func TwoSum(nums []int, target int) []int {
	numbers := make(map[int]int)

	for i, num := range nums {
		numbers[target-num] = i
	}

	for i, num := range nums {
		index, ok := numbers[num]
		if !ok || index == i {
			continue
		}

		if i < index {
			return []int{i, index}
		} else {
			return []int{index, i}
		}
	}

	return nil
}

/*
runtime: 63 ms
memory: 3.5 MB
*/
func TwoSumSlow(nums []int, target int) []int {

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

/*
WORKS ONLY FOR SORTED ARRAY(!)

pattern: two pointers
func TwoSum(nums []int, target int) []int {
	start, end := 0, len(nums)-1 // 0 2
	for start < end {
		sum := nums[start] + nums[end] //6
		if sum == target {
			return []int{start, end}
		}
		if sum < target {
			start++
		} else {
			end--
		}
	}

	return []int{start, end}
}

WORKS ONLY FOR SORTED ARRAY(!)
*/
