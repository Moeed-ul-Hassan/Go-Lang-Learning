# Pointers in Go

Pointers are variables that store the **memory address** of another variable. They are a core concept for performance and data manipulation.

## 1. The Basics
- `&` (Address-of): Gets the memory address.
- `*` (Value-at / Dereference): Gets the value stored at an address.

```go
var x int = 10
var ptr *int = &x // ptr holds the address of x

fmt.Println(ptr)  // Prints something like 0xc000012088
fmt.Println(*ptr) // Prints 10 (dereferencing)
```

## 2. Passing by Value vs. Reference
By default, Go passes everything by **value** (it makes a copy).

### Pass by Value (Copy)
```go
func changeValue(val int) {
    val = 20 // Only changes the local copy
}
```

### Pass by Reference (Pointer)
```go
func changeValue(val *int) {
    *val = 20 // Changes the original variable
}
```

## 3. The `new` Function
You can create a pointer to a zero-valued type using `new`.
```go
p := new(int) // p is a *int pointing to a 0
```

## 4. Nil Pointers
A pointer that hasn't been assigned an address is `nil`. Dereferencing a `nil` pointer causes a **panic**.
```go
var p *int
// fmt.Println(*p) // PANIC!
if p != nil {
    fmt.Println(*p)
}
```

## Why use Pointers?
1. **Efficiency**: Avoid copying large structs or arrays.
2. **Modification**: Allow functions to update the data passed to them.
3. **Consistency**: Share a single instance of data across different parts of your program.

## Summary
Pointers in Go are simpler than in C/C++ (no pointer arithmetic), but they are just as powerful for managing memory and state.
