# URL Shortener Project

## 📌 How It Works
The goal of this project is to build a service that takes a long URL and converts it into a short, manageable link. When someone visits the short link, they are automatically redirected to the original long URL.

### The Flow
1.  **Shortening (POST /shorten)**
    *   **Input**: User sends a JSON request with a long URL (e.g., `https://www.google.com`).
    *   **Process**:
        *   The server receives the request.
        *   It generates a random 6-character "Short Code" (e.g., `aX9d21`).
        *   It saves the pair `aX9d21` -> `https://www.google.com` in memory (a map).
    *   **Output**: The server returns the short URL (e.g., `http://localhost:8080/aX9d21`).

2.  **Redirecting (GET /{shortCode})**
    *   **Input**: User visits `http://localhost:8080/aX9d21` in their browser.
    *   **Process**:
        *   The server extracts the code `aX9d21` from the URL.
        *   It looks up `aX9d21` in its memory map.
        *   If found, it gets the original URL.
    *   **Output**: The server sends a "302 Found" redirect response, sending the user's browser to `https://www.google.com`.

---

## 📂 Project Structure
To keep the code clean and organized, we will split it into three files, all part of `package main`:

1.  **`model.go`** (The Data)
    *   Defines the structure of the data we expect from the user (`ShortenRequest`) and the data we send back (`ShortenResponse`).
2.  **`handler.go`** (The Logic)
    *   Contains the functions that do the actual work: `shortenHandler` (create short link) and `redirectHandler` (handle clicks).
    *   Manages the storage (the map).
3.  **`main.go`** (The Entry Point)
    *   Sets up the web server.
    *   Tells the server which function to call for which URL path.
    *   Starts listening for requests.

---

## 🚀 Implementation Instructions

### Step 1: Define the Data (`model.go`)
Create the structs to handle JSON data.
*   `ShortenRequest`: To capture the `{"url": "..."}` from the user.
*   `ShortenResponse`: To send back `{"short_url": "..."}`.

### Step 2: Implement the Logic (`handler.go`)
*   Create a global map `urlStore` to save the URLs.
*   Create a `generateShortCode` function to make random strings.
*   **`shortenHandler`**:
    1.  Parse the incoming JSON.
    2.  Generate a code.
    3.  Save it to the map.
    4.  Return the new short URL.
*   **`redirectHandler`**:
    1.  Get the code from the URL path.
    2.  Find the original URL in the map.
    3.  Redirect the user.

### Step 3: Wire it Up (`main.go`)
*   Use `http.HandleFunc` to connect `/shorten` to `shortenHandler`.
*   Use `http.HandleFunc` to connect `/` to `redirectHandler`.
*   Start the server with `http.ListenAndServe`.

---

## 🧠 Concepts Practiced
Building this project will help you master the following Go concepts:

1.  **Structs & JSON Tags** (`model.go`)
    *   Defining data shapes.
    *   Using tags like `` `json:"url"` `` to map Go fields to JSON keys.
2.  **HTTP Handlers & Routing** (`main.go`, `handler.go`)
    *   Using `http.HandleFunc` to map URLs to functions.
    *   Writing handler functions with `w http.ResponseWriter` and `r *http.Request`.
3.  **JSON Encoding/Decoding** (`handler.go`)
    *   Reading JSON from a request body (`json.NewDecoder`).
    *   Sending JSON back to the client (`json.NewEncoder`).
4.  **Maps & State Management** (`handler.go`)
    *   Using a `map` to store data in memory.
    *   Understanding that this data is lost when the server restarts (persistence comes later!).
5.  **Concurrency Safety (Mutex)** (`handler.go`)
    *   **Crucial Concept**: Web servers handle multiple requests at once.
    *   Using `sync.Mutex` (or `sync.RWMutex`) to prevent "race conditions" when writing to the map.
6.  **Randomness & Slices** (`handler.go`)
    *   Generating random strings using `math/rand` and byte slices.

