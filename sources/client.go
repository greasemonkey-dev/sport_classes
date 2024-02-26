package sources

import (
	"fmt"
	"os"
)

type DataClient interface {
	GetSportClasses() ([]string, error)
	FilterByDateRange(sport string, minScore float64, date int64) ([]string, error)
}

func NewClient(dataSource string) (DataClient, error) {
	switch dataSource {
	case "JSON": // Return file path for JSON source
		filePath := os.Getenv("JSON_PATH")
		return NewJSONClient(filePath)
	case "DB": // Return database client for DB source
		dbClient, err := NewDBClient(os.Getenv("HOST_NAME"), os.Getenv("USER_NAME"), os.Getenv("PASSWORD"))
		if err != nil {
			return nil, err
		}
		return dbClient, nil
	default:
		return nil, fmt.Errorf("Unsupported data source type: %T", dataSource)
	}
}
