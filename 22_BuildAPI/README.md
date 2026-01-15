# BuildAPI - Go REST API Practice

This project demonstrates how to build a robust RESTful API in Go using the `gorilla/mux` router. It covers essential concepts like routing, JSON handling, and CRUD (Create, Read, Update, Delete) operations using a slice as a mock database.

## 🚀 Features

*   **RESTful Architecture**: Implements standard HTTP methods (GET, POST, PUT, DELETE).
*   **Gorilla Mux**: Uses the popular `gorilla/mux` package for efficient routing.
*   **JSON Handling**: Demonstrates encoding and decoding of JSON data.
*   **Mock Database**: Uses a Go slice (`[]Course`) to simulate a database for storing course data.
*   **Seeding**: Includes initial data seeding for testing.
*   **Random ID Generation**: Generates unique IDs for new courses.

## 🛠️ Tech Stack

*   **Language**: Go (Golang)
*   **Router**: [github.com/gorilla/mux](https://github.com/gorilla/mux)
*   **Standard Libraries**: `encoding/json`, `fmt`, `log`, `math/rand`, `net/http`, `strconv`, `time`

## 📂 Project Structure

*   `main.go`: Contains the entire application logic, including models, middleware, controllers, and routing.
*   `go.mod`: Go module definition and dependencies.

## 🔌 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | Home route (Welcome message) |
| `GET` | `/courses` | Get all courses |
| `GET` | `/course/{id}` | Get a single course by ID |
| `POST` | `/course` | Create a new course |
| `PUT` | `/course/{id}` | Update an existing course |
| `DELETE` | `/course/{id}` | Delete a course |

## 🏃‍♂️ How to Run

1.  **Clone/Navigate** to the directory:
    ```bash
    cd "22_BuildAPI"
    ```

2.  **Install Dependencies**:
    ```bash
    go get -u github.com/gorilla/mux
    ```

3.  **Run the Application**:
    ```bash
    go run main.go
    ```

4.  **Test the API**:
    *   The server will start at `http://localhost:4000`.
    *   Use a tool like **Postman**, **Thunder Client**, or **curl** to send requests to the endpoints.

## 📝 Data Model

**Course Struct:**
```go
type Course struct {
    CourseId    string  `json:"courseid"`
    CourseName  string  `json:"coursename"`
    CoursePrice int     `json:"price"`
    Author      *Author `json:"author"`
}
```

**Author Struct:**
```go
type Author struct {
    Fullname string `json:"fullname"`
    Website  string `json:"website"`
}
```

## 🧠 Key Concepts Learned

*   **Routing**: Mapping URLs to specific handler functions.
*   **Middleware**: Basic validation (e.g., checking if a course is empty).
*   **JSON Encoding/Decoding**: Using `json.NewEncoder` and `json.NewDecoder`.
*   **Slices**: Manipulating slices to add, update, and remove items.
*   **Pointers**: Using pointers for nested structs (`*Author`).
