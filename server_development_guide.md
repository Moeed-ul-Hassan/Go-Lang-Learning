# High-Performance Server Development Guide

> A comprehensive guide to building production-ready servers capable of handling 1000+ requests per second

---

## Table of Contents

1. [Steps to Create a Server](#steps-to-create-a-server)
2. [20 Rules for High-Performance Servers](#20-rules-for-high-performance-servers)
3. [Key Points to Keep in Mind](#key-points-to-keep-in-mind)
4. [Performance Testing](#performance-testing)
5. [Additional Resources](#additional-resources)

---

## Steps to Create a Server

### 1. Planning Phase

**Define Requirements**
- Identify API endpoints and functionality
- Choose architecture pattern (REST, GraphQL, gRPC, WebSocket)
- Design database schema and relationships
- Plan authentication and authorization strategy
- Define SLAs and performance requirements

**Technology Stack Selection**
- Programming language (Go recommended for high performance)
- Web framework (Gin, Echo, Chi, or standard library)
- Database (PostgreSQL, MySQL, MongoDB)
- Caching layer (Redis, Memcached)
- Message queue (RabbitMQ, Kafka, NATS)

### 2. Setup Phase

**Project Initialization**
```bash
# Initialize Go module
go mod init github.com/username/project-name

# Create project structure
mkdir -p cmd/api internal/{handlers,models,middleware,database,config}
```

**Directory Structure**
```
project/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── middleware/
│   ├── database/
│   └── config/
├── pkg/
├── migrations/
├── tests/
├── docker/
├── .env.example
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

**Environment Configuration**
- Create `.env` file for local development
- Set up configuration management
- Separate configs for dev, staging, production

### 3. Development Phase

**Core Implementation**
1. Set up router and routes
2. Implement middleware stack
   - Logger
   - Recovery (panic handler)
   - CORS
   - Authentication/Authorization
   - Rate limiting
   - Request ID tracking
3. Create handlers and business logic
4. Implement database layer
5. Add validation
6. Implement error handling

**Example Server Setup**
```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // Create server
    server := &http.Server{
        Addr:         ":8080",
        Handler:      setupRouter(),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    // Start server in goroutine
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }
}
```

### 4. Testing & Optimization

**Testing Strategy**
- Unit tests for business logic
- Integration tests for API endpoints
- Load testing for performance validation
- Security testing (OWASP Top 10)

**Performance Optimization**
- Profile with `pprof`
- Optimize database queries
- Implement caching strategy
- Reduce memory allocations

### 5. Deployment

**Containerization**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

**Production Checklist**
- Set up CI/CD pipeline
- Configure monitoring and alerting
- Implement log aggregation
- Set up backup and disaster recovery
- Configure SSL/TLS certificates
- Implement health checks

---

## 20 Rules for High-Performance Servers

### Performance & Scalability

#### 1. **Use Connection Pooling**

> Never create new connections per request

- **Database**: Limit pool size (10-100 connections based on load)
- **HTTP Clients**: Reuse connections with proper timeout settings
- **Redis/Cache**: Maintain persistent connection pools

```go
// Good: Connection pool
db, err := sql.Open("postgres", dsn)
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)

// Bad: New connection per request
// db, err := sql.Open("postgres", dsn) // DON'T DO THIS IN HANDLER
```

#### 2. **Implement Proper Timeouts**

> Prevent resource exhaustion from hanging connections

```go
server := &http.Server{
    Addr:              ":8080",
    Handler:           router,
    ReadTimeout:       5 * time.Second,   // Time to read request
    WriteTimeout:      10 * time.Second,  // Time to write response
    IdleTimeout:       120 * time.Second, // Keep-alive timeout
    ReadHeaderTimeout: 2 * time.Second,   // Time to read headers
}

// Also set timeouts for HTTP clients
client := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

#### 3. **Use Goroutines Wisely**

> Unbounded concurrency leads to resource exhaustion

```go
// Good: Worker pool pattern
type WorkerPool struct {
    tasks chan Task
    wg    sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
    wp := &WorkerPool{
        tasks: make(chan Task, 100),
    }
    
    for i := 0; i < numWorkers; i++ {
        wp.wg.Add(1)
        go wp.worker()
    }
    
    return wp
}

// Bad: Unlimited goroutines
// for _, item := range millions {
//     go process(item) // DON'T DO THIS
// }
```

#### 4. **Leverage Caching**

> Reduce database load and improve response times

**Caching Layers:**
- **L1**: In-memory cache (map with sync.RWMutex or sync.Map)
- **L2**: Redis/Memcached for distributed caching
- **L3**: CDN for static assets

**Caching Strategies:**
- Cache-aside (lazy loading)
- Write-through
- Write-behind
- TTL-based expiration

```go
// Example: Simple in-memory cache
type Cache struct {
    mu    sync.RWMutex
    items map[string]CacheItem
}

type CacheItem struct {
    Value      interface{}
    Expiration time.Time
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, found := c.items[key]
    if !found || time.Now().After(item.Expiration) {
        return nil, false
    }
    return item.Value, true
}
```

#### 5. **Minimize Memory Allocations**

> Reduce garbage collection pressure

```go
// Use sync.Pool for frequently allocated objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func handler(w http.ResponseWriter, r *http.Request) {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()
    
    // Use buffer
}

// Preallocate slices when size is known
items := make([]Item, 0, expectedSize)

// Use strings.Builder for string concatenation
var builder strings.Builder
builder.Grow(estimatedSize)
for _, s := range parts {
    builder.WriteString(s)
}
result := builder.String()
```

---

### Concurrency & Thread Safety

#### 6. **Handle Concurrent Requests Safely**

> Prevent race conditions and data corruption

```go
// Good: Protected shared state
type SafeCounter struct {
    mu    sync.RWMutex
    count map[string]int
}

func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count[key]++
}

func (c *SafeCounter) Get(key string) int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.count[key]
}

// Test for race conditions
// go test -race ./...
```

#### 7. **Use Context for Cancellation**

> Propagate cancellation and deadlines through request chain

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    // Pass context to all downstream calls
    data, err := fetchData(ctx)
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            http.Error(w, "Request timeout", http.StatusGatewayTimeout)
            return
        }
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(data)
}

func fetchData(ctx context.Context) (*Data, error) {
    // Check context before expensive operations
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
    }
    
    // Perform operation with context
    return db.QueryContext(ctx, "SELECT ...")
}
```

#### 8. **Implement Graceful Shutdown**

> Allow in-flight requests to complete before shutdown

```go
func main() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
    
    // Channel to listen for shutdown signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    
    // Start server
    go func() {
        log.Println("Server starting on :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server error: %v", err)
        }
    }()
    
    // Wait for interrupt signal
    <-quit
    log.Println("Shutting down server...")
    
    // Give ongoing requests 30 seconds to complete
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }
    
    log.Println("Server exited gracefully")
}
```

---

### Database & I/O

#### 9. **Batch Database Operations**

> Reduce network round trips and improve throughput

```go
// Good: Batch insert
func BatchInsert(users []User) error {
    valueStrings := make([]string, 0, len(users))
    valueArgs := make([]interface{}, 0, len(users)*3)
    
    for i, user := range users {
        valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", 
            i*3+1, i*3+2, i*3+3))
        valueArgs = append(valueArgs, user.Name, user.Email, user.Age)
    }
    
    query := fmt.Sprintf("INSERT INTO users (name, email, age) VALUES %s",
        strings.Join(valueStrings, ","))
    
    _, err := db.Exec(query, valueArgs...)
    return err
}

