package main

import (
	linked_list "data_structure/linked-list"
	stack_queues "data_structure/stack-queues"
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

	myStack := stack_queues.CreateStack()
	myStack.Add("Charlie")
	myStack.Add("Bravo")
	myStack.Add("Echo")

	myStack.PrintStack()
	fmt.Println(myStack.Search("Echo"))
	sorted := stack_queues.ReverseStack(myStack)
	sorted.Peek()

}
