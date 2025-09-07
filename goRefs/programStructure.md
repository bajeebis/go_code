# Program Structure
The basic structural elements of a Go Program.


## Names
There Go functions, variables, constants, types, statement labels, and packages follow a simple rule: a name begins with a letter or an underscore and may have any number of additional letters, digits, and underscores.
Case matters: **heapSort** and **Heapsort** are different names

if an entity is declared within a function, it is local to that function. If declared outside of a function, it is visible in all files of the package to which it belongs.

If the name begins with an upper-case letter, it is exported, which means that it is visible and accessible outside of its own package and may be referred to by other parts of the program (i.e. **Printf** from **fmt** package)


## Declarations
*Declarations* names a program entity and specifies some or all of its properties. There are four major kinds of Declarations:
  1. var
  2. const
  3. type
  4. func
