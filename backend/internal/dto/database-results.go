package dto

// Holds sales results for specific year.
type YearlySalesData struct {
	Year *int32      `ch:"Year" json:"Year"`
	Data []SalesData `json:"Data"`
}

// Holds sales data for specific category. Generic to work
// with both region and genre for example.
type SalesData struct {
	Category string  `json:"Category"`
	Sales    float64 `json:"Sales"`
}
