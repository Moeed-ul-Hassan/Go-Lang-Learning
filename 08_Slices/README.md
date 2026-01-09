# Slices in Go

Slices are the most important and powerful data structure in Go. They are dynamic, flexible views into an underlying array.

## 1. Anatomy of a Slice
A slice consists of three things:
1. **Pointer**: To the first element of the underlying array.
2. **Length**: The number of elements currently in the slice.
3. **Capacity**: The maximum number of elements the slice can hold before reallocating.

## 2. Creation
```go
// Using a literal
s := []int{1, 2, 3}

// Using make (length 5, capacity 10)
s2 := make([]int, 5, 10)

// From an array
arr := [5]int{1, 2, 3, 4, 5}
s3 := arr[1:4] // Elements at index 1, 2, 3
```

## 3. Modifying Slices
### Appending
```go
s = append(s, 4, 5) // Adds elements and returns a new slice
```

### Copying
```go
dest := make([]int, len(s))
copy(dest, s) // Safely copies elements
```

## 4. Slicing Expressions
- `s[low:high]`: From `low` to `high-1`.
- `s[:high]`: From beginning to `high-1`.
- `s[low:]`: From `low` to end.

## 5. Important Behaviors
- **Reference Type**: Slices point to an underlying array. Modifying a slice modifies the original array (unless a re-allocation happens during `append`).
- **Nil Slice**: `var s []int` is `nil` and has length/capacity 0.

## Common Pitfalls
- **Memory Leaks**: A small slice pointing to a huge array keeps the entire array in memory. Use `copy` to a new slice if you only need a small part.
- **Append Side Effects**: If capacity is enough, `append` modifies the original array. If not, it creates a new one.

## Summary
Slices are the "Go way" of handling collections. They provide the performance of arrays with the flexibility of dynamic lists.
