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

func (list *LinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ->", current.Value)
		current = current.Next
	}
	fmt.Printf("nil")
}

func (list *LinkedList) Reverse() {
	var prev *Node
	current := list.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	list.Head = prev
}

func main() {
	list := LinkedList{}

	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtEnd(40)
	list.InsertAtEnd(50)

	list.PrintList()
	fmt.Println()
	list.Reverse()
	list.PrintList()
}