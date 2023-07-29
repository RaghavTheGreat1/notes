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
	fmt.Println(linkedList.Size())
	fmt.Println(linkedList.IsEmpty())

	linkedList.InsertFront(1)
	linkedList.InsertFront(11)
	linkedList.InsertFront(21)
	linkedList.InsertFront(31)
	linkedList.Display()
	fmt.Println(linkedList.Size())

	linkedList.InsertRear(1)
	linkedList.InsertRear(11)
	linkedList.InsertRear(21)
	linkedList.InsertRear(31)
	linkedList.Display()

	fmt.Println(linkedList.Size())
	fmt.Println(linkedList.IsEmpty())

	if err := linkedList.InsertAt(8, 5); err != nil {
		fmt.Println(err)
	}
	linkedList.Display()
	if data, err := linkedList.GetByIndex(8); err == nil {
		fmt.Println(data)
	}

	if data, err := linkedList.DeleteFront(); err == nil {
		fmt.Println("Delete Front: ", data)
	}
	if data, err := linkedList.DeleteFront(); err == nil {
		fmt.Println("Delete Front: ", data)
	}
	linkedList.Display()

	if data, err := linkedList.DeleteRear(); err == nil {
		fmt.Println("Delete Rear: ", data)
	}
	if data, err := linkedList.DeleteRear(); err == nil {
		fmt.Println("Delete Rear: ", data)
	}
	linkedList.Display()

	if data, err := linkedList.DeleteAt(4); err == nil {
		fmt.Println("Delete At index 4: ", data)
	}
	if data, err := linkedList.DeleteAt(1); err == nil {
		fmt.Println("Delete Rear index 1: ", data)
	}
	if data, err := linkedList.DeleteAt(1); err == nil {
		fmt.Println("Delete Rear index 1: ", data)
	}
	linkedList.Display()

	if err := linkedList.DeleteAll(); err != nil {
		fmt.Println("Delete All: ", err)
	}

	linkedList.Display()

	fmt.Println()
	fmt.Println("===== ===== ===== ===== =====")

}
