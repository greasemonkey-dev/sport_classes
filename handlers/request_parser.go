package handlers

import (
	"net/http"
	"os"
	"sport_classes/sources"
	"strconv"
)

func Parse(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	sport := r.URL.Query().Get("sport")
	if sport == "" {
		http.Error(w, "missing sport parameter", http.StatusBadRequest)
	}

	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "missing date parameter", http.StatusBadRequest)
	}
	date, err := strconv.ParseInt(dateStr, 10, 64)

	minScoreStr := r.URL.Query().Get("minScore")

	// Parse and validate individual parameters
	minScore, err := strconv.ParseFloat(minScoreStr, 64)
	if err != nil {
		http.Error(w, "Invalid score format", http.StatusBadRequest)
		return
	}
	if !IsValidScore(minScore) {
		http.Error(w, "Score must be between 0 to 10", http.StatusBadRequest)
		return
	}
	dataSource := os.Getenv("DATA_SOURCE")
	client, err := sources.NewClient(dataSource)
	if err != nil {
		http.Error(w, "Error creating the data client", http.StatusBadRequest)
	}
	// Retrieve existing classes from the data client
	existingClasses, err := client.GetSportClasses()
	if err != nil {
		http.Error(w, "Error retrieving classes names", http.StatusBadRequest)
		return
	}
	if !IsValidSportClass(sport, existingClasses) {
		http.Error(w, "Invalid sport class name", http.StatusBadRequest)
	}

	// Use the parsed parameters to filter and retrieve data
	instructors, err := client.FilterByDateRange(sport, minScore, date)
	if err != nil {
		http.Error(w, "Error creating the data client", http.StatusBadRequest)
	}
	// Respond with the filtered instructors
	// (implementation depends on your desired response format)
	writeJSONResponse(w, instructors)
}
func writeJSONResponse(w http.ResponseWriter, instructors []string) {
	// Assuming instructors is a slice of strings containing instructor names
	// This is a basic example, you might want to structure your response accordingly
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, instructor := range instructors {
		// Assuming your response format is a JSON array of names
		w.Write([]byte(`"` + instructor + `",`))
	}
}
