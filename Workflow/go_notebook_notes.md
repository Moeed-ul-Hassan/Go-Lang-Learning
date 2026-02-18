# 📓 Go Engineering: The Senior Notebook

These points are designed to be short, punchy, and easy to copy into your physical notebook for long-term memory.

---

## 1. Professional API Design (Let's Go Further)
*   **Version Everything:** Always prefix routes (e.g., `/v1/`). Never break existing client code.
*   **Graceful Shutdown:** Capture `SIGTERM`. Let active requests finish before the server dies.
*   **Enforced Timeouts:** Every request must have a `context.Context` timeout. Don't let requests hang forever.
*   **Rate Limiting:** Use the "Token Bucket" algorithm to prevent server abuse.
*   **Structured Logs:** Use JSON format (`{"level":"info", "msg":"..."}`). Avoid plain text logs in production.
*   **Centralized Errors:** Create a single helper to handle all API errors ensures consistent JSON responses.

---

## 2. Language Internals (Writing an Interpreter)
*   **Lexical Analysis:** The process of turning raw text into "Tokens" (e.g., `let` -> `TOKEN_LET`).
*   **Abstract Syntax Tree (AST):** A tree representation of code structure. The bridge between text and execution.
*   **The Parser:** The "grammar police." It ensures brackets are closed and syntax is valid.
*   **Zero-Dependency:** True understanding comes from building logic without using 3rd-party libraries.
*   **Recursive Descent:** A common parsing technique where functions call themselves to process nested logic (like nested `if` statements).

---

## 3. Avoiding Pitfalls (100 Go Mistakes)
*   **Slice Leakage:** Slicing a large array keeps the *whole* array in memory. Use `copy()` to save only what you need.
*   **Loop Variable Trap:** In `for` loops, the variable is reused. Be careful when passing its address to goroutines.
*   **Interface Overuse:** Don't start with interfaces. Start with concrete types. Abstract only when you have 2+ implementations.
*   **The `nil` Check:** Checking `if x != nil` on an interface can be tricky if the underlying value is nil but the type isn't.
*   **Init() Poison:** Avoid `init()` functions. They make code hard to test and follow. Use explicit setup functions.

---

## 4. Senior Developer Mindset (Design vs Coding)
*   **System > Code:** A senior architect builds a system (Security, Scaling, Monitoring), not just a feature.
*   **API-First:** The JSON contract is a "Legal Document." Design it before you write the handler.
*   **Layers are Walls:**
    *   *Handlers:* Only speak HTTP.
    *   *Services:* Only speak Business Logic.
    *   *Repositories:* Only speak Database.
*   **The Happy Path is Easy:** Seniors spend 80% of their time planning for the **Unhappy Path** (Errors, Timeouts, Crashes).
*   **Document the "Why":** Comments shouldn't explain *what* the code does. They should explain the *decision* behind it.

---

## 5. Professional Standards (Official Docs)
*   **Gofmt:** There is no debate on style. If `gofmt` likes it, it's correct.
*   **Effective Go:** The rulebook for "Idiomatic" Go. If it feels weird in Go, it's probably wrong.
*   **Go Modules:** `go.mod` is the source of truth for dependencies. Keep it tidy with `go mod tidy`.
*   **Internal Folders:** Use the `/internal` directory for code that should NOT be imported by other projects.

---

## 6. Advanced Architecture (The Greenlight Project)
*   **Response Enveloping:** Wrap JSON in a parent key (e.g., `{"movie": {...}}`). It adds security and structure for clients.
*   **Dependency Hijacking:** Use a central `application` struct. Attach handlers to it as methods to share DB and Logger without global variables.
*   **Optimistic Concurrency:** Use a `version` field in SQL. Only update if the version matches the one the client read. Prevents "lost updates."
*   **SQL Migrations:** Use "Up/Down" scripts (e.g., `golang-migrate`) to track DB changes over time. Never manually edit live production schemas.
*   **Connection Pooling:** Tuned via `MaxOpenConns` and `MaxIdleConns`. Prevents the API from "choking" under high traffic.
*   **Token Bucket Middleware:** Use Go's `x/time/rate` or custom mutex-logic to limit requests per IP.
*   **Expvar Monitoring:** Use the `expvar` package to expose live memory and DB stats to Prometheus/Grafana.
*   **Build Metadata:** Use `-ldflags` during `go build` to inject the Git Commit Hash and Build Date into the binary.
*   **Process Management:** Use `systemd` on Linux to ensure the server auto-restarts after a crash or reboot.
*   **Reverse Proxy:** Place the API behind **Caddy** or **Nginx**. This handles SSL (HTTPS) and provides an extra layer of security.
