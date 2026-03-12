package main

import (
	"math"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode

	for len(lists) > 0 {
		min := math.MaxInt32
		minIdx := -1

		i := 0
		for i < len(lists) {
			if lists[i] == nil {
				lists = slices.Delete(lists, i, i+1)
				continue
			}
			if lists[i].Val < min {
				min = lists[i].Val
				minIdx = i
			}
			i++
		}

		if minIdx > -1 {
			result = &ListNode{Val: min, Next: result}
			lists[minIdx] = lists[minIdx].Next
		}
	}

	var reversed *ListNode
	for result != nil {
		reversed = &ListNode{Val: result.Val, Next: reversed}
		result = result.Next
	}

	return reversed
}
