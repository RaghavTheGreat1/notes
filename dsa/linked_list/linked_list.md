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


- Reference of Last Node of LL contains NULL indicating the end of list, hence this makes it a Singly Linked List

```c
struct List{
    int data;
    // Pointer to next List node
    struct List *link;
};
```

## Basic Operations on SLL
    - Traversing 
    - Insertion
    - Deletion

### Traversing
    
