package add_two_numbers

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	testCases := []struct {
		root1       []int
		root2       []int
		expectedSum []int
	}{
		{root1: []int{2, 4, 3},
			root2: []int{5, 6, 4}, expectedSum: []int{7, 0, 8}},
		{root1: []int{0, 4, 3},
			root2: []int{0, 6, 4}, expectedSum: []int{0, 0, 8}},
		{root1: []int{1, 0, 0, 0, 0, 1},
			root2: []int{5, 6, 4}, expectedSum: []int{6, 6, 4, 0, 0, 1}},
		{root1: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			root2: []int{5, 6, 4}, expectedSum: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, input := range testCases {
		input := input

		t.Run(fmt.Sprintf("%v + %v = %v", input.root1, input.root2, input.expectedSum), func(t *testing.T) {
			root1 := getLinkedList(input.root1...)
			root2 := getLinkedList(input.root2...)
			actualSum := AddTwoNumbers(root1, root2)

			require.Equal(t, input.expectedSum, unwrapLinkedList(actualSum))
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
