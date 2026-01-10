# 🧠 Logic and Thinking Building: 20 Rules for Professional Engineering

To move from a student to a professional Go developer, follow these rules in every project you build:

### 1. Understand Before You Code
Never write a single line of code until you can explain the logic in plain English (or Urdu). If you can't explain it simply, you don't understand it well enough to code it.

### 2. Expect the "Worst Case" (Edge Cases)
Don't just code for when things go right. Ask: "What if the user enters a negative number?", "What if the file is missing?", "What if the input is empty?" A pro codes for the 1% of things that go wrong.

### 3. Errors are Not Suggestions
In Go, errors are values. Never ignore them. Every time a function returns an `err`, you must check it (`if err != nil`) and handle it gracefully.

### 4. Small and Focused Functions
A function should do **one thing** and do it well. If your function is 100 lines long, it's doing too much. Break it down into smaller, reusable pieces.

### 5. Name with Purpose
Avoid names like `x`, `data`, or `temp`. Use descriptive names like `accountBalance`, `isTransactionValid`, or `userEmail`. Code should read like a story.

### 6. Don't Repeat Yourself (DRY)
If you find yourself copy-pasting code, stop. Create a function or a struct instead. Duplicate code is where bugs hide.

### 7. Validate All Input
Never trust the user. Users will type letters where numbers should be, or leave fields blank. Always check and clean the data before your program uses it.

### 8. Separation of Concerns
Keep your "Business Logic" (math, data processing) separate from your "UI" (printing to the terminal). This makes it easy to change your UI later without breaking the math.

### 9. Plan for Growth (Scalability)
Don't just code for today. Ask: "If I have 10,000 users tomorrow, will this structure still work?" This is why we use Structs and Interfaces.

### 10. Write Code for Humans
You don't write code for the computer; the computer will understand anything. You write code for your future self and your teammates. Make it clean, organized, and easy to read.

---

# 🚀 Part 2: Rules for Clean and Efficient Code

To reach the next level, apply these advanced principles to your daily coding:

### 11. YAGNI (You Ain't Gonna Need It)
Don't write code for features you *think* you might need in the future. Only write what you need **now**. Over-engineering makes code harder to read and maintain.

### 12. KISS (Keep It Simple, Stupid)
If there are two ways to solve a problem, choose the simpler one. Complexity is the enemy of reliability. A simple solution is easier to test and harder to break.

### 13. Comments explain 'Why', Code explains 'How'
If your code is clean, you shouldn't need a comment to explain *what* it is doing. Use comments to explain the **reasoning** behind a complex decision.

### 14. The Boy Scout Rule
"Always leave the campground cleaner than you found it." If you see a small mess in a file (even if you didn't make it), fix it while you are there. This prevents "code rot."

### 15. Avoid "Magic Numbers"
Never use raw numbers like `86400` in your code. Use a constant like `const SecondsInADay = 86400`. This makes the code self-documenting.

### 16. Fail Fast (Guard Clauses)
Check for errors and invalid conditions at the very beginning of your function. If something is wrong, `return` immediately. This keeps the "happy path" of your code from being buried inside deep `if` statements.

### 17. Consistent Formatting (gofmt)
In Go, there is only one way to format code. Always run `go fmt`. Consistent code is easier for the community and your future team to read.

### 18. Composition over Inheritance
Go doesn't have traditional classes. Use **Structs** and **Interfaces** to build complex behavior. This makes your code more flexible and easier to test.

### 19. Premature Optimization is the Root of All Evil
Don't worry about making your code "super fast" until it is working and correct. First make it work, then make it right, then (only if needed) make it fast.

### 20. Think in Data, not just Syntax
A professional engineer understands how data flows through the system. Before writing a function, think: "What is the input? How does it change? What is the output?"

---
**"A junior developer writes code that a computer can understand. A senior developer writes code that a human can understand."**

---
> **Moeed ul Hassan**  
> *Written on 10 January 2026*  
> **-- Programmer of Go --**
