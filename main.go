package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func parseTime(value string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, value)
	if err != nil {
		log.Fatalf("Invalid time format: %v", err)
	}
	return parsedTime
}

type Habit struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Frequency   string    `json:"frequency"`
	StartDate   time.Time `json:"startDate"`
}

type TrackingEntry struct {
	ID        string    `json:"id"`
	HabitID   string    `json:"habitID"`
	Timestamp time.Time `json:"timestamp"`
	Note      string    `json:"note"`
}

var habits = []Habit{
	{ID: "abc", Name: "Exercise", Description: "Exercise to get stronger!!", Frequency: "3 times weekly", StartDate: parseTime("2024-11-01T08:00:00Z")},
	{ID: "def", Name: "Read", Description: "Read more regurarly", Frequency: "Daily", StartDate: parseTime("2024-11-12T08:00:00Z")},
	{ID: "ghj", Name: "Cook more", Description: "Cook homecooked meals", Frequency: "Daily", StartDate: parseTime("2024-11-05T08:00:00Z")},
}

var entries = []TrackingEntry{
	{ID: "123", HabitID: "abc", Timestamp: parseTime("2024-11-02T15:00:00Z"), Note: "Went running."},
	{ID: "456", HabitID: "abc", Timestamp: parseTime("2024-11-05T09:00:00Z"), Note: "Went to the gym."},
	{ID: "789", HabitID: "abc", Timestamp: parseTime("2024-11-10T16:00:00Z"), Note: "Did some yoga."},
	{ID: "234", HabitID: "def", Timestamp: parseTime("2024-11-13T21:00:00Z"), Note: "Read one chapter."},
	{ID: "567", HabitID: "def", Timestamp: parseTime("2024-11-15T21:00:00Z"), Note: "Read two chapters."},
}

func getHabits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(habits)
}

func getHabitByID(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 || pathParts[2] == "" {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}
	id := pathParts[2]

	ok := false
	for _, habit := range habits {
		if habit.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(habit)
			ok = true
			return
		}
	}
	if !ok {
		http.Error(w, "Habit not found", http.StatusNotFound)
	}
}

func addHabit(w http.ResponseWriter, r *http.Request) {
	var newHabit Habit
	if err := json.NewDecoder(r.Body).Decode(&newHabit); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newHabit.StartDate = time.Now()
	newHabit.ID = uuid.New().String()
	habits = append(habits, newHabit)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newHabit)
}

func updateHabit(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 || pathParts[2] == "" {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}
	id := pathParts[2]

	var updates map[string]string
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, habit := range habits {
		if habit.ID == id {
			if name, ok := updates["name"]; ok {
				habits[i].Name = name
			}
			if desc, ok := updates["description"]; ok {
				habits[i].Description = desc
			}
			if freq, ok := updates["frequency"]; ok {
				habits[i].Frequency = freq
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(habits[i])
			return
		}
	}

	http.Error(w, "Habit not found", http.StatusNotFound)
}

func deleteHabit(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 || pathParts[2] == "" {
		http.Error(w, "Invalid habit ID", http.StatusBadRequest)
		return
	}
	id := pathParts[2]

	for i, habit := range habits {
		if habit.ID == id {
			habits = append(habits[:i], habits[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Habit not found", http.StatusNotFound)
}

func addTrackingEntry(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 4 || pathParts[2] == "" || pathParts[3] != "tracking" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	habitID := pathParts[2]

	var habitExists bool
	for _, habit := range habits {
		if habit.ID == habitID {
			habitExists = true
			break
		}
	}
	if !habitExists {
		http.Error(w, "Habit not found", http.StatusNotFound)
		return
	}

	var entry TrackingEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	entry.ID = uuid.New().String()
	entry.HabitID = habitID
	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now()
	}

	entries = append(entries, entry)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}

func main() {
	http.HandleFunc("/habits", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getHabits(w, r)
		case http.MethodPost:
			addHabit(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/habits/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/tracking") && r.Method == http.MethodPost {
			addTrackingEntry(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			getHabitByID(w, r)
		case http.MethodPatch:
			updateHabit(w, r)
		case http.MethodDelete:
			deleteHabit(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