// Bad: Individual inserts
// for _, user := range users {
//     db.Exec("INSERT INTO users...", user) // DON'T DO THIS
// }
```

#### 10. **Index Database Properly**

> Critical for query performance at scale

**Indexing Best Practices:**
- Index columns used in `WHERE`, `JOIN`, `ORDER BY`, `GROUP BY`
- Composite indexes for multi-column queries
- Avoid over-indexing (slows down writes)
- Monitor index usage

```sql
-- Create indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_orders_user_date ON orders(user_id, created_at DESC);

-- Analyze query performance
EXPLAIN ANALYZE SELECT * FROM users WHERE email = 'test@example.com';

-- Monitor slow queries (PostgreSQL)
SELECT query, calls, total_time, mean_time
FROM pg_stat_statements
ORDER BY total_time DESC
LIMIT 10;
```

#### 11. **Use Async I/O Where Possible**

> Don't block on external operations

```go
// Good: Async with channels
func processAsync(items []Item) <-chan Result {
    results := make(chan Result, len(items))
    
    go func() {
        defer close(results)
        for _, item := range items {
            results <- process(item)
        }
    }()
    
    return results
}

// Good: Non-blocking external API calls
func fetchMultiple(urls []string) []Response {
    results := make([]Response, len(urls))
    var wg sync.WaitGroup
    
    for i, url := range urls {
        wg.Add(1)
        go func(i int, url string) {
            defer wg.Done()
            results[i] = fetch(url)
        }(i, url)
    }
    
    wg.Wait()
    return results
}

