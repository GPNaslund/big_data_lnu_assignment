package dto

// Holds query params.
type DataQuery struct {
	Aggregate string   `query:"aggregate"`
	GroupBy   string   `query:"group"`
	StartYear string   `query:"start-year"`
	EndYear   string   `query:"end-year"`
	Filters   []string `query:"filters"`
	Search    string   `query:"search"`
}
