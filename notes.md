Chapter 1

Basics & Duplicate Code

function is main

variable usages:
s := "" < short declaration & compact
var s string default initialization with zero value
var s = "" < Rarely used except declaring multiple var
var s string = "" < redundant but ncessary in other cases where they are not of the same type


## Maps
A map is a reference to the data structure created by make. When a map is passed to a function, the function receives a copy of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller’s map reference too.
In our example, the values inserted into the counts map by countLines are seen by main.



## Loops
for initialization; condition; post {
    // zero or more statements
}

// traditional while loop
for condition {
  // ...
}

// traditional infinite loop
for {
  // ...
}

## Make
make for the most part is allocation of memory to the buffer

## Printf verb tables
%d          decimal integer
%x, %o, %b  integer in hexadecimal, octal, binary
%f, %g, %e  floating-point number
%t          boolean: true or false
%c          rune (Unicode code point)
%s          string
%q          quoted string "abc" or rune 'c'
%v          any value in a natural format
%T          type of any value
%%          literal percent sign (no operand)


A map is a reference to the data structure created by make. When a map is passed to a function, the function receives a copy of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller's map reference too.
In our example, the values inserted into the counts map by countLines are seen by main.

## Lessons

### Lissajous
- const declaration is unchangeable
- The expressions []color.Color{...} and gif.GIF{...} are composite literals, a compact notation for instantiating any of Go's composite types from a sequence of element values. Here, the first one is a slice and the second one is a struct.

### Fetch
- new two package introduction, net/http and io/ioutl.
- io.ReadAll reads the entire response, and it is stored in b.
- Body stream should be closed to avoid leaking resources
- Lastly, Printf writes th response to the standard output. First it was bytes but I converted it to strings via strings().

### Fetching URLs Concurrently
- goroutine is a concurrent function execution. A channel is a communication mechanism that allow one goroutine to pass values of a specified type to another goroutine. The function main runs in a goroutine and the go statement creates additional goroutines.




## References
dup 1 good reference: https://stackoverflow.com/a/49715256
