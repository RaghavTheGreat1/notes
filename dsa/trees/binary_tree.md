# Binary Tree

<p align="center">
<img width="806" alt="image" src="https://user-images.githubusercontent.com/28825619/216805684-c00b5a34-2621-4423-8b0d-2d4346b92fe4.png">
</p>
  
- Tree having nodes with atmost two children (except leaf nodes)
- Empty Tree is also a valid Binary Tree

<p align="center">
<img width="920" alt="image" src="https://user-images.githubusercontent.com/28825619/216805770-f322fb77-ab28-48f6-ae36-2a28d2bdef4b.png">
</p>



## Types of Binary Tree (BT)

### Strict Binary Tree

<p align="center">
<img width="943" alt="image" src="https://user-images.githubusercontent.com/28825619/216805978-90419baa-45b8-4f7a-a546-b08de85c0096.png">
</p>

- Each node has exactly two children or no children

### Full Binary Tree

<p align="center">
<img width="903" alt="image" src="https://user-images.githubusercontent.com/28825619/216805891-9f64ef4e-ec3e-48c6-9e13-6b151727a18c.png">
</p>

- Each node has exactly two children
- All leaf node at same level


## Structure of Binary Tree

<p align="center">
<img width="953" alt="image" src="https://user-images.githubusercontent.com/28825619/216806370-5c52ae41-7f9e-4859-a55b-b593dc167951.png">
</p>

```c
struct BinaryTreeNode{
  int data;
  struct BinaryTreeNode *left;
  struct BinaryTreeNode *right;
};
```

## Operations on Binary Tree

- Insert
- Delete
- Search
- Traverse

## Binary Tree Traversal 

- Visitng each & every nodes of a tree is called tree traversal
- Each node is processed only once but can be traversed many times
- There aare 6 possibilities, but technically it's 3.

Three traversal techniques:
1. Preorder (NLR) - Current NODE -> Left Sub Tree -> Right Sub Tree
2. Inorder (LNR) - Left Sub Tree -> Current NODE -> Right Sub Tree
3. Postorder (LRN) - Left Sub Tree -> Right Sub Tree -> Current NODE

All these possibilities are based on direction of traversal and the processing of it.




