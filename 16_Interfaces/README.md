# Interfaces in Go

Interfaces are Go's way of defining **contracts** or **behavior**. They are one of the most powerful features for writing flexible, decoupled code.

## 1. Definition
An interface is a set of method signatures.
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

## 2. Implicit Implementation
In Go, you don't use an `implements` keyword. A type implements an interface simply by having all the required methods.

```go
type Circle struct { Radius float64 }

// Circle now implements Shape automatically
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
```

## 3. The Empty Interface `interface{}`
An interface with no methods. Every type implements it.
```go
var x interface{}
x = 42      // OK
x = "hello" // OK
```
*Note: In modern Go (1.18+), `any` is an alias for `interface{}`.*

## 4. Type Assertions and Switches
To get the concrete type back from an interface:

### Assertion
```go
s, ok := myInterface.(string)
```

### Type Switch
```go
switch v := myInterface.(type) {
case int:
    fmt.Println("It's an int:", v)
case string:
    fmt.Println("It's a string:", v)
}
```

## 5. Interface Best Practices
- **Keep them small**: Interfaces with 1-3 methods are common (e.g., `io.Reader`, `fmt.Stringer`).
- **Accept interfaces, return structs**: This provides maximum flexibility for the caller.
- **Don't over-engineer**: Only create an interface when you actually have multiple implementations or need to decouple code.

## Summary
Interfaces allow you to focus on what a type **does** rather than what it **is**. They are the key to building large, maintainable systems in Go.
