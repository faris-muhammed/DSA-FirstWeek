package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{Value: value, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *LinkedList) Length() int {
	current := list.Head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (list *LinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ->", current.Value)
		current = current.Next
	}
	fmt.Printf("nil")
}

func main() {
	list := LinkedList{}

	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)
	list.InsertAtEnd(50)
	list.InsertAtEnd(60)

	fmt.Println("Length of the list : ", list.Length())
	list.PrintList()
}
