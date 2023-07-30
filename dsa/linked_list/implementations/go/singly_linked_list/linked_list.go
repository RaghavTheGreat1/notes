package singly_linked_list

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

	if list.IsEmpty() {
		fmt.Println("nil")
	} else {
		fmt.Println()
	}

}

func (list *LinkedList) DeleteAll() error {
	if list.IsEmpty() {
		return fmt.Errorf("no element found: deletion failed since list is empty")
	}

	list.Head = nil

	return nil
}

// DeleteAt deletes element at specified index
//
// Throws error if list is empty or index goes beyond range
//
//	Index Range: [0, Size)
//
// Time Complexity: O(N)
func (list *LinkedList) DeleteAt(index int) (int, error) {
	size := list.Size()
	if index >= size {
		return 0, fmt.Errorf("index out of range, valid index range: [0, %d)", size)
	}

	if list.IsEmpty() {
		return 0, fmt.Errorf("no element found: deletion failed since list is empty")
	} else {
		node := list.Head
		prevNode := list.Head
		count := index

		for node.Next != nil && count != 0 {
			prevNode = node
			node = node.Next
			count--
		}

		dataAboutDelete := node.Data

		// Check if only one element is left
		if prevNode.Next == nil {
			list.Head = nil
			return dataAboutDelete, nil
		} else {
			prevNode.Next = node.Next
			return dataAboutDelete, nil
		}
	}
}

// DeleteFront deletes the first element or 0th index element in the list
//
// # Throws error if list is empty
//
// Time Complexity: O(1)
func (list *LinkedList) DeleteFront() (int, error) {

	if list.IsEmpty() {
		return 0, fmt.Errorf("no element found: deletion failed since list is empty")
	} else {
		dataAboutDelete := list.Head.Data
		list.Head = list.Head.Next
		return dataAboutDelete, nil
	}
}

// DeleteRear deletes the last or nth -1 index element in the list
//
// # Throws error if list is empty
//
// Time Complexity: O(N)
func (list *LinkedList) DeleteRear() (int, error) {
	if list.IsEmpty() {
		return 0, fmt.Errorf("no element found: deletion failed since list is empty")
	} else {
		node := list.Head

		prevNode := list.Head
		for node.Next != nil {
			prevNode = node
			node = node.Next
		}

		dataAboutDelete := node.Data

		// Check if only one element is left
		if prevNode.Next == nil {
			list.Head = nil
			return dataAboutDelete, nil
		} else {
			prevNode.Next = nil
			return dataAboutDelete, nil
		}
	}
}

// GetByIndex gets the element present at the specified index
//
// Throws error if the specified index goes beyond the range
//
//	Index Range: [0, Size)
//
// Time Complexity: O(N)
func (list *LinkedList) GetByIndex(index int) (int, error) {
	size := list.Size()

	if list.IsEmpty() {
		return 0, nil
	}

	if index >= size {
		return 0, fmt.Errorf("index out of range, valid index range: [0, %d)", size)
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
// # Throws error if the specified index goes beyond the range
//
//	Index Range: [0, Size]
//
// Time Complexity: O(N)
func (list *LinkedList) InsertAt(index int, data int) error {

	size := list.Size()

	if index > size {
		return fmt.Errorf("index out of range, valid index range: [0 - %d]", size)
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
