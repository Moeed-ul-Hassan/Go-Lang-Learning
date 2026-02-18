# Writing Clean Go Code: The Complete Guide

> **What does "Clean Code" even mean?**  
> Clean code is code that is easy to read, easy to understand, easy to test, and easy to maintain. It's not just about making code work—it's about making code that other developers (including future you) can understand and modify without fear.

---

## 🎯 Core Principles of Clean Code

**Clean code should:**
- ✅ Be self-explanatory (minimal comments needed)
- ✅ Do one thing well (Single Responsibility Principle)
- ✅ Be easily testable
- ✅ Have meaningful names
- ✅ Be formatted consistently
- ✅ Handle errors explicitly
- ✅ Be performant without premature optimization

---

# 30 Rules for Writing Clean, Fast, and Efficient Go Code

## 📝 Section 1: Naming Conventions (Rules 1-6)

### Rule 1: Use Short, Descriptive Names for Local Variables
```go
// ❌ Bad
numberOfUsersInDatabase := 10

// ✅ Good
userCount := 10

// ✅ Good for loop variables
for i, user := range users {
    // i is clear in context
}
```

**Why:** Go favors brevity. Short names are fine when scope is small.

---

### Rule 2: Use Longer Names for Package-Level and Exported Functions
```go
// ❌ Bad
func Get(id int) {}

// ✅ Good
func GetUserByID(id int) (*User, error) {}

// ✅ Good
func CalculateMonthlyRevenue(startDate, endDate time.Time) (float64, error) {}
```

**Why:** Exported names need to be clear without context.

---

### Rule 3: Use MixedCaps (camelCase/PascalCase), Never snake_case
```go
// ❌ Bad
var user_count int
func get_user_name() string {}

// ✅ Good
var userCount int
func GetUserName() string {}
```

**Why:** Go convention. Exported names start with capital, unexported with lowercase.

---

### Rule 4: Avoid Stuttering in Names
```go
// ❌ Bad
user.UserName
user.GetUserAge()

// ✅ Good
user.Name
user.Age()
```

**Why:** The context (user) already tells us it's about a user.

---

### Rule 5: Use Meaningful Variable Names, Not Single Letters (Except Conventions)
```go
// ❌ Bad
func f(a, b int) int { return a + b }

// ✅ Good
func Add(x, y int) int { return x + y }

// ✅ Acceptable conventions
for i := range items {}        // i for index
for k, v := range map {}        // k, v for key-value
var mu sync.Mutex               // mu for mutex
var wg sync.WaitGroup           // wg for wait group
```

---

### Rule 6: Use Consistent Naming Across Codebase
```go
// ❌ Bad - inconsistent
func GetUser() {}
func FetchProduct() {}
func RetrieveOrder() {}

// ✅ Good - consistent
func GetUser() {}
func GetProduct() {}
func GetOrder() {}
```

---

## 🏗️ Section 2: Function Design (Rules 7-12)

### Rule 7: Functions Should Do One Thing
```go
// ❌ Bad - does too much
func ProcessUser(id int) error {
    user := getUser(id)
    validateUser(user)
    saveToDatabase(user)
    sendEmail(user)
    logActivity(user)
    return nil
}

// ✅ Good - single responsibility
func GetUser(id int) (*User, error) {
    // Only fetch user
}

func ValidateUser(user *User) error {
    // Only validate
}

func SaveUser(user *User) error {
    // Only save
}
```

---

### Rule 8: Keep Functions Short (Ideally < 50 Lines)
```go
// ❌ Bad - 100+ lines function
func HandleRequest(w http.ResponseWriter, r *http.Request) {
    // 100+ lines of logic
}

// ✅ Good - break it down
func HandleRequest(w http.ResponseWriter, r *http.Request) {
    user, err := parseUser(r)
    if err != nil {
        handleError(w, err)
        return
    }
    
    if err := validateUser(user); err != nil {
        handleError(w, err)
        return
    }
    
    if err := saveUser(user); err != nil {
        handleError(w, err)
        return
    }
    
    sendSuccessResponse(w, user)
}
```

---

