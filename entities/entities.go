package entities

type SportInstructor struct {
	Name           string      `json:"name"`
	Sports         []string    `json:"sports"`
	AvailableDates []DateRange `json:"availableDates"`
	Score          float64     `json:"score"`
}

type DateRange struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type Filter struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type OrderBy struct {
	Field string `json:"field"`
	Asc   bool   `json:"asc"`
}
