package main

import "testing"

func toListNode(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	var head *ListNode
	for i := len(input) - 1; i >= 0; i-- {
		node := &ListNode{Val: input[i], Next: head}
		head = node
	}
	return head
}

func compareLists(list1, list2 *ListNode) bool {
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return (p1 == nil) && (p2 == nil)
}

func Test(t *testing.T) {
	testcases := []struct {
		description string
		input       []*ListNode
		expect      *ListNode
	}{
		{description: "basic", input: []*ListNode{toListNode([]int{1, 4, 5}), toListNode([]int{1, 3, 4}), toListNode([]int{2, 6})}, expect: toListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{description: "basic", input: []*ListNode{}, expect: toListNode([]int{})},
		{description: "basic", input: []*ListNode{toListNode([]int{})}, expect: toListNode([]int{})},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			output := mergeKLists(testcase.input)
			if !compareLists(output, testcase.expect) {
				t.Errorf("Failed to merge lists")
			}
		})
	}
}