### Rule 9: Return Early to Avoid Deep Nesting
```go
// ❌ Bad - deep nesting
func Process(data string) error {
    if data != "" {
        if len(data) > 10 {
            if isValid(data) {
                // do something
                return nil
            } else {
                return errors.New("invalid")
            }
        } else {
            return errors.New("too short")
        }
    } else {
        return errors.New("empty")
    }
}

// ✅ Good - return early
func Process(data string) error {
    if data == "" {
        return errors.New("empty")
    }
    
    if len(data) <= 10 {
        return errors.New("too short")
    }
    
    if !isValid(data) {
        return errors.New("invalid")
    }
    
    // do something
    return nil
}
```

---

### Rule 10: Accept Interfaces, Return Concrete Types
```go
// ❌ Bad
func NewReader(r io.Reader) io.Reader {
    return bufio.NewReader(r)
}

// ✅ Good
func NewReader(r io.Reader) *bufio.Reader {
    return bufio.NewReader(r)
}
```

**Why:** Gives callers more flexibility and type information.

---

### Rule 11: Use Named Return Values for Documentation
```go
// ✅ Good - clear what's returned
func Divide(a, b float64) (result float64, err error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Also acceptable for simple cases
func Add(a, b int) int {
    return a + b
}
```

**Use named returns when:** they improve clarity (especially with multiple returns).

---

### Rule 12: Limit Function Parameters (Max 3-4)
```go
// ❌ Bad - too many parameters
func CreateUser(name, email, phone, address, city, country, zip string, age int) {}

// ✅ Good - use a struct
type UserInput struct {
    Name    string
    Email   string
    Phone   string
    Address Address
    Age     int
}

func CreateUser(input UserInput) error {}
```

---

## 🔧 Section 3: Error Handling (Rules 13-16)

### Rule 13: Always Handle Errors Explicitly
```go
// ❌ Bad - ignoring errors
file, _ := os.Open("file.txt")

// ✅ Good
file, err := os.Open("file.txt")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
defer file.Close()
```

---

### Rule 14: Use Error Wrapping for Context
```go
// ❌ Bad - losing context
func SaveUser(user *User) error {
    if err := db.Save(user); err != nil {
        return err
    }
    return nil
}

// ✅ Good - adding context
func SaveUser(user *User) error {
    if err := db.Save(user); err != nil {
        return fmt.Errorf("saving user %s: %w", user.ID, err)
    }
    return nil
}
```

**Use `%w` to wrap errors** - allows `errors.Is()` and `errors.As()` to work.

---

### Rule 15: Create Custom Errors for Domain Logic
```go
// ✅ Good
var (
    ErrUserNotFound     = errors.New("user not found")
    ErrInvalidEmail     = errors.New("invalid email format")
    ErrDuplicateUser    = errors.New("user already exists")
)

func GetUser(id string) (*User, error) {
    user, found := db.Get(id)
    if !found {
        return nil, ErrUserNotFound
    }
    return user, nil
}

// Caller can check
if errors.Is(err, ErrUserNotFound) {
    // handle not found
}
```

---

### Rule 16: Check Errors Before defer Close()
```go
// ❌ Bad - might panic
defer file.Close()

// ✅ Good
if file != nil {
    defer file.Close()
}

// ✅ Even better - handle close errors for writes
defer func() {
    if err := file.Close(); err != nil {
        log.Printf("failed to close file: %v", err)
    }
}()
```

---

## 🧩 Section 4: Code Organization (Rules 17-21)

### Rule 17: Group Related Code Together
```go
// ❌ Bad - scattered
type User struct { Name string }
type Product struct { Title string }
func GetUser() {}
type Order struct { ID int }
func GetProduct() {}

// ✅ Good - grouped
// User domain
type User struct { Name string }
func GetUser() {}
func SaveUser() {}

// Product domain
type Product struct { Title string }
func GetProduct() {}
func SaveProduct() {}
```

---

### Rule 18: Order Struct Fields Logically
```go
// ❌ Bad - random order
type User struct {
    UpdatedAt time.Time
    Name      string
    ID        int
    Email     string
    CreatedAt time.Time
}

// ✅ Good - logical order (ID first, timestamps last)
type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

---

### Rule 19: Use Package-Level Variables Sparingly
```go
// ❌ Bad - mutable global state
package app

var Config map[string]string

func UpdateConfig(key, value string) {
    Config[key] = value
}

// ✅ Good - encapsulate
package config

type Config struct {
    values map[string]string
    mu     sync.RWMutex
}