// For heavy tasks: Use message queue
func enqueueTask(task Task) error {
    payload, _ := json.Marshal(task)
    return messageQueue.Publish("tasks", payload)
}
```

---

### Security

#### 12. **Validate All Input**

> Never trust user input - it's the #1 security principle

```go
type CreateUserRequest struct {
    Username string `json:"username" validate:"required,alphanum,min=3,max=20"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"required,gte=18,lte=120"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    // Limit request body size
    r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB
    
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // Validate
    validate := validator.New()
    if err := validate.Struct(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // Sanitize - use allowlists
    username := sanitizeUsername(req.Username)
    
    // Process request...
}
```

#### 13. **Implement Rate Limiting**

> Protect against abuse and DDoS attacks

```go
// Token bucket rate limiter
type RateLimiter struct {
    mu      sync.Mutex
    limiter map[string]*rate.Limiter
    rate    rate.Limit
    burst   int
}

func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
    return &RateLimiter{
        limiter: make(map[string]*rate.Limiter),
        rate:    r,
        burst:   b,
    }
}

func (rl *RateLimiter) getLimiter(ip string) *rate.Limiter {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    limiter, exists := rl.limiter[ip]
    if !exists {
        limiter = rate.NewLimiter(rl.rate, rl.burst)
        rl.limiter[ip] = limiter
    }
    
    return limiter
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := getIP(r)
        limiter := rl.getLimiter(ip)
        
        if !limiter.Allow() {
            http.Error(w, "Too many requests", http.StatusTooManyRequests)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

// Usage: 100 requests per second per IP, burst of 50
limiter := NewRateLimiter(100, 50)
router.Use(limiter.Middleware)
```

#### 14. **Use HTTPS Only**

> Encrypt all traffic - no exceptions

```go
// Force HTTPS redirect
func redirectToHTTPS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.TLS == nil {
            url := "https://" + r.Host + r.RequestURI
            http.Redirect(w, r, url, http.StatusMovedPermanently)
            return
        }
        
        // Security headers
        w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        
        next.ServeHTTP(w, r)
    })
}

// TLS configuration
tlsConfig := &tls.Config{
    MinVersion:               tls.VersionTLS12,
    PreferServerCipherSuites: true,
    CurvePreferences: []tls.CurveID{
        tls.CurveP256,
        tls.X25519,
    },
}

server := &http.Server{
    Addr:      ":443",
    Handler:   router,
    TLSConfig: tlsConfig,
}

server.ListenAndServeTLS("cert.pem", "key.pem")
```

#### 15. **Protect Against Common Attacks**

> Defense in depth approach

**SQL Injection Prevention:**
```go
// Good: Parameterized queries
rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)

// Bad: String concatenation
// query := "SELECT * FROM users WHERE email = '" + email + "'" // NEVER DO THIS
```

**XSS Prevention:**
```go
import "html/template"

// HTML templates automatically escape
tmpl := template.Must(template.ParseFiles("page.html"))
tmpl.Execute(w, data)

// For JSON APIs: Set proper content type
w.Header().Set("Content-Type", "application/json")
w.Header().Set("X-Content-Type-Options", "nosniff")
```

**CSRF Prevention:**
```go
// Use CSRF tokens
import "github.com/gorilla/csrf"

csrfMiddleware := csrf.Protect(
    []byte("32-byte-long-auth-key"),
    csrf.Secure(true), // Require HTTPS
)

router := mux.NewRouter()
router.Use(csrfMiddleware)
```

**CORS Configuration:**
```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Don't use wildcard (*) in production
        w.Header().Set("Access-Control-Allow-Origin", "https://yourdomain.com")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Max-Age", "3600")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

---

### Monitoring & Reliability

#### 16. **Implement Comprehensive Logging**

> Logs are your production debugging tool

```go
import "go.uber.org/zap"

// Structured logging setup
logger, _ := zap.NewProduction()
defer logger.Sync()

// Request logging middleware
func loggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            
            // Generate request ID
            requestID := uuid.New().String()
            ctx := context.WithValue(r.Context(), "requestID", requestID)
            
            // Wrap response writer to capture status code
            wrapped := wrapResponseWriter(w)
            
            next.ServeHTTP(wrapped, r.WithContext(ctx))
            
            logger.Info("request completed",
                zap.String("request_id", requestID),
                zap.String("method", r.Method),
                zap.String("path", r.URL.Path),
                zap.Int("status", wrapped.status),
                zap.Duration("duration", time.Since(start)),
                zap.String("ip", getIP(r)),
            )
        })
    }
}

