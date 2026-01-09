# Type Conversions in Go

Go is a **strongly typed** and **statically typed** language. Unlike some languages (like JavaScript or C++), Go **never** performs implicit type conversion. You must always be explicit.

## 1. Numeric Conversions
You cannot perform operations between different numeric types (e.g., `int` + `float64`).

### Explicit Casting
```go
var myInt int = 10
var myFloat float64 = 3.5

// result := myInt + myFloat // ERROR: invalid operation
result := float64(myInt) + myFloat // OK: 13.5
```

## 2. String Conversions (The `strconv` Package)
Converting between strings and numbers requires the `strconv` package.

### String to Integer (`Atoi`)
```go
str := "100"
val, err := strconv.Atoi(str)
if err != nil {
    fmt.Println("Error converting string to int:", err)
}
```

### Integer to String (`Itoa`)
```go
val := 42
str := strconv.Itoa(val)
```

### Parsing Floats and Bools
```go
f, _ := strconv.ParseFloat("3.14", 64)
b, _ := strconv.ParseBool("true")
```

## 3. String and Byte Slices
Converting between strings and byte slices is common for I/O operations.
```go
s := "Hello"
b := []byte(s) // Convert string to []byte
s2 := string(b) // Convert []byte back to string
```

## Key Rules
- **No Implicit Conversion**: Even `int32` and `int64` are different types and require conversion.
- **Precision Loss**: Converting from `float64` to `int` will truncate the decimal part (e.g., `int(3.9)` becomes `3`).
- **Error Handling**: Always check for errors when using `strconv` functions.

## Summary
Explicit conversions make Go code slightly more verbose but significantly **safer** by preventing hidden bugs caused by automatic type changes.
