# Intuition
- The most interesting thing is **splitting** apart.
- Avoid index out of range!

# Approach
- Handle corner case with only 1 word presented.
- Iterate through a bytes array, but keep in mind, that you should cast an indexed value to string if you'd like to compare with another string.
- Once you found a space symbol just compare a prev letter with a next letter.

# Complexity
- Time complexity:
O(n) -> due to iterating through the all length

- Space complexity:
O(1) -> due to absence of any variables

# Code
```
func trimMean(arr []int) float64 {
    sort.Ints(arr)

    w := int(0.05 * float64(len(arr)))

    sum := 0.00
    for i := w; i < len(arr) - w; i++{
        sum += float64(arr[i])
    }
    
    return sum / float64(len(arr) - w * 2)
}
```