package main

import (
	"DataStructures/linked_list/singlyLinkedList"
	"fmt"
)

func main() {
	list := new(singlyLinkedList.SinglyLinkedList)

	list.InsertFirst(1)
	list.InsertLast(3)
	fmt.Println(" --- Insert First and Insert Last")
	list.Print()

	err := list.InsertAt(2, 1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("\n --- Insert 2 at index 1")
	list.Print()

	list.InsertLast(4)
	fmt.Println("\n --- Insert 4 with Insert Last")
	list.Print()

	_ = list.RemoveAt(1)
	fmt.Println("\n --- Remove node at index 1")
	list.Print()
}
