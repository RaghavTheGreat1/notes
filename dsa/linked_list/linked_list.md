# Linked Lists

## What is Linked List?

- Data Structure that can be used to store data in nodes
- Each node has two items, *data* & *reference* to next node
- Number of nodes is not fixed, hence can be used to dynamically expand & shrink the size of DS
- Entry node is called *head*
- Last node of LL has reference = NULL
- Solves problem of memory wastage


# Singly Linked Lists

<img width="1010" alt="image" src="https://user-images.githubusercontent.com/28825619/212460591-c67be464-1b88-473a-84d3-8e7017345694.png">



Reference of Last Node of LL contains NULL indicating the end of list, hence this makes it a Singly Linked List


```c
struct List{
    int data;
    // Pointer to next List node
    struct List *link;
};
```

## Basic Operations on SLL
    - Insertion
        a. Front
        b. End
        c. Index
    - Traversing 
    - Deletion

### Traversing

- Traverse until we reach the *last node* of SLL
- Last Node has reference = NULL, so we have to iterate till the ```link/reference != NULL```
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
    