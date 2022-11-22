# Introduction

## 1.1 Variables 
Variables in computer are containers or space that can store data. 

## 1.2 Data Types
Data type is the specific type of data that can be stored by a variable. Based on the type, numerous operations can be performed. Example: integer, float, array, character, string, etc.

All the data types has some predefined space allocation in the memory. For example, int can store 2 bytes (16 bits, since 1 byte = 8 bits), float has 4 bytes of memory allocation.

Note: Memory allocation may vary depending upon the comipler.

There are two types of data types:
1. System-defined data types (Primitive Data Types)
2. User-defined data types

### System-defined Data Types (Primitive Data Types) 
These are the pre-defined primitive data types that a programming language already have it built-in. 
Example:
  1. int
  2. float
  3. char
  4. double
  5. string
  6. bool

### User Defined Data Types
These are data types which developers can define on their own. These data types are usually composed of primitive data types itself. 
Structures(in C), & Classes (in Dart) are examples of such data types.

Examples:

```dart
class Student{
  String name;
  int rollNumber;
  int phoneNumber;
}
```



```c
struct Student{
  char name[];
  int rollNumber;
  int phoneNumber;
};
```

