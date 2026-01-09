# Interfaces in Go

Interfaces are a way to define **behavior**. Instead of saying what an object **is** (like a struct), an interface says what an object can **do**.

## Key Concepts

1.  **Implicit Implementation**: You don't need to explicitly say a struct "implements" an interface. If a struct has all the methods defined in an interface, it automatically implements it.
2.  **Decoupling**: Interfaces allow you to write functions that can work with different types as long as they follow the same behavior.
3.  **The Empty Interface `interface{}`**: An interface with zero methods. Every type in Go implements the empty interface, meaning it can hold any value.

## Example
```go
type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

func MakeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}
```

## Why use Interfaces?
- **Flexibility**: You can swap out implementations without changing the code that uses them.
- **Testing**: You can easily create "mock" objects for testing.
- **Clean Code**: It helps in organizing code based on what it does rather than what it is.
