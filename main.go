package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
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
	habits = append(habits, newHabit)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newHabit)
}

func main() {
	http.HandleFunc("/habits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getHabits(w, r)
		} else if r.Method == http.MethodPost {
			addHabit(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/habits/", getHabitByID)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
