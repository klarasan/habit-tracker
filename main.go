package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Habit struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var habits = []Habit{
	{ID: "abc", Name: "Exercise", Description: "Exercise to get stronger!!"},
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
