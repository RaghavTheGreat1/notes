package main

import "fmt"

func main() {
	linkedListDemo()
}

func linkedListDemo() {
	fmt.Println("===== Linked List Demo =====")
	fmt.Println()
	linkedList := LinkedList{}
	linkedList.Display()
	fmt.Println(linkedList.size())
	fmt.Println(linkedList.isEmpty())

	linkedList.InsertFront(1)
	linkedList.InsertFront(11)
	linkedList.InsertFront(21)
	linkedList.InsertFront(31)
	linkedList.Display()
	fmt.Println(linkedList.size())

	linkedList.InsertRear(1)
	linkedList.InsertRear(11)
	linkedList.InsertRear(21)
	linkedList.InsertRear(31)
	linkedList.Display()

	fmt.Println(linkedList.size())
	fmt.Println(linkedList.isEmpty())
	fmt.Println()
	fmt.Println("===== ===== ===== ===== =====")

}
