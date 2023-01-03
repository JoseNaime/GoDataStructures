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
