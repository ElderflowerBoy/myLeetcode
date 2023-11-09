package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) add(val int) {
	var tail = node

	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = &ListNode{Val: val}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}
	curr := res

	for l1 != nil || l2 != nil {

		if l1 != nil {
			curr.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			curr.Val += l2.Val
			l2 = l2.Next
		}

		if curr.Val > 9 {
			curr.Val -= 10
			curr.Next = &ListNode{Val: 1}
			curr = curr.Next
		} else if l1 != nil || l2 != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
		}

	}

	return res
}

func main() {
	var num1 = &ListNode{Val: 8}
	var num2 = &ListNode{Val: 5}
	num1.add(8)
	num1.add(6)
	num1.add(5)
	num1.add(4)

	num2.add(5)
	num2.add(4)
	num2.add(3)
	num2.add(2)

	addTwoNumbers(num1, num2)

}
