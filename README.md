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
curl    -X POST http://localhost:8080/habits \
        -H "Content-Type: application/json" \
        -d '{"id":"xyz","name":"Meditate","description":"Daily meditation","frequency":"Daily"}'
```