// Log levels
logger.Debug("Debug info", zap.String("key", "value"))
logger.Info("Request processed")
logger.Warn("Deprecated API called")
logger.Error("Database error", zap.Error(err))
logger.Fatal("Critical failure") // Exits with os.Exit(1)
```

**Logging Best Practices:**
- Use structured logging (JSON format)
- Include request IDs for tracing
- Log at appropriate levels
- Don't log sensitive data (passwords, tokens, PII)
- Use log sampling for high-volume endpoints

#### 17. **Add Health Checks**

> Enable load balancers and orchestrators to manage your service

```go
// Liveness probe - "Is the service alive?"
func livenessHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

// Readiness probe - "Is the service ready to handle traffic?"
func readinessHandler(db *sql.DB, cache *redis.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Check database
        ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
        defer cancel()
        
        if err := db.PingContext(ctx); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            json.NewEncoder(w).Encode(map[string]string{
                "status": "unavailable",
                "reason": "database connection failed",
            })
            return
        }
        
        // Check Redis
        if err := cache.Ping(ctx).Err(); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            json.NewEncoder(w).Encode(map[string]string{
                "status": "unavailable",
                "reason": "cache connection failed",
            })
            return
        }
        
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{
            "status": "ready",
        })
    }
}

// Metrics endpoint (Prometheus format)
func metricsHandler(w http.ResponseWriter, r *http.Request) {
    promhttp.Handler().ServeHTTP(w, r)
}

// Register endpoints
router.HandleFunc("/health", livenessHandler)
router.HandleFunc("/ready", readinessHandler(db, cache))
router.HandleFunc("/metrics", metricsHandler)
```

#### 18. **Monitor Key Metrics**

> You can't improve what you don't measure

**Critical Metrics:**

| Metric | Description | Target |
|--------|-------------|--------|
| **Latency P50** | 50th percentile response time | < 100ms |
| **Latency P95** | 95th percentile response time | < 500ms |
| **Latency P99** | 99th percentile response time | < 1000ms |
| **Error Rate** | Percentage of failed requests | < 0.1% |
| **Throughput** | Requests per second | 1000+ req/s |
| **CPU Usage** | CPU utilization | < 70% |
| **Memory Usage** | RAM utilization | < 80% |
| **Connection Pool** | Active DB connections | Within limits |
| **Goroutines** | Number of goroutines | Stable over time |

```go
import "github.com/prometheus/client_golang/prometheus"

