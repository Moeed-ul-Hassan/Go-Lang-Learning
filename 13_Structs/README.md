# Structs in Go

Structs are user-defined types that allow you to group different types of data together. Go is not a traditional Object-Oriented language, but structs provide similar functionality.

## 1. Definition
```go
type User struct {
    Name   string
    Age    int
    Email  string
    Active bool
}
```

## 2. Initialization
```go
// Using field names (Recommended)
u1 := User{
    Name: "Moeed",
    Age:  22,
}

// Positional (Not recommended, breaks if fields change)
u2 := User{"Bob", 30, "bob@example.com", true}

// Zero value
var u3 User // All fields set to their zero values ("", 0, false)
```

## 3. Accessing Fields
Use the dot `.` operator.
```go
u1.Age = 23
fmt.Println(u1.Name)
```

## 4. Anonymous Structs
Useful for one-time data structures (like JSON responses).
```go
data := struct {
    ID   int
    Text string
}{1, "Hello"}
```

## 5. Composition (No Inheritance)
Go uses **struct embedding** to share fields and methods.
```go
type Admin struct {
    User       // Embedded struct
    Privileges []string
}

a := Admin{}
a.Name = "Boss" // Accesses User.Name directly
```

## 6. Exporting (Public vs Private)
- **Uppercase**: `Name` is exported (public) and accessible from other packages.
- **Lowercase**: `age` is unexported (private) to the package.

## Summary
Structs are the foundation of data modeling in Go. They are lightweight, efficient, and encourage composition over complex inheritance hierarchies.
