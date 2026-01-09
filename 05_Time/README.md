# The Time Package in Go

The `time` package is essential for handling dates, times, durations, and tickers in Go.

## 1. Getting the Current Time
```go
now := time.Now()
fmt.Println(now)
```

## 2. Formatting and Parsing (The "Magic Date")
Go does **not** use standard format codes like `%Y-%m-%d`. Instead, it uses a specific reference time:
**Mon Jan 2 15:04:05 MST 2006** (Memory trick: 1 2 3 4 5 6 7).

### Formatting
```go
fmt.Println(now.Format("01-02-2006 Monday")) // MM-DD-YYYY Day
fmt.Println(now.Format("15:04:05"))          // 24-hour time
```

### Parsing
```go
layout := "2006-01-02"
str := "2024-05-20"
t, _ := time.Parse(layout, str)
```

## 3. Durations
Durations represent the elapsed time between two instants.
```go
fiveSeconds := 5 * time.Second
tenMinutes := 10 * time.Minute
```

## 4. Time Arithmetic
```go
future := now.Add(24 * time.Hour)
past := now.Add(-1 * time.Hour)
diff := future.Sub(now) // Returns a time.Duration
```

## 5. Sleep and Tickers
```go
time.Sleep(2 * time.Second) // Pause execution

ticker := time.NewTicker(1 * time.Second)
for t := range ticker.C {
    fmt.Println("Tick at", t)
}
```

## Common Pitfalls
- **Time Zones**: `time.Now()` uses the local time zone. Use `time.Now().UTC()` for consistent server-side time.
- **Comparing Times**: Use `t1.Equal(t2)` instead of `t1 == t2` to account for monotonic clock readings.

## Summary
The `time` package is powerful but requires getting used to the unique formatting system. Once mastered, it provides robust tools for any time-related task.
