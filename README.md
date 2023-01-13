# Pure Error

Provide a very tiny Comparable, wrapped-support error.
It only contains code, message, and wrapped error.


## Comparable
You can easily use `errors.Is(err, target)` to check if this error is in wrapped error chain.
For pure error, it returns true just when the error code is the same.

```go
# Pure error
perr := pureerror.New("example", nil).Why("human fault")

# Normal error
err := errors.New("normal error")

# Pure error wrap normal error
e1 := pureerror.New("example", err)

# Normal error wrap pure error
e2 := fmt.Errorf("test: %w", perr)

# Comparable
errors.Is(e1, perr) # true: match by the same pure error code
errors.Is(e1, err)  # true: match by the error in wrapped error chain
errors.Is(e2, perr) # true: match by the pure error with the same code in wrapped error chain
```
