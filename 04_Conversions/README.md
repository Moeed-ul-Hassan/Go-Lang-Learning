# Type Conversions

Go requires explicit type conversion. You cannot implicitly convert between types like `int` and `float64`.

## Example
```go
var myInt int = 42
var myFloat float64 = float64(myInt)
```

## String Conversions
Use the `strconv` package for converting strings to numbers and vice versa.
```go
i, _ := strconv.Atoi("123")
s := strconv.Itoa(123)
```