// Define metrics
var (
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request latencies in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "endpoint", "status"},
    )
    
    requestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )
    
    activeConnections = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "active_connections",
            Help: "Number of active connections",
        },
    )
)

func init() {
    prometheus.MustRegister(requestDuration)
    prometheus.MustRegister(requestCounter)
    prometheus.MustRegister(activeConnections)
}

// Metrics middleware
func metricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        wrapped := wrapResponseWriter(w)
        
        activeConnections.Inc()
        defer activeConnections.Dec()
        
        next.ServeHTTP(wrapped, r)
        
        duration := time.Since(start).Seconds()
        status := strconv.Itoa(wrapped.status)
        
        requestDuration.WithLabelValues(r.Method, r.URL.Path, status).Observe(duration)
        requestCounter.WithLabelValues(r.Method, r.URL.Path, status).Inc()
    })
}
```

---

### Code Quality

#### 19. **Handle Errors Properly**

> Good error handling prevents cascading failures

```go
// Define custom error types
type AppError struct {
    Code    string
    Message string
    Err     error
}

func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

// Wrap errors for context
func fetchUser(id int) (*User, error) {
    user, err := db.GetUser(id)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch user %d: %w", id, err)
    }
    return user, nil
}

// Error handling in handlers
func handler(w http.ResponseWriter, r *http.Request) {
    user, err := fetchUser(123)
    if err != nil {
        // Log full error with stack trace
        logger.Error("failed to process request",
            zap.Error(err),
            zap.String("request_id", getRequestID(r)),
        )
        
        // Return sanitized error to client
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(user)
}

// Never ignore errors
// _, err := someOperation()
// if err != nil {
//     // Handle it!
// }

// Use errors.Is and errors.As for error checking
if errors.Is(err, sql.ErrNoRows) {
    // Handle not found
}

var appErr *AppError
if errors.As(err, &appErr) {
    // Handle custom error type
}
```

#### 20. **Design for Horizontal Scaling**

> Build stateless services that can scale out

**Stateless Design Principles:**

1. **No In-Memory Session Storage**
   ```go
   // Bad: In-memory session
   // var sessions = make(map[string]*Session)
   
   // Good: External session storage
   func getSession(sessionID string) (*Session, error) {
       var session Session
       data, err := redisClient.Get(ctx, "session:"+sessionID).Result()
       if err != nil {
           return nil, err
       }
       json.Unmarshal([]byte(data), &session)
       return &session, nil
   }
   ```

2. **Externalize Configuration**
   ```go
   // Load from environment or config service
   config := Config{
       DatabaseURL: os.Getenv("DATABASE_URL"),
       RedisURL:    os.Getenv("REDIS_URL"),
       APIKey:      os.Getenv("API_KEY"),
   }
   ```

3. **Use Sticky Sessions Only When Necessary**
   - Prefer stateless authentication (JWT)
   - Store state in distributed cache
   - Use database for persistent state

4. **Idempotent Operations**
   ```go
   // Use idempotency keys for critical operations
   func processPayment(idempotencyKey string, amount int) error {
       // Check if already processed
       exists, err := redis.Exists(ctx, "payment:"+idempotencyKey).Result()
       if err != nil {
           return err
       }
       if exists > 0 {
           return nil // Already processed
       }
       
       // Process payment
       err = chargePayment(amount)
       if err != nil {
           return err
       }
       
       // Store idempotency key
       redis.Set(ctx, "payment:"+idempotencyKey, "processed", 24*time.Hour)
       return nil
   }
   ```

5. **Load Balancer Configuration**
   - Use health checks
   - Configure session affinity if needed
   - Set proper timeout values
   - Enable connection draining for graceful shutdown

---

## Key Points to Keep in Mind

### Architecture Decisions

#### Stateless vs Stateful
- **Stateless** (preferred for scaling)
  - Session data in Redis/database
  - JWT for authentication
  - Can add/remove instances freely
  
- **Stateful** (avoid when possible)
  - Requires sticky sessions
  - Complex deployment coordination
  - Harder to scale

#### Monolith vs Microservices
- **Start with Monolith**
  - Easier to develop and deploy
  - Better performance (no network calls)
  - Simpler debugging and testing
  
- **Split to Microservices When:**
  - Team size grows (>8-10 developers)
  - Different scaling requirements per component
  - Need independent deployment cycles
  - Clear bounded contexts emerge

#### API Design
- **RESTful Principles**
  - Use proper HTTP methods (GET, POST, PUT, DELETE)
  - Meaningful resource URLs
  - Proper status codes
  
- **Versioning Strategy**
  - URL versioning: `/api/v1/users`
  - Header versioning: `Accept: application/vnd.api.v1+json`
  - Maintain backward compatibility

- **Pagination**
  ```go
  // Cursor-based pagination (recommended)
  GET /api/v1/users?cursor=eyJpZCI6MTAwfQ&limit=50
  
  // Offset-based pagination (simpler but slower)
  GET /api/v1/users?page=2&per_page=50
  ```

### Resource Management

#### Memory Leaks Prevention
```go
// Profile regularly
import _ "net/http/pprof"

go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()

// Access profiles at:
// http://localhost:6060/debug/pprof/heap
// http://localhost:6060/debug/pprof/goroutine
```

**Common Leak Sources:**
- Goroutines that never exit
- Growing maps without cleanup
- Unclosed HTTP response bodies
- Circular references
- Timers not stopped

#### File Descriptors
```bash
# Check limits
ulimit -n

# Set higher limits in production
# /etc/security/limits.conf
* soft nofile 65536
* hard nofile 65536
```

#### Connection Limits
- Database: 10-100 connections per instance
- HTTP client: Reuse with proper idle timeout
- WebSocket: Monitor and limit concurrent connections

### Development Best Practices

#### Code Organization
```
internal/
├── api/
│   ├── handlers/      # HTTP handlers
│   ├── middleware/    # Middleware functions
│   └── routes.go      # Route definitions
├── domain/
│   ├── user/          # User domain logic
│   └── product/       # Product domain logic
├── infrastructure/
│   ├── database/      # Database implementations
│   ├── cache/         # Cache implementations
│   └── queue/         # Message queue
└── config/            # Configuration
```

#### Configuration Management
```go
type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
    AWS      AWSConfig
}

