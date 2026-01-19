# 🐹 Go Syntax Cheatsheet

A quick reference guide for Go (Golang) syntax.

## 📦 Basics

### Variables
```go
var x int = 10          // Explicit
y := 20                 // Short declaration (inside functions only)
var a, b = 1, "hello"   // Multiple
const Pi = 3.14         // Constant
```

### Basic Types
```go
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32 (Unicode code point)
float32 float64
complex64 complex128
```

### Type Conversion
```go
i := 42
f := float64(i)
u := uint(f)
```

---

## 🔄 Control Flow

### If / Else
```go
if x > 10 {
    fmt.Println("Big")
} else if x == 10 {
    fmt.Println("Ten")
} else {
    fmt.Println("Small")
}

// With initialization
if err := doSomething(); err != nil {
    log.Fatal(err)
}
```

### Switch
```go
switch os {
case "darwin":
    fmt.Println("macOS")
case "linux":
    fmt.Println("Linux")
default:
    fmt.Println("Other")
}

// Tagless switch (like if-else chain)
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Morning")
case t.Hour() < 17:
    fmt.Println("Afternoon")
}
```

### For Loop (The only loop in Go)
```go
// Standard C-style
for i := 0; i < 10; i++ {
    sum += i
}

// While-style
for sum < 1000 {
    sum += sum
}

// Infinite
for {
    // loop forever
}

// Range (Slices, Maps, Strings)
for index, value := range mySlice {
    fmt.Printf("%d: %s\n", index, value)
}
```

---

## 🏗️ Data Structures

### Arrays & Slices
```go
var a [10]int           // Array (Fixed size)
primes := [6]int{2, 3, 5, 7, 11, 13}

// Slice (Dynamic size)
var s []int = primes[1:4]
s := []int{1, 2, 3}
s = append(s, 4, 5)     // Add elements
len(s)                  // Length
cap(s)                  // Capacity
make([]int, 5, 10)      // Create with len 5, cap 10
```

### Maps
```go
m := make(map[string]int)
m["Answer"] = 42
v, ok := m["Answer"]    // Check existence
delete(m, "Answer")     // Delete key
```

### Structs
```go
type Vertex struct {
    X, Y int
}

v := Vertex{1, 2}
v.X = 4
p := &v                 // Pointer to struct
p.X = 1e9               // Access field via pointer (auto-dereference)
```

---

## 🔧 Functions

### Basic
```go
func add(x int, y int) int {
    return x + y
}
```

### Multiple Return Values
```go
func swap(x, y string) (string, string) {
    return y, x
}
```

### Named Return Values
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // Naked return
}
```

### Variadic Functions
```go
func sum(nums ...int) {
    // nums is a slice []int
}
```

### Closures
```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

### Defer
```go
func main() {
    defer fmt.Println("world") // Runs when function returns
    fmt.Println("hello")
}
```

---

## 👉 Pointers

```go
i, j := 42, 2701
p := &i         // Point to i
*p = 21         // Set i through the pointer
fmt.Println(i)  // See the new value of i
```

---

## 🧬 Methods & Interfaces

### Methods
```go
type Vertex struct {
    X, Y float64
}

// Receiver (v Vertex)
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pointer Receiver (can modify v)
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
```

### Interfaces
```go
type Abser interface {
    Abs() float64
}

// Implicit implementation (no "implements" keyword)
```

---

## ⚡ Concurrency

### Goroutines
```go
go doSomething() // Runs in a new goroutine
```

### Channels
```go
ch := make(chan int)    // Create channel
ch <- v                 // Send v to channel ch
v := <-ch               // Receive from ch, and assign to v
close(ch)               // Close channel

// Buffered Channel
ch := make(chan int, 100)
```

### Select
```go
select {
case i := <-c:
    // use i
case <-quit:
    return
default:
    // receiving from c would block
}
```

---

## 🚨 Error Handling

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
