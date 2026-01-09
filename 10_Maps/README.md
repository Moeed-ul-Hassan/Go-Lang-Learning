# Maps in Go

Maps are Go's built-in hash tables. they store data in **key-value pairs** and provide fast lookups.

## 1. Creation
```go
// Using make (preferred)
scores := make(map[string]int)

// Using a literal
colors := map[string]string{
    "red":   "#FF0000",
    "green": "#00FF00",
}
```

## 2. Basic Operations
```go
// Add/Update
scores["Moeed"] = 100

// Delete
delete(scores, "Moeed")

// Length
fmt.Println(len(scores))
```

## 3. The "Comma Ok" Idiom
Checking if a key exists:
```go
val, ok := scores["Bob"]
if ok {
    fmt.Println("Bob's score:", val)
}
```

## 4. Iterating
Maps are **unordered**. Every time you loop, the order might be different.
```go
for key, value := range colors {
    fmt.Printf("Key: %s, Value: %s\n", key, value)
}
```

## 5. Reference Type
Maps are reference types. Passing a map to a function allows that function to modify the original map.

## Important Rules
- **Keys**: Must be of a type that is comparable (strings, ints, floats, pointers). Slices and Maps cannot be keys.
- **Nil Maps**: `var m map[string]int` is `nil`. You can read from it, but **writing to it will cause a panic**. Always use `make`.

## Summary
Maps are the go-to structure for fast data retrieval. Just remember to initialize them with `make` and handle the "not found" case using the comma-ok idiom.
