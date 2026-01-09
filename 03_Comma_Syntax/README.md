# Comma Ok Syntax

In Go, the "comma ok" idiom is used to check if a key exists in a map or if a type assertion was successful.

## Maps
```go
value, ok := myMap["key"]
if ok {
    // Key exists
}
```

## Type Assertions
```go
str, ok := myInterface.(string)
if ok {
    // Successfully asserted to string
}
```
