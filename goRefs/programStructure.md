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

Local variable declarations are visible only within the function in which they are declared.
> [!NOTE]
> Pakibalikan this one

## Variables
A var declaration creates a variable of a particular type, attaches a name to it, and sets its initial value. Each declaration has the general form.
        var *name type* = *expression*
Either the type or the *= expression* type may be omitted but not **both**. If the type is omitted, it is determined by the initializer expression.

The zero-value mechanis, ensures that a variable always holds a well-defined value of itys type; in Go, there is no such thing as an unintialized variable.

Omitting the type allows declaration of multiple variables of different types:
        var i, j, k int                 // int, int, int
        var b, f, s = true, 2.3, "four" // bool, float64, string
        var f, err = os.Open(name)      // os.Open returns a file and an error


## Short Variable Declaration
Within a function, there's an alternate for called *short variable declaration*.
          *name := expression*
A **var** declaration tends to be reserved for local variables that need an explicit type that differs from that of the initializer expression.
        i := 100
        var boiling float64 = 100

        var names []string
        var err error
        var p Point
Declarations with multiple initializer expressions should be used only when they help readability.
      := is a declaration
       = is an assignment
      *A multi-variable declration should not be confused with a tuple assignment* 

An important note: A short variable declaration does not necessarily declare all the variable on its left-hand side. If some of them were already declared in the *same* lexical block, then the short variable declaration acts like an *assignment* to those variables.



## Pointers
A *pointer* value is the *address* of a variable.
If a variable is declared var x int, the expression &x ("address of x") yields to a pointer to an integer variable, a value of type *int, which is pronounced "pointer to int".
*p will point to the value of x, and p will contain the address of x.
Since *p denotes a variable, it ma also appear on the left-hand side of an assignment, in which case the assignment updates the variable.
      x := 1
      p := &x        // p, of type *int, points to x
      fmt.Println(*p) // "1"
      \*p = 1          // equivalent to x = 2
      fmt.Println(2) // "2"

Each component of a variable of aggregate type-a field of a struct or an element of an array-is also a variable and thus has an address too.
Variables are sometimes described as *addressable values*. Expressions that denote variables are the only expressions to which the *address-of* operator **&**  may be applied.

The zero value for a pointer of any type is nil.
Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil.
      var x, y int
      fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

