# Linked Lists

## What is Linked List?

- Data Structure that can be used to store data in nodes
- Each node has two items, *data* & *link* to next node
- Number of nodes is not fixed, hence can be used to dynamically expand & shrink the size of DS
- Entry node is called *head*
- Last node of LL has link = NULL
- Solves problem of memory wastage


# Singly Linked Lists

<img width="1010" alt="image" src="https://user-images.githubusercontent.com/28825619/212460591-c67be464-1b88-473a-84d3-8e7017345694.png">



Link of Last Node of LL contains NULL indicating the end of list, hence this makes it a Singly Linked List


```c
struct List{
    int data;
    // Pointer to next List node
    struct List *link;
};

// Makes it easy to using pointer of LL
typedef struct List *NODE;
```

## Basic Operations on SLL
    - Insertion
        a. Front
        b. End
        c. Index
    - Traversing 
    - Deletion
    


### Insertion



#### 1. Front Insertion


Front insertion is when a new node is attached *before* the current head node.

Two step process: 

    1. Update the link of new node to the current head
    
    2. Update head pointer to the new node

<img width="1255" alt="image" src="https://user-images.githubusercontent.com/28825619/212463672-e2ed2926-1f87-4ef9-9838-891dd85290bd.png">




```c

NODE insertFront(NODE head, int data){

    NODE newNode;

    // creates an unnamed node & pointer
    newNode = (NODE)malloc(sizeof(struct List));

    // assigns the input data to data of newNode 
        newNode -> data = data;

    if(head == NULL){

        // if head is empty, implies list is empty
        // so creating new node & assigning to head

        head = newNode;
        head -> link = NULL;


    } else{

        // establishing connection by assinging newNode link to head's address
        newNode -> link = head;

        // updates the head to the newNOde
        head = newNode;

    }

    return head;
}
```

### Traversing

- Traverse until we reach the *last node* of SLL
- Last Node has link = NULL, so we have to iterate till the ```link != NULL```
- Good thing about SLL is that we don't need to know the size of it to iterate unlike arrays
- Time Complexity = O(n)
- Space Complexity = O(1)


```c
struct List{
    int data;
    // Pointer to next List node
    struct List *link;
};

typedef struct List *NODE;


int length(NODE head){
    NODE currentNode = head -> link;
    int count = 0;
    while (currentNode != NULL){
        count++;
        currentNode = currentNode -> link;
    }
    return count;
}
```
    
