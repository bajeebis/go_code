## Pointers
A *pointer* value is the *address* of a variable. Ex.: *p, *x
If a variable is declared var x int, the expression &x ("address of x") yields to a pointer to an integer variable, a value of type *int, which is pronounced "pointer to int".
*p will point to the value of x, and p will contain the address of x.
Since *p denotes a variable, it ma also appear on the left-hand side of an assignment, in which case the assignment updates the variable.
```go
      x := 1
      p := &x        // p, of type *int, points to x
      fmt.Println(*p) // "1"
      \*p = 1          // equivalent to x = 2
      fmt.Println(2) // "2"
```

Each component of a variable of aggregate type-a field of a struct or an element of an array-is also a variable and thus has an address too.
Variables are sometimes described as *addressable values*. Expressions that denote variables are the only expressions to which the *address-of* operator **&**  may be applied.

The zero value for a pointer of any type is nil.
Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil.
```go
      var x, y int
      fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

```
