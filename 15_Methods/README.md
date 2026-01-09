# Methods in Go

A method is simply a function with a special **receiver** argument. It allows you to "attach" functions to structs.

## 1. Syntax
```go
func (receiver ReceiverType) MethodName(parameters) returnType {
    // code
}
```

## 2. Value vs. Pointer Receivers
This is the most important concept in Go methods.

### Value Receiver `(u User)`
- Works on a **copy** of the struct.
- Changes made inside the method **do not** affect the original.
- Used for: Reading data, small structs.

### Pointer Receiver `(u *User)`
- Works on the **actual** struct in memory.
- Changes made inside the method **affect the original**.
- Used for: Modifying data, large structs (performance).

## 3. Example
```go
type User struct {
    Name string
    Age  int
}

// Value receiver: Can't change age
func (u User) Greet() {
    fmt.Printf("Hi, I'm %s\n", u.Name)
}

// Pointer receiver: Can change age
func (u *User) Birthday() {
    u.Age++
}
```

## 4. Why use Methods?
- **Encapsulation**: Group behavior with data.
- **Interfaces**: Methods are required to implement interfaces.
- **Namespace**: You can have the same method name on different types.

## Common Pitfalls
- **Mixing Receivers**: It's generally best practice to be consistent. If some methods need a pointer receiver, make them all pointer receivers.
- **Nil Receivers**: Go allows calling methods on `nil` pointers. You should check for `nil` inside the method if necessary.

## Summary
Methods turn structs into active objects. Choosing between value and pointer receivers is a key decision for both correctness and performance.
