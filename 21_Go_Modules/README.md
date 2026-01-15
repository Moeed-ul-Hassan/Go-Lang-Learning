# Go Modules (MOD) 🏎️

Go Modules are the official dependency management system for Go. They allow you to manage third-party packages and ensure reproducible builds.

## Key Commands 🛠️

1.  **Initialize a Module**:
    ```bash
    go mod init <module-name>
    ```
    Example: `go mod init github.com/Moeed-ul-Hassan/Go-Modules`
    This creates a `go.mod` file.

2.  **Tidy Dependencies**:
    ```bash
    go mod tidy
    ```
    This command cleans up your `go.mod` and `go.sum` files. It downloads any missing dependencies and removes unused ones.

3.  **Add a Dependency**:
    ```bash
    go get -u <package-url>
    ```
    Example: `go get -u github.com/gorilla/mux`

4.  **Verify Dependencies**:
    ```bash
    go mod verify
    ```
    Checks if the dependencies on your machine match the ones in `go.sum`.

## The `go.mod` File 📄
This file contains:
- The module path.
- The Go version.
- A list of required dependencies with their versions.

## The `go.sum` File 📄
This file contains the expected cryptographic hashes of the content of specific module versions. It ensures that the dependencies haven't been tampered with.

---
**Next Step**: Let's initialize this folder as a module!
Run: `go mod init modules_learning`
