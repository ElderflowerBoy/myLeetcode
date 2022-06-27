package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}))
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return node
}
