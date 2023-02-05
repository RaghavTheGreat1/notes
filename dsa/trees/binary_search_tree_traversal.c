// Binary Search Tree Traversal

// Credit goes to @divyej (Divye Jain)

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Tree
{
    int data;
    struct Tree *left;
    struct Tree *right;
};

typedef struct Tree *NODE;

#define MALLOC(ptr, s, t)                \
    ptr = (t)malloc(s);                  \
    if (ptr == NULL)                     \
    {                                    \
        printf("Insufficient Memory\n"); \
        exit;                            \
    }

NODE create(NODE);
NODE createBST(NODE, int);
void preOrder(NODE);
void inOrder(NODE);
void postOrder(NODE);

int main()
{
    int choice, done;
    NODE root;
    root = NULL;
    done = 0;

    while (!done)
    {
        printf("1. Create\t2. Preorder\n3. Inorder\t4. Postorder\n5. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);
        switch (choice)
        {
        case 1:
            root = create(root);
            break;
        case 2:
            preOrder(root);
            printf("\n");
            break;
        case 3:
            inOrder(root);
            printf("\n");
            break;
        case 4:
            postOrder(root);
            printf("\n");
            break;
        case 5:
            done = 1;
            break;
        default:
            printf("Invalid Choice\n");
        }
    }
    return 0;
}
// Function to create a Binary Search Tree
NODE create(NODE root)
{
    int i, item, numberOfElements;
    NODE temp;

    printf("Enter number of elements: \n");
    scanf("%d", &numberOfElements);

    if (root == NULL)
    {
        for (i = 0; i < numberOfElements; i++)
        {
            printf("Enter the element to insert: ");
            scanf("%d", &item);
            root = createBST(root, item);
        }
        return root;
    }
    else
        printf("Binary Tree already has some data\n");
    return root;
}

// Function to create a Binary Search Tree
NODE createBST(NODE root, int item)
{
    NODE currentNode, previousNode, temp;
    temp = (NODE)malloc(sizeof(NODE));
    temp->data = item;
    temp->left = temp->right = NULL;

    if (root == NULL)
        return temp;

    previousNode = NULL;
    currentNode = root;
    while (currentNode != NULL)
    {
        previousNode = currentNode;
        if (item == currentNode->data)
        {
            printf("Duplicate Entry\n");
            free(temp);
            return root;
        }
        if (item < currentNode->data)
            currentNode = currentNode->left;
        else
            currentNode = currentNode->right;
    }
    if (item < previousNode->data)
        previousNode->left = temp;
    else
        previousNode->right = temp;
    return root;
}

// Function to traverse the BST in Preorder
void preOrder(NODE root)
{
    if (root != NULL)
    {
        printf("%d\t", root->data);
        preOrder(root->left);
        preOrder(root->right);
    }
}

// Function to traverse the BST in Inorder
void inOrder(NODE root)
{
    if (root != NULL)
    {
        inOrder(root->left);
        printf("%d\t", root->data);
        inOrder(root->right);
    }
}

// Function to traverse the BST in Postorder
void postOrder(NODE root)
{
    if (root != NULL)
    {
        postOrder(root->left);
        postOrder(root->right);
        printf("%d\t", root->data);
    }
}
