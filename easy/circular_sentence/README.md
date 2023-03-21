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
func isCircularSentence(sentence string) bool {
    if sentence[0] != sentence[len(sentence) - 1]{
        return false
    }

    for i := 0; i < len(sentence); i++{
       if string(sentence[i]) != " "{
           continue
       }

       if sentence[i-1] != sentence[i+1]{
           return false
       }
    }
    return true
}
```