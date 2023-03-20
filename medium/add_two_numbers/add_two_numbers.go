package add_two_numbers

/*
runtime: 22 ms
memory: 4.7 MB
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rootSum, prevNode *ListNode

	currentL1, currentL2 := l1, l2
	val1, val2 := currentL1.Val, currentL2.Val

	carry := 0
	setCarry := false

	validIteration := true
	for validIteration {
		sum := val1 + val2 + carry

		if setCarry {
			carry = 0
			setCarry = false
		}

		if sum > 9 {
			carry = sum / 10
			setCarry = true
			sum %= 10
		}

		newNode := ListNode{Val: sum}
		if rootSum == nil {
			rootSum = &newNode
		} else {
			prevNode.Next = &newNode
		}
		prevNode = &newNode

		validIteration = false
		if currentL1.Next == nil {
			val1 = 0
		} else {
			currentL1 = currentL1.Next
			val1 = currentL1.Val
			validIteration = true
		}

		if currentL2.Next == nil {
			val2 = 0
		} else {
			currentL2 = currentL2.Next
			val2 = currentL2.Val
			validIteration = true
		}
	}

	if carry > 0 {
		prevNode.Next = &ListNode{Val: carry}
	}

	return rootSum
}
