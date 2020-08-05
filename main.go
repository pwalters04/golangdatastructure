package main

import (
	linked_list "data_structure/linked-list"
	"fmt"
)

func main() {
	myList := linked_list.CreateList()
	myList.InsertAtHead("sugar")
	myList.Insert("water")
	myList.Insert("tomato")
	fmt.Println("Water index: ",myList.IndexOf("water"))
	myList.InsertAtPosition("beer",4)
	fmt.Println("What is at index 3: ", myList.PeekAtPosition( 3))
	fmt.Println("List Length: ",linked_list.ListLength(myList))
	linked_list.PrintList(myList)
}
