package main

import (
	"DataStructures/data_structures/linked_list"
	"fmt"
)

func main() {
	list := linked_list.SinglyLinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	fmt.Println("Poped: ", list.Pop())
	fmt.Println("Poped: ", list.Pop())
	list.Print()

}
