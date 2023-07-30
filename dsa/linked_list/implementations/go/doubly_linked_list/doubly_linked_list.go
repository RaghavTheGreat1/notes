package doubly_linked_list

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type DoublyLinkedList struct {
	Head *Node
}

// Display prints all the nodes of the list
//
// Time Complexity: O(N)
func (list *DoublyLinkedList) Display() {
	node := list.Head

	for node != nil {
		fmt.Printf("%d ", node.Data)
		node = node.Right
	}

	if list.IsEmpty() {
		fmt.Println("nil")
	} else {
		fmt.Println()
	}

}

func (list *DoublyLinkedList) DeleteAll() error {
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
func (list *DoublyLinkedList) DeleteAt(index int) (int, error) {
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

		for node.Right != nil && count != 0 {
			prevNode = node
			node = node.Right
			count--
		}

		dataAboutDelete := node.Data

		// Check if only one element is left
		if prevNode.Right == nil {
			list.Head = nil
			return dataAboutDelete, nil
		} else {
			node.Right.Left = prevNode
			prevNode.Right = node.Right
			return dataAboutDelete, nil
		}
	}
}

// DeleteFront deletes the first element or 0th index element in the list
//
// # Throws error if list is empty
//
// Time Complexity: O(1)
func (list *DoublyLinkedList) DeleteFront() (int, error) {

	if list.IsEmpty() {
		return 0, fmt.Errorf("no element found: deletion failed since list is empty")
	} else {
		dataAboutDelete := list.Head.Data
		list.Head = list.Head.Right
		return dataAboutDelete, nil
	}
}

// DeleteRear deletes the last or nth -1 index element in the list
//
// # Throws error if list is empty
//
// Time Complexity: O(N)
func (list *DoublyLinkedList) DeleteRear() (int, error) {
	if list.IsEmpty() {
		return 0, fmt.Errorf("no element found: deletion failed since list is empty")
	} else {
		node := list.Head

		prevNode := list.Head
		for node.Right != nil {
			prevNode = node
			node = node.Right
		}

		dataAboutDelete := node.Data

		// Check if only one element is left
		if prevNode.Right == nil {
			list.Head = nil
			return dataAboutDelete, nil
		} else {
			prevNode.Right = nil
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
func (list *DoublyLinkedList) GetByIndex(index int) (int, error) {
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
		node = node.Right
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
func (list *DoublyLinkedList) InsertAt(index int, data int) error {

	size := list.Size()

	if index > size {
		return fmt.Errorf("index out of range, valid index range: [0 - %d]", size)
	}

	newNode := &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
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

			for node.Right != nil && count != 0 {
				node = node.Right
				count--
			}

			newNode.Left = node
			newNode.Right = node.Right
			node.Right = newNode

		}
	}
	return nil
}

// InsertFront inserts an element to the first position or 0th index
//
// Time Complexity: O(1)
func (list *DoublyLinkedList) InsertFront(data int) {
	newNode := &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}

	if list.Head != nil {
		// connects new node's link to the current head node
		newNode.Right = list.Head
	}

	list.Head = newNode
}

// InsertRear inserts an element to the last position or nth index, where n is the size of the list
//
// Time Complexity: O(N)
func (list *DoublyLinkedList) InsertRear(data int) {
	newNode := &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}

	if list.Head == nil {
		list.Head = newNode
	} else {

		lastNode := list.Head

		// to iterate through the last element of the list
		for lastNode.Right != nil {
			lastNode = lastNode.Right
		}

		// here, we are the end of the list
		newNode.Left = lastNode
		lastNode.Right = newNode
	}
}

// IsEmpty checks whether the list is empty or not
//
// Time Complexity: O(1)
func (list *DoublyLinkedList) IsEmpty() bool {

	if list.Head == nil {
		return true
	}

	return false
}

// Size returns the number of elements in the list
//
// Time Complexity: O(N)
func (list *DoublyLinkedList) Size() int {
	count := 0

	node := list.Head

	for node != nil {
		count++
		node = node.Right
	}

	return count
}
