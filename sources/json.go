package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sport_classes/entities"
)

// JSONClient is a type that implements DataClient for JSON sources
type JSONClient struct {
	filePath    string
	data        []entities.SportInstructor // Assuming SportInstructor is defined in your entities package
	initialized bool
}

// NewJSONClient creates a new JSONClient instance and reads the JSON file
func NewJSONClient(filePath string) (*JSONClient, error) {
	client := &JSONClient{filePath: filePath}

	if err := client.init(); err != nil {
		return nil, err
	}

	return client, nil
}

// GetSportClasses retrieves sport classes from the stored JSON data
func (c *JSONClient) GetSportClasses() ([]string, error) {
	if !c.initialized {
		if err := c.init(); err != nil {
			return nil, err
		}
	}

	// Extract class names
	var classNames []string
	for _, instructor := range c.data {
		classNames = append(classNames, instructor.Name)
	}

	return classNames, nil
}

// init reads and parses the JSON file
func (c *JSONClient) init() error {
	content, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %w", err)
	}

	// Unmarshal JSON data
	err = json.Unmarshal(content, &c.data)
	if err != nil {
		return fmt.Errorf("error parsing JSON data: %w", err)
	}

	c.initialized = true
	return nil
}

// FilterByDateRange filters sport instructors by date range and sport order by instructor's score
func (c *JSONClient) FilterByDateRange(sport string, minScore float64, from int64) ([]string, error) {
	if !c.initialized {
		if err := c.init(); err != nil {
			return nil, err
		}
	}

	// Filter and sort the instructors
	var filteredInstructors []entities.SportInstructor
	for _, instructor := range c.data {
		if containsSport(instructor.Sports, sport) && instructor.Score >= minScore {
			for _, dateRange := range instructor.AvailableDates {
				if dateRange.From >= from {
					filteredInstructors = append(filteredInstructors, instructor)
					break
				}
			}
		}
	}

	// Sort the instructors by score in descending order
	sort.Slice(filteredInstructors, func(i, j int) bool {
		return filteredInstructors[i].Score > filteredInstructors[j].Score
	})

	// Extract and return the instructor names
	var availableNames []string
	for _, instructor := range filteredInstructors {
		availableNames = append(availableNames, instructor.Name)
	}

	return availableNames, nil
}

// containsSport checks if a given sport is in the list of sports
func containsSport(sports []string, targetSport string) bool {
	for _, s := range sports {
		if s == targetSport {
			return true
		}
	}
	return false
}

// Helper function to read file content
func readFile(filePath string) (string, error) {
	// Open the file with read-only permissions
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure file is closed even if errors occur

	// Read all contents of the file into a byte slice
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// Return the content as a string
	return string(content), nil
}
