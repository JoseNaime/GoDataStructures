package main

import (
	"DataStructures/linked_list"
	"fmt"
)

func main() {
	list := linked_list.SinglyLinkedList{}

	list.InsertFirst(0)
	list.InsertLast(10)
	err := list.InsertAt(5, 1)
	if err != nil {
		fmt.Print(err)
	}
	list.Print()

}
