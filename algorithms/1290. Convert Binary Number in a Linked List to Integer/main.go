package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	println(getDecimalValue(getListNode(1101)))
}

func getDecimalValue(head *ListNode) int {

	result := 0
	var number []int

	for head != nil {
		number = append(number, head.Val)
		head = head.Next
	}

	lenBinary := len(number)

	for i := 0; i < lenBinary; i++ {
		result += number[i] * int(math.Pow(2, float64(lenBinary-1-i)))
	}

	return result
}

func getListNode(binary int) *ListNode {
	head := &ListNode{}
	node := head
	x := 0
	for binary != 0 {
		x *= 10
		x += binary % 10
		binary /= 10
	}

	for x != 0 {
		node.Next = &ListNode{Val: x % 10}
		node = node.Next
		x /= 10
	}
	return head.Next
}
