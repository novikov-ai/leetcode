package reverse_linked_list

import (
	"errors"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	nodeValues := [][]int{
		{1, 2, 3, 4, 5},
		{5},
	}

	type testCase struct {
		node         *ListNode
		reversedNode *ListNode
		left, right  int
	}

	testCases := make([]testCase, len(nodeValues))

	for _, values := range nodeValues {
		var head *ListNode
		for _, n := range values {
			node := ListNode{Val: n}
			if head == nil {
				head = &node
				continue
			}
			head.appendNode(&node)
		}
		testCases = append(testCases, testCase{node: head})
		errors.Is
	}
}

/*

Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
Example 2:

Input: head = [5], left = 1, right = 1
Output: [5]
*/