func LoadConfig() (*Config, error) {
    // Load from environment
    viper.AutomaticEnv()
    viper.SetEnvPrefix("APP")
    
    // Load from file
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.ReadInConfig()
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }
    
    return &config, nil
}
```

#### Dependency Management
- Keep dependencies minimal
- Audit dependencies regularly (`go mod verify`)
- Use dependabot or similar for updates
- Pin versions in production

#### Documentation
- OpenAPI/Swagger for API documentation
- Code comments for complex logic
- README with setup instructions
- Architecture decision records (ADRs)

---

## Performance Testing

### Load Testing Tools

#### Apache Bench (ab)
```bash
# Simple load test
ab -n 10000 -c 100 http://localhost:8080/api/users

# With POST data
ab -n 1000 -c 50 -p data.json -T application/json http://localhost:8080/api/users
```

#### wrk
```bash
# Basic test
wrk -t12 -c400 -d30s http://localhost:8080/api/users

# With Lua script
wrk -t12 -c400 -d30s -s post.lua http://localhost:8080/api/users
```

#### k6 (Recommended)
```javascript
// load-test.js
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '2m', target: 100 },  // Ramp up
        { duration: '5m', target: 100 },  // Stay at 100
        { duration: '2m', target: 200 },  // Ramp to 200
        { duration: '5m', target: 200 },  // Stay at 200
        { duration: '2m', target: 0 },    // Ramp down
    ],
    thresholds: {
        http_req_duration: ['p(95)<500'], // 95% under 500ms
        http_req_failed: ['rate<0.01'],   // Less than 1% errors
    },
};

