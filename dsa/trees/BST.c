//In bst all the left tree lements should be less than the root and right tree elements should be greater than the root


#include<stdio.h> 
#include<stdlib.h>
#include<string.h>
#define MAX 30 

struct tree
{
	int info;
	struct tree *left; 
        struct tree *right;
};
 typedef struct tree *NODE;

#define MALLOC(ptr,s,t)\
   ptr=(t)malloc(s);\
    if(ptr==NULL)\
    { \
      printf("insufficient memeory\n"); \
      exit;\
    }

NODE create(NODE, int); 
NODE create_BST(NODE, int);
void Preorder(NODE);
void Postorder(NODE); 
void Inorder(NODE);
int n;

int main()
{
	int choice, done, flag, key;
        NODE root;
	root=NULL;
	done=0; 
   while(!done)
   {
      printf("1.create\t2.preorder\n");
      printf("3.inorder\t4.postorder\t5.exit\n"); 
      printf("Enter your choice\n");
      scanf("%d",&choice); 
     switch (choice)
    {
      case 1: printf("Enter number of  elements\n"); 
              scanf("%d",&n);
              root=create(root,n);
              break;
     case 2: Preorder(root);
             printf("\n");
             break;
     case 3: Inorder(root);
             printf("\n"); 
             break;
     case 4: Postorder(root);
             printf("\n");
             break;
    case 5: done=1;
            break;
   default: printf("invalid choice\n");

  }
 }
return 0;
}
// Function to create a Binary Search Tree
NODE create(NODE root, int n)
{
	int i,item;
	NODE temp;
	if(root==NULL)
	{
	  for(i=1;i<=n;i++)
	  {	
		printf("Enter the element to insert\n");
                scanf("%d", &item); 
                root=create_BST(root,item);
          }
           return root;
        }
       else
	   printf("Binary tree already has some data\n");
           return root;
}

// Function to create a Binary Search Tree
NODE create_BST(NODE root, int item)
{
    NODE cur,prev,temp;
	  temp=(NODE)malloc(sizeof(NODE); 
          temp->info=item;
	  temp->left=temp->right=NULL;
	 
      if(root==NULL) return temp;  // If BST empty

        prev=NULL;
        cur=root;
	while (cur!=NULL)
        {
               prev=cur;
           if(item==cur->info)
             {
                printf("Duplicate entry,Remove\n"); 
                free(temp);
                return root;
              }
               if( item < cur->info)
	         cur=cur->left;
		else
		cur=cur->right;
        }
          if(item< prev->info)
              prev->left=temp;
          else
              prev->right=temp;
	   return root;
}
//Function to traverse the BST in Preorder
void Preorder(NODE root)
{
	if(root!=NULL)
	{
	  printf("%d\t", root->info);
          Preorder(root->left);
          Preorder(root->right);
        } 
}

//Function to traverse the BST in Inorder
void Inorder(NODE root)
{
	if(root !=NULL)
	{
	 Inorder(root->left);
         printf("%d\t", root->info);
         Inorder(root->right);
	}
}

//Function to traverse the BST in Postorder
void Postorder(NODE root)
{
	if( root !=NULL)
	{
	  Postorder(root->left);
          Postorder(root->right);
          printf("%d\t", root->info);
        } 
}


