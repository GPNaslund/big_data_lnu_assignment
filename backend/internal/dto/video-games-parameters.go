package dto

// Struct for holding data parameters.
type VideoGamesParameters struct {
	Aggregations []string `json:"aggregations"`
	Groupings    []string `json:"groupings"`
	Regions      []string `json:"regions"`
	Genres       []string `json:"genres"`
	StartYear    int      `json:"start-year"`
	EndYear      int      `json:"end-year"`
}
