package linked_list

import "testing"

func TestInsertFirst(t *testing.T) {
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
