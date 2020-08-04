package main

import "testing"

func TestNodeAppend(t *testing.T) {
	list := List{}
	list.append("node 1")
	list.append("node 2")
	// list.Display()

	if list.head.data != "node 1" {
		t.Errorf("\nNode Append Failed: Head Excpected %s, got %s", "node 1", list.head.data)
	} else if list.tail.data != "node 2" {
		t.Errorf("\nNode Append Failed: Tail Excpected %s, got %s", "node 2", list.tail.data)
	} else {
		t.Logf("\nNode Append Passed: Head %s - Tail %s", list.head.data, list.tail.data)
	}
}
