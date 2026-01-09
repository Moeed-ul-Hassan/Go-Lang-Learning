# Slices

Slices are dynamic, flexible views into the elements of an array.

## Creation
```go
mySlice := []int{1, 2, 3}
mySlice = append(mySlice, 4)
```

## Slicing
```go
subSlice := mySlice[1:3] // Elements from index 1 to 2
```

## Key Functions
- `append()`: Adds elements to a slice.
- `make()`: Creates a slice with a specific length and capacity.
- `copy()`: Copies elements from one slice to another.
