
/*LabPROGRAM 5:Design, develop and implement a menu driven Program in C for the following operations on Doubly Linked List (DLL) of  professor Data with fields:
 ID, Name, Branch,Area of Specialization

 a. Create a DLL stack of N Professor's Data.(Insert front/delete Front or Insert rear/Delete rear)
 b. Create a DLL Queue of N Professor's Data.(Insert front/delete rear or Insert rear/Delete front)
Display status of DLL and count the number of nodes in it.*/



#include <stdio.h> 
#include <string.h>
#include <stdlib.h>
#define MAX 30

// ID, Name,Branch,Specialization  
 struct List
 {
  int ID;
  char Name[MAX];
  char Branch[MAX];
  char Specialization[MAX];
 
  struct List *llink;
  struct List *rlink;
 };

#define MALLOC(ptr,s,t)\
	ptr=(t)malloc(s);\
	if(ptr==NULL)\
        { \
	printf("Insufficient Memory\n"); \
	exit; \
}
typedef struct List *NODE;

NODE Create(NODE, int); 
NODE InsFront(NODE);
NODE DelFront(NODE); 
NODE DelRear(NODE);
void Display(NODE);
void ReadData();

int id  ;	        /* Professor id */

char name[MAX];	        /*  Professor  name */

char branch[MAX];		/* branch */ 
char spe[MAX];	/* Specialization	*/ 

int main()
{
	int choice, done;
 	NODE head;
	int n;	/* No. of Professors  */ 
        MALLOC(head, 1, NODE);
        head->llink= head;
        head->rlink= head;
	done = 0;
	 while(!done)
	 {
		printf("\n1.Create\t\n2.InsFront\t3.Delete Front\n");
                printf("4.Delete Rear\t5.Display\t6.Exit\n");
		printf("Enter Choice: ");
		scanf("%d", &choice);
		 switch (choice)
		 {
			case 1: printf("Enter No. of Employees: ");
				scanf("%d", &n);
				head = Create(head, n); 
                                break;
			case 2: ReadData();
				head = InsFront(head);
                                break;
			case 3: head = DelFront(head);
                                break;
			case 4: head = DelRear(head);
                                break; 
		        case 5: Display(head);
				break; 

			case 6: done=1;
				break;
			default: printf(" wrong option\n\n");
				break;
                  }
         }
    return 0;
}


void ReadData()
{
	printf("Enter professor ID: "); 
	scanf("%d", &id);
	printf("Enter Professore Name: "); 
	scanf("%s",name);
	printf("Enter Professor branch: ");
	scanf("%s",branch);
	printf("Enter Area of Specialization: "); 
	scanf("%s",spe); 
}

// To create a List of Employees
NODE Create(NODE head, int n)
{
	int i; 
	 if(head->rlink== head)
	{
     		for (i = 1; i <= n; i++)
		{
			printf("Enter Professor Data <%d>:\n", i);
			ReadData();
			head = InsFront(head);
                } 
	
       
        	return head;
         }
	else
		printf("List already has some  Professors  data\n");
                return head;
}

// Function to insert a node at the front end of DLL
NODE InsFront(NODE head)
{
	NODE temp,cur;
	MALLOC(temp, sizeof(struct List),NODE);
	temp->ID =  id;
	strcpy(temp->Name, name); 
	strcpy(temp->Branch,branch);
	strcpy(temp->Specialization,spe); 
            
        cur=head->rlink;
        
        cur->llink=temp;
	temp->rlink = cur;
        temp->llink=head;
        head->rlink=temp;
        return head; 
}

// Function to delete a node at the front end of DLL
NODE DelFront(NODE  head)
{	NODE cur,next;
	  if(head->rlink ==head)
          {
            printf("DLL is empty\n");
            return head;
          }     
		  cur=head->rlink;
                  next = cur->rlink; 
		   head->rlink = next; 	
                   next->llink=head;
                   printf("Node deleted is %d\n", cur->ID);
		   free(cur);
                   return head;
}
		
// Function to delete a node at the rear end of DLL
NODE DelRear(NODE  head)
{
       
	NODE last;
        NODE prev;
         	if( head->llink == head)
		{
			printf("Can't delete,  DLL list is empty\n");
			 return  head;
		}
 
                              last=head->llink;
                               prev= last->llink; 
                               prev->rlink = head;
                               head->llink=prev; 
                               printf("Deleted node is %d\n", last->ID); 
			        free(last);
				return head;
}

void Display(NODE head)
{
        NODE temp;

	int count=0; 
	    if( head->rlink == head)
           { 
		printf("\n Doubly linked List is Empty\n");
                return; 
           }
               temp=head->rlink;
              printf("ID\tPName\tPBranch\tPSpecialization\n");
              printf("--------------------------------------------------\n"); 
		while(temp!=head)
		{
		  printf("%d\t%s\t%s\t%s\n", temp->ID, temp->Name, temp->Branch, 
                          temp->Specialization);
	          temp = temp->rlink;
                  count++;
		}
		printf("count of nodes=%d\n",count);
}



