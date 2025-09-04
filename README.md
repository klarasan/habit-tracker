# habit-tracker

A simple Go API for tracking habits and logging progress. Stores habits and tracking entries in memory and exposes REST endpoints for basic operations. Server runs on http://localhost:8080.

---

## Running the API

```bash
go run main.go
```

---

## Endpoints

GET /habits : Retrieve a list of all habits.
```bash
curl http://localhost:8080/habits
```

GET /habits/{id} : Retrieve details of a specific habit based on its ID.
```bash
curl http://localhost:8080/habits/abc
```

POST /habits : Add a new habit.
```bash
curl  -X POST http://localhost:8080/habits \
      -H "Content-Type: application/json" \
      -d '{"name":"Meditate","description":"Daily meditation","frequency":"Daily"}'
```

PATCH /habits/{id} : Update information for a specific habit.
```bash
curl  -X PATCH http://localhost:8080/habits/abc \
      -H "Content-Type: application/json" \
      -d '{"name":"Daily Exercise","frequency":"4 times weekly"}'
```

---

## Dependencies

This project uses Go modules.

External dependency:
- [github.com/google/uuid](https://pkg.go.dev/github.com/google/uuid) — for generating unique IDs.

When cloning the project, run:

```bash
go mod tidy