# Go Methods: Value vs Pointer Receivers

In Go, a method is a function with a special **receiver** argument. The receiver can be either a **Value** or a **Pointer**.

## 1. Value Receiver `(u user)`
When you use a value receiver, Go creates a **copy** of the struct and passes it to the method.

- **Behavior**: Any changes made to the struct inside the method **do not** affect the original variable.
- **Use Case**: Use this when you only need to **read** data or if the struct is very small and you don't want to modify it.

```go
func (u user) greet() {
    fmt.Println("Hello,", u.name)
}
```

## 2. Pointer Receiver `(u *user)`
When you use a pointer receiver, Go passes the **memory address** of the struct.

- **Behavior**: Any changes made inside the method **will modify** the original variable in memory.
- **Use Case**: 
    1. When you need to **update/modify** the data.
    2. For **performance** with large structs (to avoid copying large amounts of data).

```go
func (u *user) incrementAge() {
    u.age++ // Modifies the actual user in main()
}
```

## Comparison Table

| Feature | Value Receiver `(u user)` | Pointer Receiver `(u *user)` |
| :--- | :--- | :--- |
| **Data** | Works on a **copy** | Works on the **original** |
| **Performance** | Slower for large structs (copying) | Faster (passes memory address) |
| **Modification** | Cannot modify original | **Can modify original** |

## Summary
- **Pointer Receiver (`*`)** = Pointing to the actual variable (Changes stay).
- **Value Receiver** = Taking a snapshot (Changes are lost).
