package merge_two_sorted_lists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		root1    []int
		root2    []int
		expected []int
	}{
		{root1: []int{1, 2, 4}, root2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		{root1: []int{1}, root2: []int{1, 3, 4}, expected: []int{1, 1, 3, 4}},
		{root1: []int{2}, root2: []int{4}, expected: []int{2, 4}},
	}

	for _, input := range testCases {
		input := input

		t.Run(fmt.Sprintf("%v + %v = %v", input.root1, input.root2, input.expected), func(t *testing.T) {
			root1 := getLinkedList(input.root1...)
			root2 := getLinkedList(input.root2...)

			actual := MergeTwoLists(root1, root2)

			require.Equal(t, input.expected, unwrapLinkedList(actual))
		})
	}
}

func getLinkedList(numbers ...int) *ListNode {
	root := ListNode{Val: numbers[0] % 10}
	numbers = numbers[1:]

	for _, v := range numbers {
		root.AddNode(v)
	}

	return &root
}

func (n *ListNode) AddNode(value int) {
	tail := &n

	for {
		if (*tail).Next == nil {
			break
		}
		tail = &(*tail).Next
	}

	newNode := ListNode{Val: value}
	(*tail).Next = &newNode
}

func unwrapLinkedList(node *ListNode) []int {
	result := make([]int, 0)

	current := node
	for {
		if current == nil {
			break
		}
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
