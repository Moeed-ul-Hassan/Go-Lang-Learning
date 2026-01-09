# Maps

Maps are Go's built-in associative data type (hash tables).

## Creation
```go
colors := make(map[string]string)
colors["red"] = "#FF0000"
```

## Deletion
```go
delete(colors, "red")
```

## Iteration
```go
for key, value := range colors {
    fmt.Println(key, value)
}
```
