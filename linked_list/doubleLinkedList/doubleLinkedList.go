package doubleLinkedList

import (
	"errors"
	"fmt"
)

type node struct {
	Val  int
	Next *node
	Prev *node
}

type DoubleLinkedList struct {
	Head *node
	Tail *node
	size int
}

// O(1) time complexity
func (l *DoubleLinkedList) InsertFirst(val int) {
	newNode := &node{Val: val}

	if l.Head == nil {
		l.initList(newNode)
	} else {
		l.Head.Prev = newNode

		newNode.Next = l.Head
		l.Head = newNode
		l.size += 1
	}
}

// O(1) time complexity
func (l *DoubleLinkedList) InsertLast(val int) {
	newNode := &node{Val: val}

	if l.Head == nil {
		l.initList(newNode)
	} else {
		newNode.Prev = l.Tail

		l.Tail.Next = newNode
		l.Tail = newNode
		l.size += 1
	}
}

// O(n) time complexity
func (l *DoubleLinkedList) InsertAt(val int, idx int) error {
	if idx == 0 { // First position
		l.InsertFirst(val)
		return nil
	} else if idx == l.size { // Last position
		l.InsertLast(val)
		return nil
	}

	if l.size >= idx {
		newNode := &node{Val: val}

		node := l.Head

		// Check if idx is in the first half of the list
		if idx <= (l.size-1)/2 {
			for i := 0; i < idx-1; i++ {
				node = node.Next
			}
		} else {
			node := l.Tail
			for i := l.size - 1; i > idx-1; i-- {
				node = node.Prev
			}
		}

		newNode.Prev = node
		newNode.Next = node.Next

		node.Next.Prev = newNode
		node.Next = newNode
		l.size += 1
	} else {
		err := fmt.Errorf("can not insert %d at position %d (must be inside the size range of the list %d)", val, idx, l.size)
		return err
	}

	return nil
}

// O(1) time complexity
func (l *DoubleLinkedList) RemoveFirst() error {
	if l.Head == nil {
		return errors.New("can not remove header, if nil")
	}

	if l.size == 1 {
		l.emptyList()
	} else {
		l.Head.Next.Prev = nil
		prevHead := l.Head
		l.Head = prevHead.Next
		prevHead = nil
		l.size -= 1
	}

	return nil
}

// O(1) time complexity
func (l *DoubleLinkedList) RemoveLast() error {
	if l.Tail == nil {
		return errors.New("can not remove tail, if nil")
	}

	if l.size == 1 {
		l.emptyList()
	} else {
		l.Tail.Prev.Next = nil
		prevTail := l.Tail
		l.Tail = prevTail.Prev
		prevTail = nil
		l.size -= 1
	}

	return nil
}

func (l *DoubleLinkedList) RemoveAt(idx int) error {
	if idx == 0 { // First position
		err := l.RemoveFirst()
		return err
	} else if idx == l.size { // Last position
		err := l.RemoveLast()
		return err
	}

	if l.size >= idx {
		node := l.Head

		// Check if idx is in the first half of the list
		if idx <= (l.size-1)/2 {
			for i := 0; i < idx-1; i++ {
				node = node.Next
			}
		} else {
			node := l.Tail
			for i := l.size - 1; i > idx-1; i-- {
				node = node.Prev
			}
		}

		nodeToDel := node.Next
		node.Next = nodeToDel.Next
		node.Next.Prev = node
		nodeToDel = nil
		l.size -= 1

	} else {
		err := fmt.Errorf("can not delete at position %d (must be inside the size range of the list %d)", idx, l.size)
		return err
	}

	return nil

}

func (l *DoubleLinkedList) Print() {
	node := l.Head

	fmt.Print("Head - ")

	for node.Next != nil {
		fmt.Print(node.Val, " -> ")
		node = node.Next
	}
	fmt.Println(node.Val, "- Tail")
}

func (l *DoubleLinkedList) initList(node *node) {
	l.Head = node
	l.Tail = node
	l.size = 1
}

func (l *DoubleLinkedList) emptyList() {
	l.Head = nil
	l.Tail = nil
	l.size = 0
}