func (c *Config) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.values[key]
}
```

---

### Rule 20: Keep Package Names Short and Clear
```go
// ❌ Bad package names
package utilityFunctionsForUsers
package helpers
package common

// ✅ Good package names
package user
package auth
package http
```

---

### Rule 21: Organize Imports in Groups
```go
// ✅ Good - standard, external, internal
import (
    // Standard library
    "fmt"
    "time"
    
    // External packages
    "github.com/gin-gonic/gin"
    "github.com/gorilla/mux"
    
    // Internal packages
    "myapp/models"
    "myapp/services"
)
```

---

## 🚀 Section 5: Performance & Efficiency (Rules 22-26)

### Rule 22: Use Pointers for Large Structs
```go
// ❌ Bad - copies 1KB struct
type LargeStruct struct {
    Data [1000]byte
}

func Process(s LargeStruct) {}  // Copies all data

// ✅ Good - passes pointer
func Process(s *LargeStruct) {}  // Only copies 8 bytes
```

**Rule of thumb:** Use pointers for structs > 64 bytes.

---

### Rule 23: Preallocate Slices When Size is Known
```go
// ❌ Bad - multiple allocations
var users []User
for i := 0; i < 1000; i++ {
    users = append(users, User{ID: i})
}

// ✅ Good - single allocation
users := make([]User, 0, 1000)
for i := 0; i < 1000; i++ {
    users = append(users, User{ID: i})
}
```

---

### Rule 24: Use string.Builder for String Concatenation
```go
// ❌ Bad - creates new string each time (slow)
var result string
for _, s := range strings {
    result += s
}

// ✅ Good - efficient
var builder strings.Builder
for _, s := range strings {
    builder.WriteString(s)
}
result := builder.String()
```

---

### Rule 25: Avoid Unnecessary Pointer Dereferences
```go
// ❌ Bad
func (u *User) GetName() string {
    return (*u).Name  // Unnecessary
}

// ✅ Good - Go dereferences automatically
func (u *User) GetName() string {
    return u.Name
}
```

---

### Rule 26: Use sync.Pool for Frequently Allocated Objects
```go
// ✅ Good - reuse expensive objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func ProcessData(data []byte) {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()
    
    buf.Write(data)
    // use buffer
}
```

---

## 🧪 Section 6: Testing & Code Quality (Rules 27-30)

### Rule 27: Write Table-Driven Tests
```go
// ✅ Good
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed", -2, 3, 1},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("Add(%d, %d) = %d, want %d", 
                    tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

---

### Rule 28: Design for Testability - Use Interfaces
```go
// ❌ Bad - hard to test
type UserService struct {}

func (s *UserService) GetUser(id int) (*User, error) {
    // directly calls database
    return db.Query("SELECT * FROM users WHERE id = ?", id)
}

// ✅ Good - easy to mock
type UserRepository interface {
    GetByID(id int) (*User, error)
}

type UserService struct {
    repo UserRepository
}

func (s *UserService) GetUser(id int) (*User, error) {
    return s.repo.GetByID(id)
}
```

---

### Rule 29: Use Context for Cancellation and Timeouts
```go
// ✅ Good
func FetchData(ctx context.Context, url string) ([]byte, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    return io.ReadAll(resp.Body)
}

// Usage with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

data, err := FetchData(ctx, "https://api.example.com")
```

---

### Rule 30: Run go vet and golint Regularly
```bash
# Check for common mistakes
go vet ./...

# Check for style issues
golangci-lint run

# Format code
go fmt ./...
```

---

## 📚 How to Write Documentation for Go Code

### 1. Package Documentation
```go
// Package user provides functionality for user management.
//
// This package handles user creation, authentication, and profile management.
// It integrates with the database layer for persistence and the auth package
// for security features.
//
// Basic usage:
//
//     svc := user.NewService(db, logger)
//     user, err := svc.CreateUser(ctx, "john@example.com", "password")
//     if err != nil {
//         // handle error
//     }
//
package user
```

**Best practices:**
- First sentence is a summary (shows in godoc)
- Include usage examples
- Mention dependencies/integrations

---

