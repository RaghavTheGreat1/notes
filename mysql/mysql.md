# mysql-notes

## What is Database(DB)?
Basically, it's any collection of related information. In old days, we used to write down mobile numbers on a notebook. That notebook is a database, which stores data like recipent name and it's number. 

## What is Database Management Systems (DBMS)?
A softare program that help users perform CRUD(Create, Read, Update, Delete) operations and maintain a database. More feature includes:
- Maintain large amount of data or information in a grouped manner
- Handles security 
- Create backups
- Import or export data
- Concurrency 
- Interact with software applications.

DBMS is not a database, it's just helps devs to perform CRUD ops on database.

## Two Main Types of Database:

1. **Relationdal Databases or SQL databases**
2. **Non Relational Databases or No SQL databases** 


## Relationdal Databases or SQL databases 
Organizes data in tables(& tables has rows and columns) and has unique keys.
System that allows to do CRUD operations on Relational Databases are called RDBMS (R stands for relational).
Some RDBMS are mySql, postgreSQL, Oracle, mariaDB, etc.
  
All above RDMS uses SQL (Structured Query Language).  

### Structured Query Language
- Standardized Language for interaction with RDBMS
- Used to perform CRUD operations as well as administrative tasks.
- Used to define table & structures.
- SQL code is not always portable from one RDBMS to another.



## Database Queries
DB cQueries are requests that will be sent to DBMS to fetch some specific information. 
As DB strucures gets more & more complex, it becomes more & more difficult to get specific pieces of information.


#####Jargons:
1. Column - Vertical
2. Row - Horizontal
3. Primary Key - A unique attribute about a specific entry that can help in identifying the rest entries, either in a row or a column based on design. It also helps in differentiating two or more entries & hence, it should be unique.

| transaction_id | title       | amount | type    | tag      |
|----------------|-------------|--------|---------|----------|
| 1              | netflix     | 2499   | outflow | ott      |
| 2              | electricity | 1999   | outflow | bills    |
| 3              | ITC         | 2500   | inflow  | dividend |

In the above case, transaction_id should be the primary key.

4. Foreign Key - An attribute(not necessarily unique) that can help link two tables. A foreign key is basically the primary key of another table.

![image](https://user-images.githubusercontent.com/28825619/140561218-dcf6191f-b1ef-43e3-91c1-64ce412deb02.png)
  
A foreign key can also be used to relate the same table that has the foreign key.

![image](https://user-images.githubusercontent.com/28825619/140561976-8ae34a53-6621-488b-8470-4b8c1a027484.png)

5. Composite Key - Two columns that combine to make a primary key(a unique key).

![image](https://user-images.githubusercontent.com/28825619/140562495-b82ad3d4-75be-4414-9ea7-1ee9760213eb.png)

![image](https://user-images.githubusercontent.com/28825619/140562942-a4cb3175-b5f9-478b-9f31-d19cf3ee33bb.png)

---

## SQL Basics

SQL is a language (not a programming language) that can be used to communicate with RDBMS.

SQL implementations vary between systems.

SQL is a hybrid language. It's a package of 4 languages:
1. Data Query Language - Used to query database info & get information that's already stored.
2. Data Definition Language - Used for defining database schemas.
3. Data Control Language - Used for defining security rules.
4. Data Manipulation Language - Used for IUD(Inserting, Updating, Deleting) ops.

### Queries
In SQL, it's set of instruction given to RDBMS to retrieve data that you want to retrieve. 

### Core Datatypes in mySQL
1. INT (Integers)
2. DECIMAL(M,N) [M-> Total number of digits, N-> Post decimal digits]
3. VARCHAR(L) [L-> Length of characters it can store]
4. BLOB [Binary Large Object, used to store Large Data]  
5. DATE [YYYY-MM-DD]
6. TIMESTAMP [YYYY-MM-DD HH-MM-SS]

### Creating Database Tables

1. With Key
```sql
CREATE TABLE <table-name> (
  <column-name> <datatype> <key>,
);
```
Example:

```sql
CREATE TABLE students (
  student_id INT PRIMARY KEY,
  name VARCHAR(20),
  major VARCHAR(20),
);
```

2. Without Key
```sql
CREATE TABLE <table-name> (
  <column-name> <datatype>,
  <key>(<column-name>),
);
```
Example:

```sql
CREATE TABLE students (
  student_id INT,
  name VARCHAR(20),
  major VARCHAR(20),
  PRIMARY KEY(student_id),
);
```

Using semi-colon helps mySQL detect the line ending, hence it's neccessary to add one.

`CREATE TABLE` is the keyword that SQL uses to create a table.
It's not necessary to write `CREATE TABLE` in all caps, you can use in small caps too but convention says to do in all caps so as to differentiate SQL commands.

### Viewing Table

```sql
DESCRIBE <table-name>;
```

Example:

```sql
DESCRIBE students;
```

### Deleteting Table

```sql
DROP TABLE <table-name>;
```

Example:

```sql
DROP TABLE students;
```

### Modifying Table

```sql
ALTER TABLE <table-name> ADD <column-name> <datatype>;
```

Example:

```sql
ALTER TABLE students ADD gpa DECIMAL(3,2);
```

### Deleting Columns

```sql
ALTER TABLE <table-name> DROP COLUMN <column-name>;
```

Example:

```sql
ALTER TABLE students DROP COLUMN gpa;
```

---

## Inserting Data

```sql
INSERT INTO <table-name> VALUES(
  <value>,
  <value>,
  .
  .
  .
);
```

Using the above command, you have to mandatorily enter all the values corresponding to respective columns & that too in same order. You can't skip any of the value else it will throw an error.
But if you want to use the above syntax and don't want to enter the values of specific column, then add ```sql NULL``` keyword, provided it can be nullable.

**Note**: Strings are always encapsulated in '' or "".

Example:
```sql
INSERT INTO students VALUES(
 1,
 'Raghav Joshi',
 'CSE',
);
```

## Inserting Specific Data Only

```sql
INSERT INTO <table-name>(<column-name>, <column-name> ,.... ) VALUES(
  <value>,
  <value>,
  .
  .
  .
);
```

The above syntax helps in entering values to only specific column. If you aren't willing to add data right now, then this is very useful as it sets the value to NULL.

Example:
```sql
INSERT INTO students(student_id, name) VALUES(
  1,
  'Raghav Joshi',
);
```

> Note: You can't insert values that uses same PRIMARY KEY(if there exists any).

### Reading Data

```sql
SELECT * FROM <table-name>;
```
* commands retrieves the table which includes 'all' entries.

Example:

```sql
SELECT * FROM students;
```

---

## Constraints

If you want that a column should not have Null values or maybe, should have a unique value, then you can define one using mysql.

```sql
CREATE TABLE transaction(
  transaciont_id INT PRIMARY KEY,
  transaction_name VARCHAR(50) NOT NULL,
  recepient_id INT UNIQUE,
);
```

The above command creates a table which has a column called transaction_name and is non-nullable, which implies, whenever we insert values in table, we have to give the transaction_name. Also, whenever we give recepient_id, that id should be unique in whole table while we insert the values.

### Default Constraint:
If you don't know some values while inserting it, yet you don't want to make it NULL, then you can specify a default value before hand.

```sql
CREATE TABLE students(
  student_id INT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  branch VARCHAR(50) DEFAULT "undecided",
);
```

### AUTO_INCREMENT
If you want to add values in a table thats auto incrementable (like serial number) then it's also possible. This helps in not entering the values that follow a sequence.

```sql
CREATE TABLE students(
  student_id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  branch VARCHAR(50) DEFAULT "undecided",
);
```
