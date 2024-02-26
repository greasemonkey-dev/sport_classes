package sources

import (
	"database/sql"
	"fmt"
)

type DBClient struct {
	db *sql.DB
}

func NewDBClient(dataSourceName string, username string, password string) (*DBClient, error) {
	// Build connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8", username, password, dataSourceName)

	// Connect to database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Test the connection and return client
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &DBClient{db: db}, nil
}

// Function to retrieve sport classes from database (replace with your actual query)
func (c *DBClient) GetSportClasses() ([]string, error) {

	rows, err := c.db.Query("SELECT DISTINCT sport_class FROM classes")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var sportClasses []string
	for rows.Next() {
		var sportClass string
		err := rows.Scan(&sportClass)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		sportClasses = append(sportClasses, sportClass)
	}

	return sportClasses, nil
}

// FilterByDateRange filters sport instructors by date range in the database
func (c *DBClient) FilterByDateRange(sport string, minScore float64, from int64) ([]string, error) {
	// TODO: Implementation specific to FilterByDateRange in the database
	return nil, nil
}
