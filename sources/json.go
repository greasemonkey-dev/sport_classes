package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"sport_classes/entities"
)

// JSONClient is a type that implements DataClient for JSON sources
type JSONClient struct {
	filePath    string
	data        []entities.SportInstructor
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

	uniqueClassNames := make(map[string]struct{})

	// Extract and store unique class names
	for _, instructor := range c.data {
		for _, sport := range instructor.Sports {
			uniqueClassNames[sport] = struct{}{}
		}
	}

	// Convert map keys to slice
	var classNames []string
	for className := range uniqueClassNames {
		classNames = append(classNames, className)
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
				if dateRange.From <= from && dateRange.To > from {
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
