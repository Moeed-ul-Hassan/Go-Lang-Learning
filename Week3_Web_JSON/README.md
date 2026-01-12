# Week 3: Web & JSON Basics

This directory contains the Go exercises for Week 3, covering HTTP server basics, URL routing, GET/POST handling, form data, and JSON serialization.

---

## 💡 Beginner's Guide to Clean Go Code

As a beginner, you might find yourself writing `if err != nil { panic(err) }` everywhere. Here are better ways to handle errors and keep your code clean.

### 1. The "Guard Clause" (Check Early, Stop Early)
Instead of nesting your "good" code inside a big `if` block, check for the error first and `return` immediately. This keeps your main logic "aligned to the left."

**❌ Instead of this (Nested):**
```go
func handle(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err == nil {
        // All your logic is stuck inside here...
        fmt.Fprint(w, "Success!")
    }
}
```

**✅ Do this (Guard Clause):**
```go
func handle(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form", 400)
        return // Stop right here!
    }

    // Now the "Happy Path" is clean and easy to read
    fmt.Fprint(w, "Success!")
}
```

### 2. The "Must" Pattern (For Critical Setup)
If something **must** work for your program to even start (like loading a config file), you can use a helper function to avoid repeating the `panic` logic.

**✅ Example Helper:**
```go
func Must[T any](val T, err error) T {
    if err != nil {
        panic(err) // Only panic if it's a fatal startup error
    }
    return val
}

// Usage:
content := Must(os.ReadFile("config.txt"))
```

### 3. Inline Error Checking
You can declare a variable and check it in the same line. This keeps the variable `err` only inside that specific block, making your code safer and tidier.

```go
if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
    return err
}
```

### 4. Why not use `panic` everywhere?
In Go, `panic` is like a "nuclear option"—it crashes the whole program.
- **Use `panic`**: Only when the program *cannot* continue (e.g., database won't connect at startup).
- **Use `return err`**: In almost all other cases (e.g., a user sent a bad password, or a file is missing). You want your server to stay alive even if one request fails!

---

## 🚀 How the Web Server Works (Walkthrough)

Since you are a beginner, here is a breakdown of what our `main.go` is actually doing:

### 1. The "Traffic Controller" (Mux)
```go
mux := http.NewServeMux()
```
Think of the **Mux** (Multiplexer) as a receptionist. When someone visits your server, the Mux looks at the URL and decides which function should handle the request.

### 2. Mapping URLs to Functions
```go
mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/hello/", helloHandler)
```
This tells the Mux: "If someone goes to `/`, call the `homeHandler` function."

### 3. Starting the Server
```go
log.Fatal(http.ListenAndServe(":8080", mux))
```
This line actually starts the server on your computer at port `8080`. The `log.Fatal` is there to catch errors if the server fails to start (e.g., if port 8080 is already being used).

### 4. What are "Handlers"?
Handlers are just functions that take two things:
1.  `w http.ResponseWriter`: This is like a blank piece of paper you write your response on to send back to the user.
2.  `r *http.Request`: This contains all the information the user sent to you (the URL, any data, their browser info).

### 5. Our Routes:
- **`/`**: Just says "Hello World".
- **`/hello/Name`**: Takes the name from the URL and greets you.
- **`/search?q=Term`**: Reads a "query parameter" (the `?q=...` part) and returns it as JSON.
- **`/echo`**: A **POST** request. It takes JSON data you send and sends it right back to you.
- **`/submit`**: Handles data sent from an HTML form.
- **`/data`**: Sends back a list of courses in JSON format.
