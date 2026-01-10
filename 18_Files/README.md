# Working with Files in Go

Go provides the `os` and `io` packages to handle file operations.

## 1. Creating and Writing to a File
Use `os.Create` to create a file and `io.WriteString` or `f.WriteString` to write to it.

```go
file, err := os.Create("./test.txt")
if err != nil {
    panic(err)
}
defer file.Close() // Always defer closing!

content := "Hello Go!"
len, err := io.WriteString(file, content)
```

## 2. Reading a File
The easiest way to read a whole file is using `os.ReadFile`.

```go
data, err := os.ReadFile("./test.txt")
if err != nil {
    panic(err)
}
fmt.Println(string(data))
```

## 3. Best Practices
- **Always check for errors**: File operations can fail (permission denied, disk full, etc.).
- **Always defer `file.Close()`**: This ensures the file is closed even if the program crashes later.
- **Use `os.ReadFile` for small files**: For very large files, use a `Scanner` or `Reader` to read line-by-line.
