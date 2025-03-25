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

func (list *LinkedList) FindNthFromNode(n int) *Node {
	first := list.Head
	second := list.Head

	for i := 0; i < n; i++ {
		if first == nil {
			return nil
		}
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
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
	n:=6
	result:=list.FindNthFromNode(n)
	if result!=nil{
		fmt.Printf("%dnth from node is %d\n",n,result.Value)
	}else{
		fmt.Printf("%dnth not exist in node\n",n)
	}
}
