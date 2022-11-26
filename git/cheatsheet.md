# Git Github Cheatsheet



## Introduction

Git is a version control system *(VCS)*. Version control systems generally allow users to revisit earlier versions of the files, compare changes between versions, undo changes, and a whole lot more.

---

## Let's Get Started

### Basic Setup Commands

**1. Setting up User Name**

   > $ git config --global user.name "\<full name>"

**2. Setting up Email Address**

  > $ git config --global user.email \<email-id>

---

### Basic Git Commands

**1. Initializing Repository**
   
   > $ git init
    
The above command tells Git to create a Git Repository in the specified folder and start keeping an eye to the changes happening in the folder.

***A line of caution**, never ever initialize a folder whose root folder is already initialized. If the mistake is already done, then simply delete the **.git** folder which is present in the folder that you have initialized by mistake*

**2. Status of Folder**

   > $ git status

git status gives information on the current status of a git repository and its contents. It can also be used to check whether the folder is initialized or not. 

**3. Staging Changes**

- Adding by file names
    > $ git add \<file-name1> \<file-name2>

    *Note: Filename should include extension.*

- Adding all files in one go
    > $ git add .

Before committing changes, we need to add the files and folders to Staging Area. It's just like adding items into the cart before checkout.

**4. Staging Changes**

- **Simple Staging**
  
   > $ git commit -m "\<message>"

- **Staging & Committing Changes Simultaneously**

   > $ git commit -a -m "\<message>"

    *Note: It adds all the files in one go.*

Staging is just like saving a checkpoint. If anything goes worng in future, one can come back to this checkpoint and can start afresh from here without completely losing the work.

**5. Checking Commit Logs**

> $ git log

git log helps you retrieve the information about who and when the commits were made, and the commit messages.

---

## Git Branch Commands

**1. Fetch List of Branches**

   > $ git branch


**2. Create New Branch**
- Simple Branching

   > $ git branch \<branch-name>

    *Note: Branches are always created based on the HEAD.*

- Branching & Switching Simultaneously

   > $ git switch -c \<branch-name>


    *Note: Always be in the branch from where you want to create a new branch.*

**3. Switch Branches**

   > $ git switch \<branch-name>

*Note 1: Always be in the branch from where you want to create a new branch.*  

*Note 2: This a new method of switching branches. Personally, this new syntax is more handy than the other one.*

*Note 3: Always try to commit changes before switching branches else it might erase all of the work.*

   > $ git checkout \<branch-name>

*Note: The above command is an older method to switch branches. I don't recommend to use this method because it has tons of other use cases.*

**4. Renaming Branches**

   > $ git branch -m \<new-branch-name>

*Note: To be able to rename using above command, you should be in the branch.*

   > $ git branch -m \<oldname> \<newname> 


**5. Deleting Branches**

   > $ git branch -d \<branch-name>

*Note: You should not be in the branch that you want to delete. You have to switch to any different branch so as to delete your desired branch.*

**6. Merging Branches**

   > $ git merge \<branch-name>

*Note 1: Branch name should be the branch that you want to merge*

*Note 2: You should be in the branch that you want your wishing branch to get merged into. For example, you want a branch called 'main' to get merged into 'production' branch, then you should be 'production' branch.*


---
