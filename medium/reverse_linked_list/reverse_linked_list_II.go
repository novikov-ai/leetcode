package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) appendNode(node *ListNode) {
	currentNode := head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = node
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}

	var prevRev, postRev *ListNode
	var firstRev, lastRev *ListNode

	current := head

	position := 1
	for {
		if position == left {
			prevRev = prev
			firstRev = current
		}
		if position > left && position < right {
			current.Next = prev
		}
		if position == right {
			postRev = prev
			lastRev = current
		}

		prev = current
		current = current.Next
		position++

		if current == nil {
			lastRev.Next = prevRev
			firstRev.Next = postRev
			break
		}
	}

	return head
}

/*
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head.Next == nil {
		return head
	}

    left -= 1
    right -= 1

	linkedList := make([]*ListNode,0,500)

    current := head
    for {
        linkedList = append(linkedList, current)
        current = current.Next
        if current == nil{
            break
        }
    }

    if len(linkedList) == 2{
        if left < right{
            return head
        }

        node := head.Next
        head.Next = nil
        return node
    }

    prevRev := linkedList[left]
    postRev := linkedList[right+1]

    reversed := linkedList[left : right+1] // len > 2
    for i := len(reversed) - 1; i >= 0; i-- {
		if i > 0 {
			reversed[i].Next = reversed[i-1]
			continue
		}
		reversed[i].Next = postRev
	}

    prevRev.Next = reversed[len(reversed)-1]

    if left == 0{
        return reversed[len(reversed) - 1]
    }

    return head
}
*/

//https://leetcode.com/problems/reverse-linked-list-ii/
