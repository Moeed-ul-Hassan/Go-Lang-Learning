# Gen-Code: Professional Scaffolding Utility

**Gen-Code** is a powerful, interactive CLI tool built in Go that automates the creation of project boilerplates. It moves beyond simple file copies to provide a structured, wizard-like experience for developers to bootstrap their next masterpiece in seconds.

## 🚀 Features

- **Interactive TUI**: A polished wizard-driven interface built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).
- **Multi-Language Support**: Scaffolds projects for **Go**, **JavaScript**, and **Python**.
- **Framework Selectors**: Dynamic selection of popular frameworks:
    - **Go**: Gin, Echo, Fiber
    - **JavaScript**: Express, Fastify
    - **Python**: Flask, FastAPI, Django
- **Customizable Complexity**: Choose between Minimal (MVP), Standard (Clean Architecture), or Enterprise levels.
- **Realistic Boilerplate**: Generates authentic code, dependency files (`go.mod`, `package.json`, `requirements.txt`), and directory structures.
- **Attribution**: Every generated project features a custom signature and developer credit for **Moeed ul Hassan**.

## 🛠 Tech Stack

- **Language**: Go
- **UI Framework**: Bubble Tea / Lip Gloss / Bubbles
- **Engine**: Custom matrix-based scaffolding engine

## 📖 Usage

1.  **Clone the repository**.
2.  **Build the application**:
    ```bash
    go build ./cmd/gen-code
    ```
3.  **Run the tool**:
    ```bash
    ./gen-code
    ```
4.  **Follow the steps**:
    - Enter your App Name.
    - Select your Language.
    - Select your Framework.
    - Select the Environment type.
    - Select the Complexity level.
    - Enter the Output Path (e.g., `./my-new-app`).

## 🏗 Architecture

The project follows a modular architecture designed for scalability and separation of concerns:

-   **`cmd/gen-code/`**: The **Entry Point**. It initializes the Bubble Tea program and handles the top-level execution loop.
-   **`internal/tui/`**: The **User Interface Layer**. Built using the **The Elm Architecture (TEA)**, it manages state transitions (Model), user input handling (Update), and terminal rendering (View).
-   **`internal/matrix/`**: The **Logic & Template Layer**. It acts as a repository of project definitions. It contains the data structures and templates for every supported language and framework.
-   **`internal/scaffold/`**: The **Execution Engine**. This layer interacts with the OS file system to create directories and write files based on the selection from the Matrix.

## 🧠 How it was Made

Building **Gen-Code** involved a step-by-step evolution from a basic prototype to a functional utility:

1.  **Requirement Analysis**: Identified the need for a tool that eliminates the "blank page" problem for developers starting new projects.
2.  **Environment Cleanup**: Removed unnecessary web-scaffolding legacy files to ensure a lean, CLI-focused codebase.
3.  **TUI Design**: Chose **Bubble Tea** for its robust state management. I implemented a state machine to guide the user through a logical 7-step wizard (App Name -> Language -> Framework -> Env -> Complexity -> Path).
4.  **Matrix Implementation**: Developed a structured way to store and retrieve boilerplate code. This involved writing realistic, running code examples for Go, JS, and Python stacks.
5.  **Scaffolding Logic**: Built a robust engine using Go's `os` and `path/filepath` packages to handle cross-platform file creation and directory nesting.
6.  **Branding & Signature**: Integrated a custom header system to ensure every generated file is signed with the project identity and developer credit.

## 📁 Project Structure

- `cmd/gen-code/`: Application entry point.
- `internal/tui/`: Bubble Tea model and view logic.
- `internal/matrix/`: Scaffolding templates and definitions.
- `internal/scaffold/`: Scaffolding execution engine.

---
Created by **Moeed ul Hassan**
