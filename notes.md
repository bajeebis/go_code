# Extra interesting stuff, really

## Fundamentals

### Slices
Dynamically sized

### goroutines
- Good analogy https://old.reddit.com/r/golang/comments/14zeqjq/can_someone_explain_me_the_goroutines/jrxh655/
Concurrent function execution

### channel
Communication mechanism that allow one goroutine to pass values to another goroutine

## io package

### Writer
writer is the interface that wraps the basic Write method
- Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes from p (0 <=n <= len(p)) and any error encountered. Returns a non-nil error if it returns n<len(p).
- Write must no modify the slice data, even temporarily

### io.Discard
Discard is a Writer on which all write calls succeed without doing anything.
i.e. a blackhole (Patapon)

