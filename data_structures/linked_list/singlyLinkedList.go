package linked_list

type Node struct {
	Val  int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func (l *SinglyLinkedList) InsertFirst(val int) {
	newNode := &Node{Val: val}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.size += 1
}

func (l *SinglyLinkedList) InsertLast(val int) {

}

func (l *SinglyLinkedList) InsertAt(val int) {

}

func (l *SinglyLinkedList) RemoveFirst() {

}

func (l *SinglyLinkedList) RemoveLast() {

}

func (l *SinglyLinkedList) RemoveAt(idx int) {

}
