# Structs

Structs are collections of fields used to group data together.

## Definition
```go
type User struct {
    Name  string
    Age   int
    Email string
}
```

## Initialization
```go
u := User{Name: "Moeed", Age: 22}
```

## Embedded Structs (Composition)
Go uses composition instead of inheritance.
```go
type Admin struct {
    User
    Role string
}
```
