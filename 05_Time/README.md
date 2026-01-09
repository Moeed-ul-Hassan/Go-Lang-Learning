# Time Package

The `time` package provides functionality for measuring and displaying time.

## Getting Current Time
```go
now := time.Now()
```

## Formatting Time
Go uses a unique reference time for formatting: `Mon Jan 2 15:04:05 MST 2006`.
```go
fmt.Println(now.Format("01-02-2006 Monday"))
```
