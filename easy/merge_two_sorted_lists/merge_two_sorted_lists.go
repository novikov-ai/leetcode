package merge_two_sorted_lists

/*
runtime: 2 ms
memory: 2.7 MB
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var mergedHead *ListNode

	if list1 == nil && list2 == nil {
		return mergedHead
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	mergedHead, list1, list2 = getNextMinNode(list1, list2)
	current := mergedHead

	for {
		if list1 == nil {
			current.Next = list2
			break
		}
		if list2 == nil {
			current.Next = list1
			break
		}

		current.Next, list1, list2 = getNextMinNode(list1, list2)
		current = current.Next
	}

	return mergedHead
}

func getNextMinNode(list1, list2 *ListNode) (current, l1, l2 *ListNode) {
	current = new(ListNode)
	if list1.Val < list2.Val {
		current = &ListNode{Val: list1.Val}
		list1 = list1.Next
	} else {
		current = &ListNode{Val: list2.Val}
		list2 = list2.Next
	}

	return current, list1, list2
}
