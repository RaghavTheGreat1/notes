# Singly Linked Lists

<img width="1010" alt="image" src="https://user-images.githubusercontent.com/28825619/212460591-c67be464-1b88-473a-84d3-8e7017345694.png">



Link of Last Node of LL contains NULL indicating the end of list, hence this makes it a Singly Linked List


```go
type Node struct {
    Data int
    Next *Node
}

type LinkedList struct {
    Head *Node
}
```

## Basic Operations on SLL
    - Insertion
        a. Front
        b. End
        c. Position
    - Traversing 
    - Deletion
        a. Front
        b. End
        c. Position    


## Insertion



### 1. Front Insertion


Front insertion is when a new node is attached *before* the current head node.

Two step process: 

    1. Update the link of new node to the current head
    2. Update head pointer to the new node

<img width="1255" alt="image" src="https://user-images.githubusercontent.com/28825619/212463672-e2ed2926-1f87-4ef9-9838-891dd85290bd.png">




Time Complexity: O(1)

Space Complexity: O(1)


```go
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
```



### 2. End Insertion

End insertion is when a new node is attached *after* the last node.

Two step process:

    1. Iterate throught SLL until we reach current last node
    2. Change link of last node from NULL to newNode address


<img width="1208" alt="image" src="https://user-images.githubusercontent.com/28825619/212464530-bab64a40-d34d-4806-b693-376a9ec2e224.png">



Time Complexity: O(n) [Since we have to iterate through SLL to reach end of SLL]

Space Complexity: O(1)


```go

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

```



### 3. Position Insertion

Position insertion is adding node to the specified position.

Three step process:

    1. Iterate until iterator reaches the given position 
    2. Assign the link of preNode to the newNode link
    3. Update the link of preNode to newNode's address

<img width="1202" alt="image" src="https://user-images.githubusercontent.com/28825619/212468258-97bb985d-7b93-45fa-bac1-2d4c623d606f.png">



Time Complexity: O(n) [Worst case scenario]

Space Complexity: O(1)

```go
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

```


## Traversing

- Traverse until we reach the *last node* of SLL
- Last Node has link = NULL, so we have to iterate till the ```link != NULL```
- Good thing about SLL is that we don't need to know the size of it to iterate unlike arrays
- Time Complexity = O(n)
- Space Complexity = O(1)


```go
func (list *LinkedList) Size() int {
    count := 0

    node := list.Head

    for node != nil {
        count++
        node = node.Next
    }

    return count
}

```
    

## Deletion

### 1. Front Deletion

Current header node gets deleted and the next node to it becomes the new header node

Two step process:

    1. Create temporary node to point the same node as of current header node
    2. Move the header node pointer to the next node & dispose temporary node
    
    
<img width="1108" alt="image" src="https://user-images.githubusercontent.com/28825619/212478358-c216a621-95ef-4f92-9b0b-4b75e4ab2949.png">



Time Complexity: O(1)

Space Complexity: O(1)


```go
func (list *LinkedList) DeleteFront() (int, error) {

    if list.IsEmpty() {
        return 0, fmt.Errorf("no element found: deletion failed since list is empty")
    } else {
        dataAboutDelete := list.Head.Data
        list.Head = list.Head.Next
        return dataAboutDelete, nil
    }
}

```


### 2. End Deletion

Tail node gets deleted and the link of previous node to it becomes the NULL

Two step process:

    1. Iterate until we reach last node, also storing the previous node
    2. Assign previous node's link to NULL & dispose currentNode

    
<img width="1111" alt="image" src="https://user-images.githubusercontent.com/28825619/212494694-10096fa7-9396-45f8-bd1e-26aea7a540fa.png">



Time Complexity: O(n)

Space Complexity: O(1)


```go
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

```


### 3. Position Deletion

Node present at given position nee

Two step process:

    1. Create temporary node to point the same node as of current header node
    2. Move the header node pointer to the next node & dispose temporary node

    
<img width="1055" alt="image" src="https://user-images.githubusercontent.com/28825619/212528704-a9a9a046-f487-46ea-a33a-9de180b4c2c2.png">


Time Complexity: O(n)

Space Complexity: O(1)


```go
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

```