export default function () {
    const res = http.get('http://localhost:8080/api/users');
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response time < 500ms': (r) => r.timings.duration < 500,
    });
    sleep(1);
}
```

```bash
# Run k6 test
k6 run --vus 100 --duration 30s load-test.js
```

### Profiling

#### CPU Profiling
```bash
# Start server with pprof
go run main.go

# Generate CPU profile
curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof

# Analyze
go tool pprof cpu.prof
# (pprof) top
# (pprof) list functionName
# (pprof) web  # Requires graphviz
```

#### Memory Profiling
```bash
# Heap profile
curl http://localhost:6060/debug/pprof/heap > heap.prof
go tool pprof heap.prof

# Allocation profiling
curl http://localhost:6060/debug/pprof/allocs > allocs.prof
go tool pprof allocs.prof
```

#### Goroutine Profiling
```bash
# Check for goroutine leaks
curl http://localhost:6060/debug/pprof/goroutine > goroutine.prof
go tool pprof goroutine.prof
```

### Testing Best Practices

1. **Baseline First**
   - Establish baseline performance
   - Document hardware specs
   - Record initial metrics

2. **Realistic Load**
   - Use production-like data
   - Simulate realistic user behavior
   - Mix of read/write operations
   - Include error scenarios

3. **Incremental Testing**
   - Start with low load
   - Gradually increase
   - Find breaking point
   - Identify bottlenecks

4. **Monitor During Tests**
   - CPU and memory usage
   - Database connections
   - Response times
   - Error rates
   - System metrics (disk I/O, network)

---

## Additional Resources

### Go Resources
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Performance Profiling](https://go.dev/blog/pprof)

### Architecture & Design
- [Twelve-Factor App](https://12factor.net/)
- [API Design Guide](https://cloud.google.com/apis/design)
- [Microservices Patterns](https://microservices.io/patterns/index.html)

### Performance
- [High Performance Browser Networking](https://hpbn.co/)
- [Database Performance Articles](https://use-the-index-luke.com/)

### Security
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Go Security Best Practices](https://github.com/guardrailsio/awesome-golang-security)

### Tools
- **Monitoring**: Prometheus, Grafana, Datadog
- **Logging**: ELK Stack, Loki, Splunk
- **Tracing**: Jaeger, Zipkin
- **Load Testing**: k6, Locust, Gatling
- **API Testing**: Postman, Insomnia, HTTPie

---

## Summary Checklist

Before deploying to production, verify:

- [ ] All timeouts configured (server, client, database)
- [ ] Connection pooling implemented
- [ ] Rate limiting active
- [ ] Input validation on all endpoints
- [ ] HTTPS enforced
- [ ] Security headers set
- [ ] Error handling covers all cases
- [ ] Logging configured with request IDs
- [ ] Health checks respond correctly
- [ ] Metrics exposed (/metrics endpoint)
- [ ] Graceful shutdown implemented
- [ ] Database indexes optimized
- [ ] Load tested at expected traffic + 50%
- [ ] Memory profiled for leaks
- [ ] Monitoring and alerting configured
- [ ] Backup and recovery tested
- [ ] CI/CD pipeline working
- [ ] Documentation up to date
- [ ] Security scan passed
- [ ] Dependencies audited

---

**Remember**: Premature optimization is the root of all evil, but planning for scale from the start is wisdom. Build it right the first time, and you'll thank yourself later.

Good luck building your high-performance server! 🚀
