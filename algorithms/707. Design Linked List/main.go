package main

func main() {
	//Your MyLinkedList object will be instantiated and called as such:
	//"addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
	//[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
	obj := Constructor()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	println(obj.Get(4))
	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)
	for obj.Head != nil {
		println(obj.Head.Val)
		obj.Head = obj.Head.Next
	}

}

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func Constructor() MyLinkedList {
	node := &Node{}
	return MyLinkedList{
		Head: node,
		Tail: node,
		Size: 0,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if l.indexNotExists(index) {
		return -1
	}
	node := l.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{Val: val}
	if l.Size == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.Size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{Val: val}
	if l.Size == 0 {
		l.Tail = newNode
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if l.indexNotExists(index) {
		return
	}

	if l.Size == index {
		l.AddAtTail(val)
		return
	}

	newNode := &Node{Val: val}

	if l.Size == 0 {
		l.Tail = newNode
		l.Head = newNode
		l.Size++
		return
	}

	head := l.Head
	for i := 0; i < index; i++ {
		head = head.Next
	}

	newNode.Next = head.Next
	head.Next = newNode
	l.Size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if l.indexNotExists(index) {
		return
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Size--
		return
	}

	node := l.Head
	for i := 1; i < index; i++ {
		node = node.Next
	}
	if node.Next != nil {
		node.Next = node.Next.Next
	}

	l.Size--
}

func (l *MyLinkedList) indexNotExists(index int) bool {
	return l.Size < index
}