### 2. Function Documentation
```go
// CreateUser creates a new user account with the given credentials.
//
// It validates the email format, checks for duplicate accounts, hashes the
// password using bcrypt, and stores the user in the database.
//
// Returns ErrInvalidEmail if the email format is invalid.
// Returns ErrDuplicateUser if a user with the same email exists.
// Returns any database errors wrapped with context.
//
// Example:
//
//     user, err := CreateUser(ctx, "john@example.com", "secure123")
//     if errors.Is(err, ErrInvalidEmail) {
//         // handle invalid email
//     }
//
func CreateUser(ctx context.Context, email, password string) (*User, error) {
    // implementation
}
```

**Best practices:**
- Start with function name
- Explain what it does (not how)
- Document parameters if not obvious
- Document all error cases
- Include examples for complex functions

---

### 3. Type Documentation
```go
// User represents a registered user in the system.
//
// Users have a unique ID and email. The password is stored as a bcrypt hash
// and should never be exposed in API responses.
type User struct {
    // ID is the unique identifier for the user.
    ID int `json:"id"`
    
    // Email must be a valid email address and is unique across all users.
    Email string `json:"email"`
    
    // PasswordHash is the bcrypt hash of the user's password.
    // This field should never be serialized to JSON.
    PasswordHash string `json:"-"`
    
    // CreatedAt is when the user account was created.
    CreatedAt time.Time `json:"created_at"`
}
```

**Best practices:**
- Document the purpose of the type
- Document non-obvious fields
- Explain validation rules
- Note fields that shouldn't be exposed

---

### 4. Example Tests (Executable Documentation)
```go
// ExampleAdd demonstrates basic usage of the Add function.
func ExampleAdd() {
    result := Add(2, 3)
    fmt.Println(result)
    // Output: 5
}

// ExampleUserService_CreateUser shows how to create a user.
func ExampleUserService_CreateUser() {
    svc := NewUserService(db, logger)
    
    user, err := svc.CreateUser(context.Background(), "john@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Created user: %s\n", user.Email)
    // Output: Created user: john@example.com
}
```

**Best practices:**
- Example tests show up in godoc
- Use `// Output:` to make them testable
- Keep examples simple and focused

---

### 5. README.md for Packages
```markdown
# User Package

User management and authentication for the application.

## Features

- User registration with email/password
- Password hashing with bcrypt
- JWT token generation
- User profile management

## Installation

```bash
go get github.com/yourorg/yourapp/user
```

## Quick Start

```go
import "github.com/yourorg/yourapp/user"

// Create service
svc := user.NewService(db, logger)

// Create user
user, err := svc.CreateUser(ctx, "email@example.com", "password")
if err != nil {
    // handle error
}
```

## API Reference

See [GoDoc](https://pkg.go.dev/github.com/yourorg/yourapp/user) for full API documentation.

## Testing

```bash
go test ./...
```
```

---

## 🛠️ Documentation Tools

### 1. Generate Documentation with godoc
```bash
# Install godoc
go install golang.org/x/tools/cmd/godoc@latest

# View documentation locally
godoc -http=:6060

# Visit http://localhost:6060/pkg/yourpackage
```

---

### 2. Use golangci-lint for Documentation Checks
```yaml
# .golangci.yml
linters:
  enable:
    - godot        # Check comments end with period
    - godox        # Detect FIXME, TODO, etc
    - misspell     # Check for spelling mistakes
```

---

### 3. Documentation Checklist

For every exported symbol:
- [ ] Has a comment starting with its name
- [ ] Comment is a complete sentence with period
- [ ] Explains what it does (not how)
- [ ] Documents error cases
- [ ] Includes examples for complex usage
- [ ] Uses consistent terminology

---

## 📖 Complete Example: Clean, Documented Code

