package linked_list

import (
	"testing"
)

func TestSinglyLinkedList_InsertFirst(t *testing.T) {
	l := SinglyLinkedList{}
	l.InsertFirst(5)

	if l.Head == nil {
		t.Error("InsertFirst() must set a Head value if the l is empty, and first item is added")
		l.Print()
	}
	if l.Tail == nil {
		t.Error("InsertFirst() must set a Tail value if the l is empty, and first item is added")
		l.Print()
	}
	if l.size != 1 {
		t.Errorf("InsertFirst() must set increase the l size. Got %d; expects 1", l.size)
		l.Print()
	}
	if l.Head.Val == 3 {
		t.Errorf("InsertFirst() must create a node with the value passed. Got %d; expects 5", l.Head.Val)
		l.Print()
	}

	l.InsertFirst(10)
	if l.Head.Val != 10 {
		t.Errorf("InsertFirst() must add elements to the l Head. Got %d; expects 10", l.Head.Val)
		l.Print()
	}
}

func TestSinglyLinkedList_InsertLast(t *testing.T) {
	l := SinglyLinkedList{}
	l.InsertLast(5)

	if l.Head == nil {
		t.Error("InsertLast() must set a Head value if the l is empty, and first item is added")
		l.Print()
	}
	if l.Tail == nil {
		t.Error("InsertLast() must set a Tail value if the l is empty, and first item is added")
		l.Print()
	}
	if l.size != 1 {
		t.Errorf("InsertLast() must set increase the l size. Got %d; expects 1", l.size)
		l.Print()
	}
	if l.Head.Val == 3 {
		t.Errorf("InsertLast() must create a node with the value passed. Got %d; expects 5", l.Tail.Val)
		l.Print()
	}

	l.InsertLast(10)
	if l.Tail.Val != 10 {
		t.Errorf("InsertLast() must add elements to the l Head. Got %d; expects 10", l.Tail.Val)
		l.Print()
	}
}

func TestSinglyLinkedList_InsertAt(t *testing.T) {
	l := SinglyLinkedList{}
	err := l.InsertAt(6, 4)

	// Test to insert in no range idx
	if err == nil || l.size != 0 {
		t.Errorf("Error must be returned as idx is out of size range, and no value must be inserted, l size: %d  idx: %d", l.size, 4)
	}

	// Test insert at 0
	err = l.InsertAt(0, 0)

	if l.size != 1 || l.Head.Val != 0 || err != nil {
		t.Error("Error inserting value at idx 0")
	}

	// Test insert at size - 1
	err = l.InsertAt(2, l.size)

	if l.size != 2 || l.Tail.Val != 2 || err != nil {
		t.Error("Error inserting value at Tail idx")
		l.Print()
	}

	// Test insert at 1
	err = l.InsertAt(1, 1)

	if l.size != 3 || l.Head.Next.Val != 1 || err != nil {
		t.Errorf("Error inserting value at idx 1, Got: %d Expects: 1", l.Head.Next.Val)
		l.Print()
	}
}

func TestSinglyLinkedList_RemoveFirst(t *testing.T) {
	l := SinglyLinkedList{}
	err := l.RemoveFirst()

	if err == nil {
		t.Errorf("Error must be returned if list is empty, error: %s", err)
	}

	l.InsertFirst(50)

	err = l.RemoveFirst()

	if err != nil || l.Head != nil || l.Tail != nil {
		t.Errorf("Error deleting when size is 1. Head and Tail must be empty, error: %s", err)
		l.Print()
	}

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)

	err = l.RemoveFirst()

	if err != nil || l.Head.Val != 2 || l.Tail.Val != 3 {
		t.Errorf("Error deleting when size is greater than 1. Head must be changed, Got: %d Expects: 2  error: %s", l.Head.Val, err)
		l.Print()
	}
}

func TestSinglyLinkedList_RemoveLast(t *testing.T) {
	l := SinglyLinkedList{}
	err := l.RemoveLast()

	if err == nil {
		t.Errorf("Error must be returned if list is empty, error: %s", err)
	}

	l.InsertFirst(50)

	err = l.RemoveLast()

	if err != nil || l.Head != nil || l.Tail != nil {
		t.Errorf("Error deleting when size is 1. Head and Tail must be empty, error: %s", err)
		l.Print()
	}

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)

	err = l.RemoveLast()

	if err != nil || l.Tail.Val != 2 || l.Head.Val != 1 {
		t.Errorf("Error deleting when size is greater than 1. Tail must be changed, Got: %d Expects: 2  error: %s", l.Tail.Val, err)
		l.Print()
	}
}

func TestSinglyLinkedList_RemoveAt(t *testing.T) {
	l := SinglyLinkedList{}
	err := l.RemoveAt(4)

	// Test to insert in no range idx
	if err == nil || l.size != 0 {
		t.Errorf("Error must be returned as idx is out of size range, l size: %d  idx: %d", l.size, 4)
	}

	l.InsertFirst(5)
	err = l.RemoveAt(0)

	if err != nil || l.Head != nil || l.Tail != nil {
		t.Errorf("If list size is 1, Remove At 0 must delete all the list content, error: %s", err)
		l.Print()
	}

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	err = l.RemoveAt(1)

	if err != nil || l.size != 2 || l.Head.Next.Val != 3 {
		t.Errorf("Do not delete value at index, Got: %d Expects: %d, error: %s", l.Head.Next.Val, 3, err)
		l.Print()
	}

}
