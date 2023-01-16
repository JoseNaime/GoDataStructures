package linked_list

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func (l *SinglyLinkedList) initList(node *Node) {
	l.Head = node
	l.Tail = node
	l.size = 1
}

func (l *SinglyLinkedList) InsertFirst(val int) {
	newNode := &Node{Val: val}

	if l.Head == nil {
		l.initList(newNode)
	} else {
		newNode.Next = l.Head
		l.Head = newNode
		l.size += 1
	}
}

func (l *SinglyLinkedList) InsertLast(val int) {
	newNode := &Node{Val: val}

	if l.Head == nil {
		l.initList(newNode)
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
		l.size += 1
	}
}

func (l *SinglyLinkedList) InsertAt(val int, idx int) error {
	if idx == 0 { // First position
		l.InsertFirst(val)
		return nil
	} else if idx == l.size { // Last position
		l.InsertLast(val)
		return nil
	}

	if l.size >= idx {
		newNode := &Node{Val: val}

		node := l.Head
		for i := 0; i < idx-1; i++ {
			node = node.Next
		}

		newNode.Next = node.Next
		node.Next = newNode
		l.size += 1
	} else {
		err := fmt.Errorf("can not insert %d at position %d (must be inside the size range of the list %d)", val, idx, l.size)
		return err
	}

	return nil
}

func (l *SinglyLinkedList) RemoveFirst() {

}

func (l *SinglyLinkedList) RemoveLast() {

}

func (l *SinglyLinkedList) RemoveAt(idx int) {

}

func (l *SinglyLinkedList) Print() {
	node := l.Head

	fmt.Print("Head - ")

	for node.Next != nil {
		fmt.Print(node.Val, " -> ")
		node = node.Next
	}
	fmt.Println(node.Val, "- Tail")
}
