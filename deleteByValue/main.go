package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
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

func (list *LinkedList) DeleteByValue(value int) {
	if list.Head == nil {
		return
	}
	if list.Head.Value == value {
		list.Head = list.Head.Next
		return
	}
	current := list.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (list *LinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d - >", current.Value)
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

	list.PrintList()

	fmt.Println()
	list.DeleteByValue(30)
	list.PrintList()
}
