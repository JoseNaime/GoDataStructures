package linked_list

import (
	"testing"
)

func TestSinglyLinkedList_InsertFirst(t *testing.T) {
	list := SinglyLinkedList{}
	list.InsertFirst(5)

	if list.Head == nil {
		t.Error("InsertFirst() must set a Head value if the list is empty, and first item is added")
	}
	if list.Tail == nil {
		t.Error("InsertFirst() must set a Tail value if the list is empty, and first item is added")
	}
	if list.size != 1 {
		t.Errorf("InsertFirst() must set increase the list size. Got %d; expects 1", list.size)
	}
	if list.Head.Val == 3 {
		t.Errorf("InsertFirst() must create a node with the value passed. Got %d; expects 5", list.Head.Val)
	}

	list.InsertFirst(10)
	if list.Head.Val != 10 {
		t.Errorf("InsertFirst() must add elements to the list Head. Got %d; expects 10", list.Head.Val)
	}
}

func TestSinglyLinkedList_InsertLast(t *testing.T) {
	list := SinglyLinkedList{}
	list.InsertLast(5)

	if list.Head == nil {
		t.Error("InsertLast() must set a Head value if the list is empty, and first item is added")
	}
	if list.Tail == nil {
		t.Error("InsertLast() must set a Tail value if the list is empty, and first item is added")
	}
	if list.size != 1 {
		t.Errorf("InsertLast() must set increase the list size. Got %d; expects 1", list.size)
	}
	if list.Head.Val == 3 {
		t.Errorf("InsertLast() must create a node with the value passed. Got %d; expects 5", list.Tail.Val)
	}

	list.InsertLast(10)
	if list.Tail.Val != 10 {
		t.Errorf("InsertLast() must add elements to the list Head. Got %d; expects 10", list.Tail.Val)
	}
}

func TestSinglyLinkedList_InsertAt(t *testing.T) {
	list := SinglyLinkedList{}
	err := list.InsertAt(6, 4)

	// Test to insert in no range idx
	if err == nil || list.size != 0 {
		t.Errorf("Error must be returned as idx is out of size range, and no value must be inserted, list size: %d  idx: %d", list.size, 4)
	}

	// Test insert at 0
	err = list.InsertAt(0, 0)

	if list.size != 1 || list.Head.Val != 0 || err != nil {
		t.Error("Error inserting value at idx 0")
	}

	// Test insert at size - 1
	err = list.InsertAt(2, list.size)

	if list.size != 2 || list.Tail.Val != 2 || err != nil {
		t.Error("Error inserting value at Tail idx")
	}

	// Test insert at 1
	err = list.InsertAt(1, 1)

	if list.size != 3 || list.Head.Next.Val != 1 || err != nil {
		t.Errorf("Error inserting value at idx 1, Got: %d Expects: 1", list.Head.Next.Val)
	}
}
