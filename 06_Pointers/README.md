# Pointers

Pointers store the memory address of a value.

## Basics
- `&` operator: Gets the memory address of a variable.
- `*` operator: Dereferences a pointer to get the value at that address.

```go
var x int = 10
var p *int = &x // p points to x
fmt.Println(*p) // Prints 10
```

## Why use Pointers?
To pass variables by reference instead of by value (copying), which is more efficient and allows modifying the original data.
