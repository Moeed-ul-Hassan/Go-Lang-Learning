# JSON Handling in Go 🚀

This guide covers the two most important concepts in JSON handling: **Marshalling** and **Unmarshalling**.

## 1. Marshalling (Encoding) 📦
**Definition:** Converting Go data structures (Structs, Maps, Slices) into JSON format (usually `[]byte` or `string`).

### How to use:
Use `json.Marshal` or `json.MarshalIndent` (for pretty-printing).

```go
type Course struct {
    Name  string `json:"coursename"`
    Price int
}

course := Course{"Go Lang", 50}
jsonData, err := json.MarshalIndent(course, "", "\t")
```

### Key Features:
- **Struct Tags:** Use `` `json:"key_name"` `` to control the JSON key names.
- **Ignore Fields:** Use `` `json:"-"` `` to prevent a field from being included in JSON.
- **Omit Empty:** Use `omitempty` to skip fields that are empty/zero.

---

## 2. Unmarshalling (Decoding) 🔓
**Definition:** Converting JSON data back into Go data structures.

### How to use:
Use `json.Unmarshal`. **Always pass a pointer!**

```go
var myCourse Course
err := json.Unmarshal(jsonData, &myCourse)
```

### Why Pointers?
We use `&` because `Unmarshal` needs to modify the original variable. Without the pointer, it would only modify a copy.

### Unmarshaling to Maps:
If you don't know the JSON structure, use a map:
```go
var myData map[string]interface{}
json.Unmarshal(jsonData, &myData)
```
*Note: Numbers in maps are unmarshaled as `float64` by default.*

---

## 3. Best Practices ✅
1. **Check Validity:** Use `json.Valid(data)` before unmarshaling if you're unsure about the source.
2. **Error Handling:** Always check the `err` returned by Marshal/Unmarshal.
3. **Use Structs:** Prefer Structs over Maps for better type safety and performance.
4. **Naming:** Use PascalCase for Go fields (to make them exported/public) and camelCase for JSON keys.

---

## What's Next? ➜
Now that you've mastered JSON, the next logical steps in your Go journey are:
1. **Go Modules (`go mod`):** Learn how to manage project dependencies.
2. **Building APIs:** Create your own REST API using `net/http` and JSON.
3. **External APIs:** Practice by consuming real-world data from public APIs.
