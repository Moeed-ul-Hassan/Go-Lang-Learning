# Arrays in Go

Arrays are fixed-size sequences of elements of a single type. In Go, the **size is part of the type**.

## 1. Declaration and Initialization
```go
// Fixed size of 3
var fruits [3]string
fruits[0] = "Apple"

// Inline initialization
numbers := [5]int{10, 20, 30, 40, 50}

// Let Go count the size
colors := [...]string{"Red", "Green", "Blue"}
```

## 2. Key Characteristics
- **Fixed Size**: You cannot change the size of an array after declaration.
- **Value Type**: Assigning an array to a new variable creates a **complete copy**.
- **Homogeneous**: All elements must be of the same type.

## 3. Iterating over Arrays
```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}

// Using range (preferred)
for index, value := range numbers {
    fmt.Printf("Index %d has value %d\n", index, value)
}
```

## 4. Multi-dimensional Arrays
```go
matrix := [2][2]int{
    {1, 2},
    {3, 4},
}
```

## Limitations
Because arrays have a fixed size and are passed by value, they are rarely used directly in Go. Instead, **Slices** are used for almost everything.

## Summary
Arrays are the building blocks for Slices. While you won't use them often, understanding them is crucial for understanding how memory is managed in Go.
