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

// Remove duplicates in a sorted linked list
func (list *LinkedList) RemoveDuplicates() {
	current := list.Head
	if current == nil {
		return
	}

	for current.Next != nil {
		if current.Value == current.Next.Value {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
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
	list.InsertAtEnd(50)

	list.RemoveDuplicates()
	list.PrintList()
}
