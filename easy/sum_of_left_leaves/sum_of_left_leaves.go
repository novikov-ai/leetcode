package sum_of_left_leaves

/*
runtime: 2 ms
memory: 2.68 MB
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	if root.Left != nil {
		leftNode := root.Left
		if leftNode.Left == nil && leftNode.Right == nil {
			sum += leftNode.Val
		} else {
			sum += SumOfLeftLeaves(leftNode)
		}
	}

	if root.Right != nil {
		sum += SumOfLeftLeaves(root.Right)
	}

	return sum
}
