# Trees

Trees are non linear Data Structures where each node is connected to various nodes & represents heirarchial nature of a structure. 
In a tree, there's a root node, to it multiple instances can be attached down the tree.

<p align="center">
<img width="511" alt="image" src="https://user-images.githubusercontent.com/28825619/216803306-ba814f97-31b9-4d50-806e-c39031e88db6.png">
</p>

## Tree Terms


<img width="1230" alt="image" src="https://user-images.githubusercontent.com/28825619/216803810-3c44e239-5e8a-4527-9a24-45203f2779ad.png">


**Root Node**: 
- The starting node of a tree, which has no parent node
- There can only be one root node in a tree


**Edge**:
- The link between parent & child nodes

**Leaf Node**:
- A node with no children 


**Level**:
- Set of alll nodes a given dept

**Ancestor & Descendant Node**:
- A node p is ancestor to node q if there exists a path from root to q & p appears on the path
- Node q is descendant of p
- Example: A B E I J are ancestors of K or K is descendant of it

**Depth**:
- Length of path from root to node
- Example: Depth of K is 5 ( A - B - E - I - J - K )
- _Pro Tip_: Just count the number of edges while tracing the path from root to specific node

**Height of Node**:
- Length of path from node to deepest node or leaf node
- Example: Height of C is 1 ( C - F ), B is 4 ( B - E - I - J - K )
- _Pro Tip_: Just count the number of edges while tracing the path from specific node to deepest node or leaf node

