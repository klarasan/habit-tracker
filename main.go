package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	{ID: "123", HabitID: "abc", Timestamp: parseTime("2024-11-01T15:00:00Z"), Note: "Went running."},
	{ID: "456", HabitID: "abc", Timestamp: parseTime("2024-13-01T09:00:00Z"), Note: "Went to the gym."},
	{ID: "789", HabitID: "abc", Timestamp: parseTime("2024-15-01T16:00:00Z"), Note: "Did some yoga."},
	{ID: "234", HabitID: "def", Timestamp: parseTime("2024-12-01T21:00:00Z"), Note: "Read one chapter."},
	{ID: "567", HabitID: "def", Timestamp: parseTime("2024-13-01T21:00:00Z"), Note: "Read two chapters."},
}

func getHabits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(habits)
}

func main() {
	http.HandleFunc("/habits", getHabits)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
