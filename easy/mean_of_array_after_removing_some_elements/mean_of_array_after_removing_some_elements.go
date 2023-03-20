package mean_of_array_after_removing_some_elements

import "sort"

func TrimMean(arr []int) float64 {
	sort.Ints(arr)

	w := int(0.05 * float64(len(arr)))

	sum := 0.00
	for i := w; i < len(arr)-w; i++ {
		sum += float64(arr[i])
	}

	return sum / float64(len(arr)-w*2)
}