```go
// Package order handles order processing for the e-commerce system.
//
// It provides functionality for creating, updating, and tracking orders.
// Orders go through a state machine: pending -> confirmed -> shipped -> delivered.
package order

import (
    "context"
    "errors"
    "fmt"
    "time"
)

// Common errors returned by this package.
var (
    ErrOrderNotFound    = errors.New("order not found")
    ErrInvalidStatus    = errors.New("invalid order status")
    ErrInsufficientStock = errors.New("insufficient stock")
)

// Status represents the current state of an order.
type Status string

const (
    StatusPending   Status = "pending"
    StatusConfirmed Status = "confirmed"
    StatusShipped   Status = "shipped"
    StatusDelivered Status = "delivered"
)

// Order represents a customer order.
//
// Orders track items, quantities, and fulfillment status.
// The total is calculated from line items and stored for performance.
type Order struct {
    // ID is the unique order identifier.
    ID string
    
    // CustomerID references the customer who placed the order.
    CustomerID string
    
    // Items contains the products and quantities ordered.
    Items []LineItem
    
    // Total is the order total in cents (e.g., $10.50 = 1050).
    Total int
    
    // Status tracks the fulfillment state.
    Status Status
    
    // CreatedAt is when the order was placed.
    CreatedAt time.Time
}

// LineItem represents a single product in an order.
type LineItem struct {
    ProductID string
    Quantity  int
    Price     int // Price in cents at time of order
}

// Repository defines storage operations for orders.
type Repository interface {
    // Get retrieves an order by ID.
    // Returns ErrOrderNotFound if the order doesn't exist.
    Get(ctx context.Context, id string) (*Order, error)
    
    // Save persists an order to storage.
    Save(ctx context.Context, order *Order) error
}

// Service handles order business logic.
type Service struct {
    repo  Repository
    stock StockChecker
}

// StockChecker verifies product availability.
type StockChecker interface {
    CheckAvailability(ctx context.Context, productID string, qty int) (bool, error)
}

// NewService creates a new order service.
func NewService(repo Repository, stock StockChecker) *Service {
    return &Service{
        repo:  repo,
        stock: stock,
    }
}

// CreateOrder creates a new order in pending status.
//
// It validates stock availability for all items before creating the order.
// The total is calculated from line items.
//
// Returns ErrInsufficientStock if any item is out of stock.
// Returns any repository errors wrapped with context.
func (s *Service) CreateOrder(ctx context.Context, customerID string, items []LineItem) (*Order, error) {
    // Validate stock for all items
    for _, item := range items {
        available, err := s.stock.CheckAvailability(ctx, item.ProductID, item.Quantity)
        if err != nil {
            return nil, fmt.Errorf("checking stock for %s: %w", item.ProductID, err)
        }
        if !available {
            return nil, ErrInsufficientStock
        }
    }
    
    // Calculate total
    total := 0
    for _, item := range items {
        total += item.Price * item.Quantity
    }
    
    // Create order
    order := &Order{
        ID:         generateID(),
        CustomerID: customerID,
        Items:      items,
        Total:      total,
        Status:     StatusPending,
        CreatedAt:  time.Now(),
    }
    
    if err := s.repo.Save(ctx, order); err != nil {
        return nil, fmt.Errorf("saving order: %w", err)
    }
    
    return order, nil
}

// ConfirmOrder transitions an order to confirmed status.
//
// Only pending orders can be confirmed.
// Returns ErrInvalidStatus if the order is not in pending status.
func (s *Service) ConfirmOrder(ctx context.Context, orderID string) error {
    order, err := s.repo.Get(ctx, orderID)
    if err != nil {
        return fmt.Errorf("getting order: %w", err)
    }
    
    if order.Status != StatusPending {
        return ErrInvalidStatus
    }
    
    order.Status = StatusConfirmed
    
    if err := s.repo.Save(ctx, order); err != nil {
        return fmt.Errorf("updating order: %w", err)
    }
    
    return nil
}

func generateID() string {
    // Implementation
    return "ORD-123"
}
```

---

## 🎯 Quick Reference Card

### Clean Code Checklist
- [ ] Functions < 50 lines
- [ ] Max 3-4 function parameters
- [ ] No deep nesting (return early)
- [ ] Errors always handled
- [ ] Meaningful variable names
- [ ] Code grouped logically
- [ ] Consistent formatting (go fmt)
- [ ] No global mutable state
- [ ] Interfaces for testability
- [ ] Tests for all public functions

### Documentation Checklist
- [ ] Every exported symbol documented
- [ ] Package has overview comment
- [ ] Complex functions have examples
- [ ] Error cases documented
- [ ] README.md exists
- [ ] godoc generates correctly
- [ ] No spelling mistakes
- [ ] Comments are sentences with periods

---

## 🚀 Next Steps

1. **Read** the [Effective Go](https://go.dev/doc/effective_go) guide
2. **Use** [golangci-lint](https://golangci-lint.run/) in your projects
3. **Practice** writing tests first (TDD)
4. **Review** your old code and refactor using these rules
5. **Contribute** to open source to see clean code examples

---

**Remember:** Clean code is a practice, not a destination. Keep refactoring, keep learning! 🎯
