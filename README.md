## Chapter 1

Basics & Duplicate Code

function is main

variable usages:
s := "" < short declaration & compact
var s string default initialization with zero value
var s = "" < Rarely used except declaring multiple var
var s string = "" < redundant but ncessary in other cases where they are not of the same type


### Maps
A map is a reference to the data structure created by make. When a map is passed to a function, the function receives a copy of the reference, so any changes the called function makes to the underlying data structure will be visible through the caller’s map reference too.
In our example, the values inserted into the counts map by countLines are seen by main.



### Loops
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

### Make
make for the most part is allocation of memory to the buffer

### Channels
A communication mechanism
- Create a new channel with make(chan val-type)
- Send a value into a channel using the channel <- syntax
- The <-channel syntax receivs a value from the channel

### Go routines
goroutine, concurrent function execution. Channels are important in goroutne's

### Printf verb tables
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
goroutine is a concurrent function execution. A channel is a communication mechanism that allow one goroutine to pass values of a specified type to another goroutine. The function main runs in a goroutine and the go statement creates additional goroutines.
Each fetch sends a value (ch <- expression) on the channel ch, and main receives all of them (<-ch). Having main do all the printing ensures that output from each goroutine is processed as a unit, no danger of interleaving if two goroutines finish at the same time
- Exercise 
  - 1.10 - was able to print the output to a file and facebook and google were always consistent. Something to note is the internet speed and if google and facebook has special cashing for any requests
  - 1.11 - program broke after adding a new site that returns a 404 but accidentally left a '/' in the url and this was also the last argument. I suspect that there's something happening in the channels that a simple '/' broke the response second output. Must see in Section 8.9 for coping outputs
  - 1.11 - IN ERRATUM - it was because the io.Copy was being executed before the lesson's own io.Copy thus making the bytes seem none.

### Web Server
Race condition is avoided via mu.Lock() (sync.Mutex)
ParseForm is nested within an if statement. This allows a local variable declaration that precedes the if condition.

## References
dup 1 good reference: https://stackoverflow.com/a/49715256
