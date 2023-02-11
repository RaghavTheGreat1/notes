/*A tree is called a binary tree if each node has a zero chilf one child or two children
Types of binary tree 
Struct Binary tree :Exactly two or no children
Full Binary tree  :Exactly two children with all nodes on same level
Complete bonary tree: A binary tree is called a complte binary tree if all leaf nodes are at height h or h-1 and also without missing any number in the sequence 
*/
//Structure of binary tree 
#include<stdio.h> 
#include<stdlib.h>
#include<string.h>
#define MAX 30 
//Structure of binary tree 
struct bt
{
	int info;
	struct bt *left; 
    struct bt *right;
};
 typedef struct bt *NODE;

/*level order traversal
Level order traversal 
//VISIT THE ROOT 
//WHILE TRAVERSAL LEVEL l KEEP ALL THE ELEMENTS AT LEVEL l+1 IN QUEUE
//GOTO NEXT LEVEL AND VISIT ALL NODES AT THAT LEVEL
//REPEAT THIS UNTIL ALL NODES ARE COMPLETED 
*/
void levelorder(NODE root)
{
    NODE temp;
    struct Queue *Q= createQueue();
    if(!root)
    return;
    enQueue(Q,root);
    while(!isEmpty(Q)){
        temp=deQueue(Q);
        //process current node
        printf("\n%d",temp->data);
        if(temp->left)
        enQueue(Q,temp->left);
        if(temp->right)
        enQueue(Q,temp->right);

    }
    deQueue(Q);
}
//Finding max element in a binary tree
int findMax(NODE root)
{
    int item,left,right,max=INT_MIN;
    IF(root!=null)
    {
        item=root->info;
        left=findMax(root->left);
        right=findMax(root->right);
        //Find the largest of three values 
        if(left>right)
        max=left;
        else
         max=right;
         if (item>max)
         max=item;
         
    }
    return max;
}