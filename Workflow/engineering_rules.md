# 📜 The 100 Golden Rules of Engineering Excellence

This document contains 100 hard-earned lessons, best practices, and principles to guide your career from Junior Developer to Principal Engineer.

## 🧠 General Coding Philosophy (1-15)
1.  **Code is for Humans**: Write code to be read by others, not just executed by machines.
2.  **KISS (Keep It Simple, Stupid)**: Complexity is the enemy. Simple code is less buggy and easier to maintain.
3.  **DRY (Don't Repeat Yourself)**: If you copy-paste code, you are creating a future bug. Extract it into a function.
4.  **YAGNI (You Aren't Gonna Need It)**: Don't build features for "someday." Build what you need *now*.
5.  **Boy Scout Rule**: Always leave the code cleaner than you found it.
6.  **Premature Optimization is the Root of All Evil**: Make it work, make it right, *then* make it fast.
7.  **Fail Fast**: Check for errors immediately. Don't let bad data propagate deep into your system.
8.  **Comments Explain "Why", Not "What"**: Code tells you what it does. Comments should explain *why* you did it that way.
9.  **Delete Dead Code**: Commented-out code is rot. Use Git history if you need it back.
10. **Single Responsibility Principle**: A function or class should do one thing and do it well.
11. **Naming is Hard but Critical**: `x` is a bad name. `daysSinceLogin` is a good name. Spend time on naming.
12. **Consistency > Perfection**: A consistently "okay" style is better than a mix of "perfect" and "bad" styles.
13. **Code Reviews are for Learning**: Don't take feedback personally. Your code is not you.
14. **Magic Numbers are Bad**: Don't use `86400` in your code. Use `const SecondsInDay = 86400`.
15. **Learn the Fundamentals**: Frameworks change. Algorithms, data structures, and design patterns stay forever.

## 🐹 Go Specifics (16-35)
16. **Idiomatic Go**: Follow `gofmt` and community standards. Don't write Java in Go.
17. **Errors are Values**: Handle errors explicitly (`if err != nil`). Never ignore them with `_`.
18. **Panic is for Emergencies**: Only panic if the program literally cannot continue (e.g., missing config). Otherwise, return errors.
19. **Interfaces are Small**: Define interfaces where they are *used*, not where they are implemented. Keep them small (1-2 methods).
20. **Accept Interfaces, Return Structs**: This makes your code more flexible and easier to mock.
21. **Channels are for Orchestration, Mutexes are for State**: Use channels to signal data flow, use mutexes to protect data access.
22. **Defer is Your Friend**: Use `defer` to close files, unlock mutexes, and clean up resources right after opening them.
23. **Context is King**: Always pass `context.Context` as the first argument to functions that do I/O or long operations.
24. **Short Variable Declarations**: Use `:=` inside functions, `var` for package-level variables.
25. **Table-Driven Tests**: Use table-driven tests for covering multiple scenarios efficiently.
26. **Avoid `init()`**: Magic initialization is hard to debug. Explicit initialization is better.
27. **Pointer vs Value Receiver**: Use pointers (`*MyStruct`) if you need to modify the struct, otherwise use values.
28. **Empty Interface `interface{}` is a Last Resort**: It bypasses type safety. Use generics or specific interfaces if possible.
29. **Goroutine Leaks**: Never start a goroutine without knowing how it will stop.
30. **Use `go mod tidy`**: Keep your dependencies clean.
31. **Vendor Your Dependencies**: For critical production apps, `go mod vendor` ensures you have a copy of your deps.
32. **Struct Tags**: Use them for JSON, database mapping, and validation.
33. **Make Zero Values Useful**: Design structs so that their default "zero" state is valid and usable.
34. **Don't Export Everything**: Start with private (lowercase) functions. Only export (Capitalize) what is necessary.
35. **Use the Standard Library**: Go's stdlib is powerful. Check `net/http`, `encoding/json`, `time` before grabbing a 3rd party lib.

## 🏗️ Architecture & Design (36-55)
36. **Separation of Concerns**: UI should not know about Database. Database should not know about HTTP.
37. **Dependency Injection**: Pass dependencies (like DB connections) into functions/structs rather than creating them inside.
38. **Repository Pattern**: Abstract database logic behind an interface. This lets you swap SQL for Mongo easily.
39. **Stateless APIs**: REST APIs should be stateless. Store session data in Redis or a Database, not in memory.
40. **Idempotency**: Making the same API request twice should not have side effects (e.g., charging a card twice).
41. **Rate Limiting**: Protect your API from abuse.
42. **Pagination**: Never return "all" records. Always paginate (limit/offset or cursors).
43. **Versioning**: APIs change. Version them (`/v1/users`) to avoid breaking clients.
44. **Health Checks**: Every service needs a `/health` endpoint for monitoring.
45. **Graceful Shutdown**: Handle `SIGTERM` to finish current requests before stopping the server.
46. **Config via Environment**: Store config in Environment Variables, not code. (The 12-Factor App).
47. **Logging Strategy**: Log *events*, not just text. Use structured logging (JSON).
48. **Metrics**: Measure what matters (Latency, Error Rate, Traffic, Saturation).
49. **Caching**: Cache expensive operations, but remember: "There are only two hard things in CS: cache invalidation and naming things."
50. **Event-Driven Architecture**: Decouple services using events (Kafka/RabbitMQ) instead of direct HTTP calls where possible.
51. **Database Normalization**: Start normalized (3NF). Denormalize only for performance reasons.
52. **Indexes**: An unindexed query on a large table is a production outage waiting to happen.
53. **Transactions**: Use transactions for operations that must be atomic (all or nothing).
54. **Backup Strategy**: A backup is not a backup until you have successfully restored from it.
55. **Security First**: Sanitize inputs. Parameterize queries (SQL Injection). Escape outputs (XSS).

## 🛠️ Workflow & Productivity (56-75)
56. **Master Your IDE**: Learn the shortcuts. If you use the mouse, you are slow.
57. **Terminal Proficiency**: Be comfortable in the shell. `grep`, `awk`, `sed`, `curl` are superpowers.
58. **Git Hygiene**: Commit often. Write meaningful commit messages. Don't commit broken code.
59. **Branching Strategy**: Use Feature Branches. Don't push directly to `main`.
60. **Pull Requests**: Keep them small. A 200-line PR is reviewed; a 2000-line PR is "looks good to me".
61. **Automate Everything**: If you do it twice, automate it. (Scripts, CI/CD).
62. **Continuous Integration (CI)**: Run tests automatically on every push.
63. **Continuous Deployment (CD)**: Automate the release process. Deploying should be one click.
64. **Documentation**: Document the "How to Run", "Architecture", and "API".
65. **Time Management**: Deep work requires uninterrupted blocks of time. Turn off notifications.
66. **Pomodoro Technique**: Work in focused bursts (25m) followed by short breaks.
67. **Rubber Duck Debugging**: Explain your problem to an inanimate object. You'll often find the solution yourself.
68. **Learn to Search**: Google is a skill. "golang http server example" vs "golang net/http listenandserve error 404".
69. **Read Other People's Code**: You learn more from reading open source code than writing your own.
70. **Pair Programming**: Two heads are better than one. It spreads knowledge and reduces bugs.
71. **Take Breaks**: Your brain solves problems while you are walking/sleeping.
72. **Ergonomics**: Invest in a good chair, keyboard, and monitor. Your body is your tool.
73. **Keyboard Typing Speed**: If you type slow, you think slow (in code). Practice touch typing.
74. **Snippets**: Create snippets for common code patterns (e.g., `if err != nil`).
75. **Backup Your Work**: Push to remote often. Hardware fails.

## 🐛 Debugging & Troubleshooting (76-85)
76. **Reproduce First**: You can't fix a bug you can't reproduce.
77. **Read the Error Message**: seriously, read it. It usually tells you exactly what is wrong.
78. **Binary Search**: Comment out half the code. If the bug persists, it's in the other half. Repeat.
79. **Check the Logs**: The server logs usually have the answer.
80. **Isolate the Variable**: Change one thing at a time.
81. **Verify Assumptions**: "I'm sure that variable is not null" -> Check it anyway.
82. **Check Network**: Is it a DNS issue? Firewall? Port closed?
83. **Check Permissions**: Is the file read-only? Does the user have access?
84. **Restart**: Sometimes, "Have you tried turning it off and on again?" actually works.
85. **Ask for Help**: Don't bang your head for 4 hours. Ask a colleague after 30 minutes of trying.

## 🚀 Career & Soft Skills (86-100)
86. **Communication is Key**: The best code is useless if you can't explain it or work with others.
87. **Under-promise, Over-deliver**: Better to say "3 days" and finish in 2, than say "1 day" and take 2.
88. **Be a Problem Solver, Not a Coder**: Companies pay you to solve business problems, not just type syntax.
89. **Learn the Domain**: Understand the business you are in (FinTech, HealthTech, E-commerce).
90. **Mentorship**: Find a mentor. Be a mentor. Teaching reinforces learning.
91. **Stay Curious**: Technology moves fast. Keep learning, but don't chase every shiny new toy.
92. **Network**: Your next job will likely come from a friend or colleague.
93. **Build a Portfolio**: Show your work. GitHub, Blog, Talks.
94. **Negotiate**: Know your worth. Don't be afraid to ask for a raise or better benefits.
95. **Work-Life Balance**: Burnout is real. A tired engineer is a bad engineer.
96. **Own Your Mistakes**: If you break production, admit it, fix it, and learn from it.
97. **Be Kind**: Tech is a small world. Don't be a jerk.
98. **Documentation is Visibility**: Writing good docs makes you visible to the whole organization.
99. **Say "No"**: Don't take on more than you can handle. Quality > Quantity.
100. **Enjoy the Journey**: Coding is a superpower. Have fun with it!
