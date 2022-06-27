package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := removeElements(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
	}, 2)

	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	node := &ListNode{}
	newHead := node
	for head != nil {
		if head.Val != val {
			node.Next = &ListNode{Val: head.Val}
			node = node.Next
		}
		head = head.Next
	}
	return newHead.Next
}
