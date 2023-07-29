# Linked Lists

## What is Linked List?

A linked list is a type of data structure used in computer programming. It is a sequence of items, where each item is linked to the next one in the sequence.

Think of a linked list like a chain of paperclips. Each paperclip (which represents a piece of data) is connected to the next one. You can add, remove, or move around the paperclips relatively easily, the same way you can manipulate the items in a linked list.

Each item in a linked list (each "paperclip") is made up of two parts: the data, and a reference (or link) to the next item in the sequence. The data can be any type of information you want to store, like numbers, strings, etc. The reference is like a pointer showing where to find the next item.

One special thing about linked lists is that they are dynamic, meaning they can grow or shrink in size as needed in real-time during your program's execution. This can be a big advantage over traditional arrays, which have a fixed size that you have to specify at the start.

- Each node has two items, *data* & *link* to next node
- Number of nodes is not fixed, hence can be used to dynamically expand & shrink the size of DS
- Entry node is called *head*
- Last node of LL has link = NULL
- Solves problem of memory wastage


## Linked List ADT

A Linked List is an example of an Abstract Data Type (ADT) that can be implemented in various ways. It's a linear data structure where each element is a separate object, often called a node. Each node contains a pointer (or reference) that points to the next node in the sequence, allowing for efficient insertion and removal of elements from any position in the sequence.

Here's how we could define the ADT of a linked list:

- Node: A unit of data containing one or more data fields and a link field. The link field typically contains a reference to the next node.

- LinkedList: A sequence of nodes, where each node is connected by a link field. There may also be a reference to the first node (head) and/or last node (tail) for quick access.

- **Operations**:

- **insert(element, position)**: Insert an element at the specified position.

- **delete(position)**: Delete an element at the specified position.

- **search(element)**: Find the position of the first occurrence of an element.

- **update(element, position)**: Update an element at the specified position.

- **isEmpty()**: Check if the list is empty.

- **size()**: Return the number of nodes in the list.


### Advantages of Linked Lists:

1. **Dynamic Size**: Unlike arrays, Linked Lists are not of a fixed size because they can grow and shrink during runtime by allocating and deallocating memory. Therefore, there is no need to determine the size of a linked list in advance.

2. **Efficient Insertions/Deletions**: Inserting a new node or deleting a node from an existing linked list is more efficient compared to an array. In a linked list, you just need to update a few pointers whereas in an array this operation potentially requires shifting elements.

3. **Memory Efficient**: Unlike an array, a linked list only allocates the memory required for values to be stored. In arrays, you need to allocate memory equivalent to the maximum size irrespective of usage.


### Disadvantages of Linked Lists:

1. **Random Access Not Allowed**: Nodes in a Linked List cannot be accessed randomly; it has to accessed sequentially starting from the head node. This makes accessing nodes relatively time consuming.

2. **Memory Overhead**: Each node in the Linked List needs extra memory to store the address field with the data object. Therefore, Linked Lists require more memory to store the same number of elements than arrays.

3. **Reversing is Difficult**: Reversing the links in a linked list is complex and requires a good understanding of data structures.

4. **No Cache Locality**: Unlike arrays, linked list nodes may not be stored in consecutive memory locations. This could result in a slight time delay due to different memory locations and no locality of reference.

Remember, the choice of data structure greatly depends on the problem you are trying to solve. If your application requires frequent addition and deletion of data elements, linked lists could be a good choice. If you need fast access to elements in a sequence, arrays could be more suitable.