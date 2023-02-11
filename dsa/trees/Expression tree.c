//A TREE REPRESENTING AN EXPRESSION IS CALLED EXPRESSION TREE,LEAF NODES ARE OPERANDS AND NON LEAF NODES ARE OPERATORS
//INTERNAL NODES -> OPERATORS
//EXTERNAL NODES ->OPERANDS
#include<stdio.h> 
#include<stdlib.h>
#include<string.h>
#define MAX 30 

struct bt
{
	int info;
	struct bt *left; 
    struct bt *right;
};
 typedef struct bt *NODE;

 //ALGO FOR BUILDING EXPRESSION TREE FROM POSTFIX EXPRESSION

 NODE buildExprTree(char postfix[],int size){
    struct Stack *s= Stack(size);
    for (int i=0;i<size;i++){
        if(postfix[i] is a operator)
        {
                NODE temp;
                temp=(NODE)malloc(sizeof(NODE));
                temp->data=postfix[i];
                temp->left=temp->right=NULL;
                push(s,temp);
        }
        else{
            NODE T2=pop(s);
            NODE T1=pop(s);
            NODE temp;
            temp=(NODE)malloc(sizeof(NODE));
            temp->data=postfix[i];
            temp->left=T1;
            temp->right=T2;
            push(s,temp);
        }
    }
    return s;

 }