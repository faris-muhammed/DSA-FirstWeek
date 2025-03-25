package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{Value: value, Next: list.Head}
	list.Head = newNode
}

func (list *LinkedList) PrintList() {
	current := list.Head

	for current != nil {
		fmt.Printf("%d ->", current.Value)
		current = current.Next
	}
}

func main() {
	list := LinkedList{}

	list.InsertAtBeginning(10)
	list.InsertAtBeginning(20)
	list.InsertAtBeginning(30)
	list.InsertAtBeginning(40)
	list.InsertAtBeginning(50)

	list.PrintList()
}
