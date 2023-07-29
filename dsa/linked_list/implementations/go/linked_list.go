package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

// Display prints all the nodes of the list
//
// Time Complexity: O(N)
func (list *LinkedList) Display() {
	node := list.Head

	for node != nil {
		fmt.Printf("%d ", node.Data)
		node = node.Next
	}

	fmt.Println("")
}

// GetByIndex gets the element present at the specified index
//
// Throws error if the specified index goes beyond the range
// Index Range: [0, Size)
//
// Time Complexity: O(N)
func (list *LinkedList) GetByIndex(index int) (int, error) {
	size := list.Size()

	if size == 0 {
		return 0, nil
	}

	if index >= size {
		return 0, fmt.Errorf("index out of range, valid index range: 0 - %d", size)
	}

	node := list.Head
	count := index
	for node != nil && count != 0 {
		node = node.Next
		count--
	}

	return node.Data, nil
}

// InsertAt inserts an element into the specified index.
//
// Throws error if the specified index goes beyond the range
// Index Range: [0, Size]
//
// Time Complexity: O(N)
func (list *LinkedList) InsertAt(index int, data int) error {

	size := list.Size()

	if index > size {
		return fmt.Errorf("index out of range, valid index range: 0 - %d", size)
	}

	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newNode
	} else {

		if index == 0 {
			list.InsertFront(data)

		} else if index == size {
			list.InsertRear(data)

		} else {
			node := list.Head
			count := index - 1

			for node.Next != nil && count != 0 {
				node = node.Next
				count--
			}

			newNode.Next = node.Next
			node.Next = newNode

		}
	}
	return nil
}

// InsertFront inserts an element to the first position or 0th index
//
// Time Complexity: O(1)
func (list *LinkedList) InsertFront(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if list.Head != nil {
		// connects new node's link to the current head node
		newNode.Next = list.Head
	}

	list.Head = newNode
}

// InsertRear inserts an element to the last position or nth index, where n is the size of the list
//
// Time Complexity: O(N)
func (list *LinkedList) InsertRear(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newNode
	} else {

		lastNode := list.Head

		// to iterate through the last element of the list
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
		// here, we are the end of the list

		lastNode.Next = newNode
	}
}

// IsEmpty checks whether the list is empty or not
//
// Time Complexity: O(1)
func (list *LinkedList) IsEmpty() bool {

	if list.Head == nil {
		return true
	}

	return false
}

// Size returns the number of elements in the list
//
// Time Complexity: O(N)
func (list *LinkedList) Size() int {
	count := 0

	node := list.Head

	for node != nil {
		count++
		node = node.Next
	}

	return count
}
