# Go for DevOps – Learning Journey

This repository documents my 6-week journey of learning the Go programming language with a focus on DevOps, Cloud, and Kubernetes.

The goal is to move from Go fundamentals to production-ready patterns used in real-world cloud-native systems.

---

## 🚀 Progress

✔ Day 1 — Go module + first binary  
✔ Day 2 — Concurrency with Goroutines  
✔ Day 3 — Channels + WaitGroups  
✔ Day 4 — Built first CLI tool using the `flag` package  
✔ Day 5 — JSON parsing & file handling  
✔ Day 6 — Parallel API calls using Goroutines  
✔ Day 7 — Reliable API fetcher with retries, timeouts, and logging  
✔ Day 8 — Interfaces & polymorphism (multiple loggers via a single contract)  
✔ Day 9 — Go error handling (wrapping, custom errors, production patterns)  
✔ Day 10 — Real-world Go project structure (`cmd/`, `internal/`, dependency injection)  
✔ Day 11 — Configuration & Environment Variables (12-Factor App: Config)  
✔ Day 12 — Flags + Environment Variables (priority-based configuration)  
✔ Day 13 — Graceful Shutdown & OS signal handling (SIGINT, SIGTERM)  
✔ Day 14 — HTTP Server with Graceful Shutdown (Go for DevOps)

### Concepts  
- net/http  
- OS signal handling  
- Graceful shutdown  
- Kubernetes-ready services  

### Run  
```bash  
go run main.go  
```  

### Test  
```bash  
curl http://localhost:9090  
```




---

## 📂 Repository Structure

- `week1/day1` — Go basics and first Go program
- `week1/day2` — Goroutines (Concurrency)
- `week1/day3` — Channels + WaitGroups
- `week1/day4` — CLI tool using `flag` package
- `week1/day5` — JSON parsing CLI tool
- `week1/day6` — Parallel API fetcher using Goroutines
- `week1/day7` — Reliable API fetcher (retry + timeout + logging)
- `week1/day8` — Interfaces & polymorphism
- `week1/day9` — Error handling in Go
- `week1/day10` — Real-world Go project structure
- `week2/day11` — Configuration & Environment Variables
- `week2/day12` — Flags + Environment Variables (Production-Ready Config)  
- `week2/day13` — Graceful Shutdown & OS Signals (Production Essential)  
- `week2/day14` — HTTP Server with Graceful Shutdown 

---

## ▶️ How to Run Each Day

Each day is organized as an **independent Go module**.  
You can run any day individually.

### 1️⃣ Navigate to the desired day

```bash
cd week1/dayX
```

Replace dayX with the day you want to run (e.g., day3, day8, day10).

### 2️⃣ Run the program (development mode)

```bash
go run .
```

This command:

- Compiles the code temporarily
- Runs the program
- Does not create a permanent binary

### 3️⃣ Build a binary (recommended practice)

```bash
go build -o app
```

Run the compiled binary:

```bash
./app
```

Compiled binaries are intentionally ignored using .gitignore.

### 4️⃣ Running projects with cmd/ structure (Day-10 onwards)

Some days use a real-world Go project layout with cmd/ and internal/.

Run the application:

```bash
go run ./cmd/app
```

Build the binary:

```bash
go build -o app ./cmd/app
```

```bash
./app
```

### 5️⃣ Managing dependencies

Each day has its own go.mod file.

If you add, remove, or change imports:

```bash
go mod tidy
```

This keeps dependencies clean and reproducible.

## 📝 Notes

- Each day is self-contained
- Generated files (binaries, logs) are not committed
- The structure mirrors production-grade Go projects
- Concepts progress from fundamentals to real-world patterns
