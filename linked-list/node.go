// A basic doubly linked list in go for practice
package main

import (
	"fmt"
)

// Node an obj. to store data
type Node struct {
	data string
	next *Node
	prev *Node
}

// List will use nodes to and chain them together
type List struct {
	head *Node
	tail *Node
}

// Append a node to the list (not exported to show that export is at package level)
func (l *List) append(val string) {
	// Refrence to head node
	// if there is no  head node create it
	//TODO create a current node to assign to head / tail
	if l.head == nil {
		node := &Node{
			data: val,
			next: nil,
			prev: nil,
		}
		l.head = node
		l.tail = node
		return
	}
	currentNode := l.head
	// While current node has next progress to it
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &Node{val, nil, currentNode}
	l.tail = currentNode.next
}

// Display a list
func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%s", list.data)
		if list.next != nil {
			fmt.Print(" -> ")
		}
		list = list.next
	}
	fmt.Print("\n")
}
