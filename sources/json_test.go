package sources

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sport_classes/entities"
	"testing"
	"time"
)

func TestJSONClient(t *testing.T) {
	// Define a test JSON file path
	testJSONFilePath := "./test.json"
	createTestJSONFile(testJSONFilePath)

	// Create a new JSONClient for testing
	jsonClient, err := NewJSONClient(testJSONFilePath)
	if err != nil {
		t.Fatalf("Failed to create JSONClient: %v", err)
	}

	// Run tests
	t.Run("TestGetSportClasses", func(t *testing.T) {
		testGetSportClasses(t, jsonClient)
	})

	t.Run("TestFilterByDateRange", func(t *testing.T) {
		testFilterByDateRange(t, jsonClient)
	})
	removeTestJSONFile(testJSONFilePath)
}

// Helper function to create a temporary test JSON file
func createTestJSONFile(fileName string) {
	// Create test data
	testData := []entities.SportInstructor{
		{
			Name:   "Instructor1",
			Sports: []string{"SportA", "SportB"},
			AvailableDates: []entities.DateRange{
				{From: time.Now().Add(-time.Hour).Unix(), To: time.Now().Add(1 * time.Hour).Unix()},
			},
			Score: 8.5,
		}, {
			Name:   "Instructor2",
			Sports: []string{"SportC", "SportA"},
			AvailableDates: []entities.DateRange{
				{From: time.Now().Add(-time.Hour).Unix(), To: time.Now().Add(2 * time.Hour).Unix()},
			},
			Score: 8.9,
		}, {
			Name:   "Instructor3",
			Sports: []string{"SportC", "SportA"},
			AvailableDates: []entities.DateRange{
				{From: time.Now().Add(-time.Hour).Unix(), To: time.Now().Unix()},
			},
			Score: 7.0,
		},
	}

	// Convert to JSON
	jsonData, err := json.Marshal(testData)
	if err != nil {
		panic(err)
	}

	// Write to a temporary file
	err = ioutil.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

// Helper function to remove the temporary test JSON file
func removeTestJSONFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}

// Test GetSportClasses method
func testGetSportClasses(t *testing.T, client *JSONClient) {
	expectedClasses := []string{"SportA", "SportB", "SportC"} // Adjust based on your test data

	classes, err := client.GetSportClasses()
	if err != nil {
		t.Fatalf("Error retrieving sport classes: %v", err)
	}

	// Compare the obtained classes with the expected ones
	if !equalSlices(classes, expectedClasses) {
		t.Errorf("GetSportClasses result does not match expected classes")
	}
}

// Test FilterByDateRange method
func testFilterByDateRange(t *testing.T, client *JSONClient) {
	// Adjust based on your test data
	sport := "SportA"
	minScore := 8.0
	from := time.Now().Unix()

	expectedInstructors := []string{"Instructor2", "Instructor1"}

	instructors, err := client.FilterByDateRange(sport, minScore, from)
	if err != nil {
		t.Fatalf("Error filtering by date range: %v", err)
	}

	// Compare the obtained instructors with the expected ones
	if !equalSlices(instructors, expectedInstructors) {
		t.Errorf("FilterByDateRange result does not match expected instructors")
	}
}

// Helper function to check if two slices are equal
func equalSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}
