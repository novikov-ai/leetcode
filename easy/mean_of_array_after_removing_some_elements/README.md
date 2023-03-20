# Intuition
- The most interesting thing is **sorting**.
- I was thinking to implement sorting by myself, but if Leetcode could accept an answer with stlib func, then I will use it.

# Approach
- Sort first.
- You should do nothing with the array, just skip Min and Max 5% values.
- Return mean, but beware of types!

# Complexity
- Time complexity:
O(n*log(n)) -> due to sorting func

- Space complexity:
O(n) -> due to an input slice

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