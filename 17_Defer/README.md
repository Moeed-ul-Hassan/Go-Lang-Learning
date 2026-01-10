# Defer in Go

`defer` is a keyword used to delay the execution of a function until the surrounding function returns.

## 1. How it works
When you use `defer`, the function call is pushed onto a **stack**. When the main function finishes, the deferred calls are executed in **Last-In, First-Out (LIFO)** order.

```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
// Output:
// Hello
// World
```

## 2. Multiple Defers (LIFO)
If you have multiple defers, they execute in reverse order.
```go
func main() {
    defer fmt.Println("One")
    defer fmt.Println("Two")
    defer fmt.Println("Three")
}
// Output:
// Three
// Two
// One
```

## 3. Why use it?
- **Cleaning up resources**: Closing files, closing database connections.
- **Unlocking Mutexes**: Ensuring a lock is released even if the function panics.
- **Logging**: Recording when a function finishes.

## 4. Important Note
The arguments to a deferred function are evaluated **immediately**, but the function itself doesn't run until the end.
